package alipay

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/guestin/go_pay_cli/internal"
	"github.com/guestin/mob/mvalidate"
	"github.com/ooopSnake/assert.go"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

var defaultReqTimeout = time.Second * 60

type Request interface {
	ApiName() string
	NotifyUrl() string
	NeedEncrypt() bool
	HasBizContent() bool
	Params() url.Values
}

type Response interface {
	error
	setCodeMsg(code, msg string)
	GetCode() string
	IsCodeSuccess() bool
	GetMsg() string
	setSubCodeMsg(subCode, subMsg string)
	GetSubCode() string
	GetSubMsg() string
	IsSystemError() bool
	setRespContent(content string)
	GetRespContent() string
}

type responseInternal struct {
	Sign         string      `json:"sign" validate:"required"`
	AliPayCertSn string      `json:"alipay_cert_sn"`
	ContentRaw   string      `json:"-"`
	IsErrorResp  bool        `json:"-"`
	Content      interface{} `json:"-"`
	errResp      error
}

func (this *responseInternal) Err() error {
	return this.errResp
}

type client struct {
	ctx                context.Context
	validator          mvalidate.Validator
	gateway            string
	sandbox            bool
	appId              string
	format             string
	charset            string
	apiVersion         string
	signType           string
	appPrivateKeyRaw   string
	appPrivateKey      *rsa.PrivateKey
	alipayPublicKeyRaw string
	alipayPublicKey    *rsa.PublicKey
	signProvider       AsymmetricEncryptor
	encryptType        string
	encryptKey         string
	appAuthToken       string
	sellerId           string
	isvPid             string
	rwLock             *sync.RWMutex
	rootCertContent    string //支付宝根证书内容，用于验证下载的支付宝公钥证书有效性
	rootCertSN         string //支付宝根证书序列号，用于每次调动OpenAPI时发送给网关
	aliPublicCertSN    string
	appCertSN          string //商户证书序列号，用于每次调动OpenAPI时发送给网关
	aliPublicKeyList   map[string]*rsa.PublicKey
}

func DefaultAliPayClient(appId, appPrivateKey, signType string, sandbox ...bool) Client {
	assert.Must(len(appId) > 0, "app id must not be empty").Panic()
	assert.Must(len(appPrivateKey) > 0, "private key must not be empty").Panic()
	assert.Must(signType == "RSA" || signType == "RSA2", "sign type must be one of ['RSA','RSA2']").Panic()
	appPriKey, err := parserPrivateKey(appPrivateKey)
	assert.Must(err != nil, "invalid app private key").Panic()
	apiUrl := productionApiUrl
	sb := len(sandbox) > 0 && sandbox[0]
	if sb {
		apiUrl = sandBoxApiUrl
	}
	validator, err := mvalidate.NewValidator("en")
	assert.NoError(err).Panic()
	out := &client{
		ctx:              context.Background(),
		validator:        validator,
		gateway:          apiUrl,
		sandbox:          sb,
		appId:            appId,
		format:           DefaultFormat,
		charset:          DefaultCharset,
		apiVersion:       ApiVersion,
		signType:         signType,
		appPrivateKeyRaw: appPrivateKey,
		appPrivateKey:    appPriKey,
		signProvider:     GetEncryptor(signType),
		encryptType:      "",
		encryptKey:       "",
		rwLock:           &sync.RWMutex{},
		rootCertContent:  "",
		rootCertSN:       "",
		appCertSN:        "",
		aliPublicKeyList: map[string]*rsa.PublicKey{},
	}
	return out
}

func (this *client) LoadAliPayPublicKey(alipayPublicKey string) error {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	key, err := parserPublicKey(alipayPublicKey)
	if err != nil {
		return errors.Wrap(err, "invalid alipay public key :")
	}
	this.alipayPublicKey = key
	this.alipayPublicKeyRaw = alipayPublicKey
	return nil
}

func (this *client) LoadCert(appPublicKeyRaw, alipayRootCertRaw, alipayPublicCertRaw []byte) error {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	//APP public key cert
	appPublicCert, err := ParseCertificate(appPublicKeyRaw)
	if err != nil {
		return errors.Wrap(err, "invalid app public key cert")
	}
	this.appCertSN = getCertSN(appPublicCert)

	//AliPay root cert
	var alipayRootCertStrList = strings.Split(string(alipayRootCertRaw), kCertificateEnd)
	var alipayRootCertSNList = make([]string, 0, len(alipayRootCertStrList))
	for _, certStr := range alipayRootCertStrList {
		certStr = certStr + kCertificateEnd
		var cert, _ = ParseCertificate([]byte(certStr))
		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			alipayRootCertSNList = append(alipayRootCertSNList, getCertSN(cert))
		}
	}
	buf := bytes.Buffer{}
	buf.Write(alipayPublicCertRaw)
	this.rootCertContent = buf.String()
	this.rootCertSN = strings.Join(alipayRootCertSNList, "_")

	//AliPay public cert
	alipayPublicKeyCert, err := ParseCertificate(alipayPublicCertRaw)
	if err != nil {
		return err
	}
	//cache
	this.aliPublicCertSN = getCertSN(alipayPublicKeyCert)
	key, ok := alipayPublicKeyCert.PublicKey.(*rsa.PublicKey)
	if ok {
		this.aliPublicKeyList[this.aliPublicCertSN] = key
	}
	return nil
}

func (this *client) LoadCertFromFile(appPublicKeyFile, alipayRootCertFile, alipayPublicCertFile string) error {
	appPublicKeyRaw, err := ioutil.ReadFile(appPublicKeyFile)
	if err != nil {
		return errors.Wrapf(err, "load app public key from file '%s' failed :", appPublicKeyFile)
	}
	alipayRootCertRaw, err := ioutil.ReadFile(alipayRootCertFile)
	if err != nil {
		return errors.Wrapf(err, "load alipay root cert from file '%s' failed :", alipayRootCertFile)
	}
	alipayPublicCertRaw, err := ioutil.ReadFile(alipayPublicCertFile)
	if err != nil {
		return errors.Wrapf(err, "load alipay public key cert from file '%s' failed :", alipayPublicCertFile)
	}
	return this.LoadCert(appPublicKeyRaw, alipayRootCertRaw, alipayPublicCertRaw)
}

func (this *client) LoadCertFromBase64(appPublicKey, alipayRootCert, alipayPublicCert string) error {
	appPublicKeyRaw, err := base64.StdEncoding.DecodeString(appPublicKey)
	if err != nil {
		return errors.Wrap(err, "load app public key from base64 str failed :")
	}
	alipayRootCertRaw, err := base64.StdEncoding.DecodeString(alipayRootCert)
	if err != nil {
		return errors.Wrap(err, "load alipay root cert from base64 str failed :")
	}
	alipayPublicCertRaw, err := base64.StdEncoding.DecodeString(alipayPublicCert)
	if err != nil {
		return errors.Wrap(err, "load alipay public key cert from  base64 str failed :")
	}
	return this.LoadCert(appPublicKeyRaw, alipayRootCertRaw, alipayPublicCertRaw)
}

func (this *client) GetAppId() string {
	return this.appId
}

func (this *client) SetAppAuthToken(appAuthToken string) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	this.appAuthToken = appAuthToken
}

func (this *client) GetAppAuthToken() string {
	this.rwLock.RLock()
	defer this.rwLock.RUnlock()
	return this.appAuthToken
}

func (this *client) SetSellerId(sellerId string) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	this.sellerId = sellerId
}

func (this *client) GetSellerId() string {
	this.rwLock.RLock()
	defer this.rwLock.RUnlock()
	return this.sellerId
}

func (this *client) SetISVPid(isvPid string) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	this.isvPid = isvPid
}

func (this *client) GetISVPid() string {
	this.rwLock.RLock()
	defer this.rwLock.RUnlock()
	return this.isvPid
}

func (this *client) prepareReq(in Request) (url.Values, error) {
	err := this.validator.Struct(in)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("app_id", this.appId)
	params.Set("method", in.ApiName())
	params.Set("format", this.format)
	params.Set("charset", this.charset)
	params.Set("sign_type", this.signType)
	params.Set("timestamp", internal.TimeNow().Format(longTimeFormat))
	params.Set("version", this.apiVersion)
	if in.ApiName() != APISystemOauthToken &&
		//in.ApiName() != APIFundAuthOrderAppFreeze &&
		len(this.appAuthToken) > 0 {
		params.Set("app_auth_token", this.appAuthToken)
	}
	if this.appCertSN != "" {
		params.Set("app_cert_sn", this.appCertSN)
	}
	if this.rootCertSN != "" {
		params.Set("alipay_root_cert_sn", this.rootCertSN)
	}
	extParams := in.Params()
	if len(extParams) > 0 {
		for k := range extParams {
			v := extParams.Get(k)
			if len(v) > 0 {
				params.Set(k, v)
			}
		}
	}
	if in.HasBizContent() {
		biz, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		params.Set("biz_content", string(biz))
	}
	notifyUrl := in.NotifyUrl()
	if len(notifyUrl) > 0 {
		params.Add("notify_url", notifyUrl)
	}
	if in.NeedEncrypt() {
		//todo
	}
	sign, err := this.sign(params, this.appPrivateKey)
	if err != nil {
		return nil, err
	}
	params.Set("sign", sign)
	return params, nil
}

func (this *client) execute(in Request, out Response) error {
	assert.Must(in != nil, "in must not be nil").Panic()
	assert.Must(out != nil, "out must not be nil").Panic()
	params, err := this.prepareReq(in)
	if err != nil {
		return err
	}
	header := map[string]string{
		"Content-Type": fmt.Sprintf("%s%s", "application/x-www-form-urlencoded;charset=", this.charset),
	}
	//do post
	respBody, err := internal.HttpDo(this.ctx, defaultReqTimeout, http.DefaultClient, "POST",
		fmt.Sprintf("%s?charset=%s", this.gateway, this.charset), params.Encode(), header)
	if err != nil {
		return err
	}
	parser := getAopParser(this.format)
	resp, err := parser.parse(in, string(respBody))
	if err != nil {
		return err
	}
	out.setRespContent(resp.ContentRaw)
	//verify sign
	err = this.verifyRespSign(in, resp)
	if err != nil {
		return err
	}
	if resp.IsErrorResp {
		return resp.Err()
	}
	return parser.unMarshall(in, resp, out)
}

func (this *client) sdkExecute(in Request, out Response) error {
	assert.Must(in != nil, "in must not be nil").Panic()
	assert.Must(out != nil, "out must not be nil").Panic()
	params, err := this.prepareReq(in)
	if err != nil {
		return err
	}
	out.setCodeMsg(CodeSuccess, "OK")
	out.setSubCodeMsg(CodeSuccess, "OK")
	out.setRespContent(params.Encode())
	return nil
}

func (this *client) addAliPayPublicKey(cert *x509.Certificate) {
	assert.Must(cert != nil, "cert must not be nil").Panic()
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	certSn := getCertSN(cert)
	this.aliPublicKeyList[certSn] = cert.PublicKey.(*rsa.PublicKey)
}

func (this *client) getAliPublicKey(sn string) (*rsa.PublicKey, error) {
	this.rwLock.RLock()
	key, ok := this.aliPublicKeyList[sn]
	if ok {
		this.rwLock.RUnlock()
		return key, nil
	}
	this.rwLock.RUnlock()
	//download cert
	downResp, err := this.CertDownload(&CertDownloadReq{
		AlipayCertSn: sn,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "download alipay public cert '%s'", sn)
	}
	certBytes, err := base64.StdEncoding.DecodeString(downResp.AlipayCertContent)
	cert, err := ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}
	err = VerifyCertChain(cert, []byte(this.rootCertContent))
	if err != nil {
		return nil, errors.Wrapf(err, "verify download alipay public cert cert failed :")
	}
	key, ok = cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("download invalid alipay public cert")
	}
	this.addAliPayPublicKey(cert)
	return key, nil
}

func (this *client) sign(param url.Values, privateKey *rsa.PrivateKey) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}
	var paramList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			paramList = append(paramList, key+"="+value)
		}
	}
	sort.Strings(paramList)
	var src = strings.Join(paramList, "&")
	sign, err := this.signProvider.Sign([]byte(src), privateKey)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sign)
	return s, nil
}

func (this *client) verifyRespSign(req Request, resp *responseInternal) (err error) {
	if req.ApiName() == APICertDownload {
		return nil
	}
	if len(resp.Sign) == 0 || (len(resp.AliPayCertSn) == 0 && len(this.alipayPublicKeyRaw) == 0) {
		return nil
	}
	key := this.alipayPublicKey
	//公钥证书
	if len(resp.AliPayCertSn) > 0 {
		key, err = this.getAliPublicKey(resp.AliPayCertSn)
		if err != nil {
			return err
		}
	}
	return this.signProvider.Verify([]byte(resp.ContentRaw), key, resp.Sign)
}
