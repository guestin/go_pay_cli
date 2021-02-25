package wxpay

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/guestin/go_pay_cli/internal"
	"github.com/guestin/mob/mvalidate"
	"github.com/ooopSnake/assert.go"
	"github.com/pkg/errors"
	"io"
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
	toUrlValues() url.Values
}

type client struct {
	ctx            context.Context
	rwLock         *sync.RWMutex
	validator      mvalidate.Validator
	sandbox        bool
	apiBaseUrl     string
	signType       string //签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	apiKey         string //签名秘钥
	appSecret      string //公众号app secret
	appId          string //普通商户：微信支付分配的公众账号ID（企业号corpid即为此appId）;子商户：服务商商户的APPID
	mchId          string //普通商户：微信支付分配的商户号;子商户：微信支付分配的商户号
	subAppId       string //普通商户：无；子商户：微信分配的子商户公众账号ID，如需在支付完成后获取sub_openid则此参数必传。
	subMchId       string //普通商户：无；子商户：微信支付分配的子商户号
	signProvider   SignProvider
	sandBoxSignKey string //缓存沙箱密钥
	tlsClient      *http.Client
	httpCli        *http.Client
}

func DefaultWxPayClient(appId, mchId, apiKey, appSecret, subAppId, subMchId, signType string, sandbox ...bool) Client {
	assert.Must(len(appId) > 0, "app id must not be null").Panic()
	assert.Must(len(mchId) > 0, "mch id must not be null").Panic()
	assert.Must(len(apiKey) > 0, "api key must not be null ").Panic()
	assert.Must(len(appSecret) > 0, "app secret must not be null ").Panic()
	assert.Must(signType == "MD5" || signType == "HMAC-SHA256",
		"sign type must be one of ['MD5','HMAC-SHA256']").Panic()
	apiUrl := ProductionApiUrl
	sb := len(sandbox) > 0 && sandbox[0]
	if sb {
		apiUrl = SandBoxApiUrl
	}
	validator, err := mvalidate.NewValidator("en")
	assert.NoError(err).Panic()
	cli := &client{
		ctx:            context.Background(),
		rwLock:         &sync.RWMutex{},
		validator:      validator,
		sandbox:        sb,
		apiBaseUrl:     apiUrl,
		signType:       signType,
		apiKey:         apiKey,
		appId:          appId,
		mchId:          mchId,
		appSecret:      appSecret,
		subAppId:       subAppId,
		subMchId:       subMchId,
		signProvider:   GetSignProvider(signType),
		sandBoxSignKey: "",
		httpCli:        http.DefaultClient,
	}
	return cli
}

func (this *client) LoadCert(raw []byte) error {
	assert.Must(len(raw) > 0, "cert data must not be null").Panic()
	cert, err := P12ToPem(raw, this.mchId)
	if err != nil {
		return err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	transport := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	this.setTlsCli(&http.Client{Transport: transport})
	return nil
}

func (this *client) LoadCertFromFile(p12File string) error {
	assert.Must(len(p12File) > 0, "p12File must not be null").Panic()
	raw, err := ioutil.ReadFile(p12File)
	if err != nil {
		return err
	}
	return this.LoadCert(raw)
}

func (this *client) LoadCertFromBase64(certBase64Str string) error {
	assert.Must(len(certBase64Str) > 0, "certBase64Str must not be null").Panic()
	raw, err := base64.StdEncoding.DecodeString(certBase64Str)
	if err != nil {
		return err
	}
	return this.LoadCert(raw)
}

func (this *client) setTlsCli(cli *http.Client) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	this.tlsClient = cli
}

func (this *client) getTlsCli() *http.Client {
	assert.Must(this.tlsClient != nil, "cert not load yet,please load cert first").Panic()
	this.rwLock.RLock()
	defer this.rwLock.RUnlock()
	return this.tlsClient
}

func (this *client) getSignKey(apiName string) (string, error) {
	if !this.sandbox || apiName == APIPayGetSignKey {
		return this.apiKey, nil
	}
	//get cache
	if len(this.sandBoxSignKey) > 0 {
		return this.sandBoxSignKey, nil
	}
	req := &payGetSignKeyReq{}
	resp := new(payGetSignKeyResp)
	err := this.executeWithoutCert(req, resp)
	if err != nil {
		return "", errors.Wrap(err, "get sandbox sign key failed :")
	}
	if !resp.IsSuccess() {
		return "", errors.Wrap(resp, "get sandbox sign key failed :")
	}
	this.sandBoxSignKey = resp.SandboxSignkey
	return resp.SandboxSignkey, nil
}

func (this *client) GetAppId() string {
	return this.appId
}
func (this *client) GetMchId() string {
	return this.subAppId
}

func (this *client) GetSubAppId() string {
	return this.subAppId
}

func (this *client) GetSubMchId() string {
	return this.subMchId
}

func (this *client) buildUrl(in Request) string {
	return fmt.Sprintf("%s%s", strings.TrimRight(this.apiBaseUrl, "/"), in.ApiName())
}

func (this *client) executeWithoutCert(in Request, out interface{}) error {
	return this.execute(this.httpCli, in, out)
}

func (this *client) executeWithCert(in Request, out interface{}) error {
	httpCli := this.getTlsCli()
	return this.execute(httpCli, in, out)
}

func (this *client) execute(httpCli *http.Client, in Request, out interface{}) error {
	err := this.validator.Struct(in)
	if err != nil {
		return err
	}
	key := this.apiKey
	key, err = this.getSignKey(in.ApiName())
	if err != nil {
		return err //errors.Wrap(err, "get sandbox sign key failed :")
	}
	params := in.toUrlValues()
	//params := make(XMLMap)
	//pb, err := xml.Marshal(in)
	//if err != nil {
	//	return err
	//}
	//err = xml.Unmarshal(pb, &params)
	//if err != nil {
	//	return err
	//}
	params.Set("appid", this.appId)
	params.Set("mch_id", this.mchId)
	if this.subAppId != "" && this.subMchId != "" {
		params.Set("sub_appid", this.subAppId)
		params.Set("sub_mch_id", this.subMchId)
	}
	params.Set("sign_type", this.signType)
	params.Set("nonce_str", internal.RandomString(16))
	notifyUrl := in.NotifyUrl()
	if len(notifyUrl) > 0 {
		params.Add("notify_url", notifyUrl)
	}
	sign, err := this.sign(params, key)
	if err != nil {
		return err
	}
	params.Set("sign", sign)
	header := map[string]string{
		"Accept":       "application/xml",
		"Content-Type": "application/xml;charset=utf-8",
	}
	apiUrl := fmt.Sprintf("%s%s", strings.TrimRight(this.apiBaseUrl, "/"), in.ApiName())
	reqBody := URLValueToXML(params)
	//do post
	respBody, err := internal.HttpDo(this.ctx, defaultReqTimeout, httpCli, "POST", apiUrl, reqBody, header)
	if err != nil {
		return err
	}
	if in.ApiName() == APIPayDownloadBill {
		_ = xml.Unmarshal(respBody, out)
		downResp := out.(*PayDownloadBillResp)
		if !downResp.IsSuccess() {
			return downResp
		}
		downResp.Data = respBody
		return nil
	}
	if err = this.verifyRespSign(respBody, key); err != nil {
		return err
	}
	return xml.Unmarshal(respBody, out)
}

func (this *client) makeStrToSign(param url.Values, key string) string {
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
	if len(key) > 0 {
		paramList = append(paramList, "key="+key)
	}
	return strings.Join(paramList, "&")
}

func (this *client) sign(param url.Values, key string) (s string, err error) {
	str := this.makeStrToSign(param, key)
	//log.Println("sign str = ", str)
	sign, err := this.signProvider.Sign([]byte(str), key)
	if err != nil {
		return "", err
	}
	return sign, nil
}

func (this *client) verifyRespSign(data []byte, key string) (err error) {
	var param = make(XMLMap)
	err = xml.Unmarshal(data, &param)
	if err != nil {
		return err
	}
	sign := param.Get("sign")
	if len(sign) == 0 {
		return nil
	}
	delete(param, "sign")
	str := this.makeStrToSign(url.Values(param), key)
	err = this.signProvider.Verify([]byte(str), key, sign)
	if err != nil {
		return errors.Wrap(err, "验签失败")
	}
	return nil
}

type XMLMap url.Values

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m XMLMap) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	for {
		var e xmlMapEntry
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(m)[e.XMLName.Local] = []string{e.Value}
	}
	return nil
}

func (v XMLMap) Get(key string) string {
	if v == nil {
		return ""
	}
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func (v XMLMap) Set(key, value string) {
	v[key] = []string{value}
}

func (v XMLMap) Add(key, value string) {
	v[key] = append(v[key], value)
}

func (v XMLMap) Del(key string) {
	delete(v, key)
}
