package wxpay

import (
	"encoding/json"
	"fmt"
	"github.com/guestin/go_pay_cli/internal"
	"github.com/pkg/errors"
	"net/url"
)

type CertLoader interface {
	LoadCert(raw []byte) error
	LoadCertFromFile(p12File string) error
	LoadCertFromBase64(certBase64Str string) error
}

type Client interface {
	CertLoader
	GetAppId() string
	GetMchId() string
	GetSubAppId() string
	GetSubMchId() string
	Execute(in Request, out interface{}) error
	ExecuteWithCert(in Request, out interface{}) error
	// https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_1
	PayUnifiedOrder(req *PayUnifiedOrderReq) (resp *PayUnifiedOrderResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_10&index=1
	PayMicroPay(req *PayMicroPayReq) (resp *PayMicroPayResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_2
	PayOrderQuery(req *PayOrderQueryReq) (resp *PayOrderQueryResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_11&index=3
	PayReverse(req *PayReverseReq) (resp *PayReverseResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_3
	PayCloseOrder(req *PayCloseOrderReq) (resp *PayCloseOrderResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_4
	PayRefund(req *PayRefundReq) (resp *PayRefundResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_5
	PayRefundQuery(req *PayRefundQueryReq) (resp *PayRefundQueryResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_6
	PayDownloadBill(req *PayDownloadBillReq) (resp *PayDownloadBillResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
	GetBrandWCPayRequest(req *GetBrandWCPayRequestReq) (resp *GetBrandWCPayRequestResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_0&index=1
	DepositFacePay(req *DepositFacePayReq) (resp *DepositFacePayResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_1&index=2
	DepositMicroPay(req *DepositMicroPayReq) (resp *DepositMicroPayResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_2&index=3
	DepositOrderQuery(req *DepositOrderQueryReq) (resp *DepositOrderQueryResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_3&index=4
	DepositReverse(req *DepositReverseReq) (resp *DepositReverseResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_4&index=5
	DepositConsume(req *DepositConsumeReq) (resp *DepositConsumeResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_5&index=6
	DepositRefund(req *DepositRefundReq) (resp *DepositRefundResp, err error)
	//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_6&index=7
	DepositRefundQuery(req *DepositRefundQueryReq) (resp *DepositRefundQueryResp, err error)

	//https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#1
	UserAccessToken(req *UserAccessTokenReq) (resp *UserAccessTokenResp, err error)
	UserRefreshToken(req *UserRefreshTokenReq) (resp *UserAccessTokenResp, err error)

	VerifySign(data []byte) error
}

func (this *client) Execute(in Request, out interface{}) error {
	return this.executeWithoutCert(in, out)
}

func (this *client) ExecuteWithCert(in Request, out interface{}) error {
	return this.executeWithCert(in, out)
}

func (this *client) PayUnifiedOrder(req *PayUnifiedOrderReq) (resp *PayUnifiedOrderResp, err error) {
	resp = new(PayUnifiedOrderResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) PayMicroPay(req *PayMicroPayReq) (resp *PayMicroPayResp, err error) {
	resp = new(PayMicroPayResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) PayOrderQuery(req *PayOrderQueryReq) (resp *PayOrderQueryResp, err error) {
	resp = new(PayOrderQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) PayReverse(req *PayReverseReq) (resp *PayReverseResp, err error) {
	resp = new(PayReverseResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) PayCloseOrder(req *PayCloseOrderReq) (resp *PayCloseOrderResp, err error) {
	resp = new(PayCloseOrderResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) PayRefund(req *PayRefundReq) (resp *PayRefundResp, err error) {
	resp = new(PayRefundResp)
	err = this.ExecuteWithCert(req, resp)
	return
}

func (this *client) PayRefundQuery(req *PayRefundQueryReq) (resp *PayRefundQueryResp, err error) {
	resp = new(PayRefundQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) PayDownloadBill(req *PayDownloadBillReq) (resp *PayDownloadBillResp, err error) {
	resp = new(PayDownloadBillResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) GetBrandWCPayRequest(req *GetBrandWCPayRequestReq) (resp *GetBrandWCPayRequestResp, err error) {
	if len(req.PrepayId) == 0 {
		return nil, errors.New("prepayId 为必需参数")
	}
	key := this.apiKey
	key, err = this.getSignKey(req.ApiName())
	if err != nil {
		return nil, err //errors.Wrap(err, "get sandbox sign key failed :")
	}
	now := internal.TimeNow()
	appId := this.appId
	if len(req.AppId) > 0 {
		appId = req.AppId
	}
	resp = &GetBrandWCPayRequestResp{
		AppId:     appId,
		Timestamp: fmt.Sprintf("%d", now.Unix()),
		NonceStr:  internal.RandomString(16),
		Package:   fmt.Sprintf("prepay_id=%s", req.PrepayId),
		SignType:  this.signType,
	}
	params := url.Values{}
	params.Set("appId", resp.AppId)
	params.Set("timeStamp", resp.Timestamp)
	params.Set("signType", resp.SignType)
	params.Set("nonceStr", resp.NonceStr)
	params.Set("package", resp.Package)
	sign, err := this.sign(params, key)
	if err != nil {
		return nil, err
	}
	//params.Set("sign", sign)
	resp.PaySign = sign
	return resp, nil
}

func (this *client) DepositFacePay(req *DepositFacePayReq) (resp *DepositFacePayResp, err error) {
	resp = new(DepositFacePayResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) DepositMicroPay(req *DepositMicroPayReq) (resp *DepositMicroPayResp, err error) {
	resp = new(DepositMicroPayResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) DepositOrderQuery(req *DepositOrderQueryReq) (resp *DepositOrderQueryResp, err error) {
	resp = new(DepositOrderQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) DepositReverse(req *DepositReverseReq) (resp *DepositReverseResp, err error) {
	resp = new(DepositReverseResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) DepositConsume(req *DepositConsumeReq) (resp *DepositConsumeResp, err error) {
	resp = new(DepositConsumeResp)
	err = this.ExecuteWithCert(req, resp)
	return
}

func (this *client) DepositRefund(req *DepositRefundReq) (resp *DepositRefundResp, err error) {
	resp = new(DepositRefundResp)
	err = this.ExecuteWithCert(req, resp)
	return
}

func (this *client) DepositRefundQuery(req *DepositRefundQueryReq) (resp *DepositRefundQueryResp, err error) {
	resp = new(DepositRefundQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) UserAccessToken(req *UserAccessTokenReq) (resp *UserAccessTokenResp, err error) {
	if len(req.AppId) == 0 {
		req.AppId = this.appId
	}
	if len(req.AppSecret) == 0 {
		req.AppSecret = this.appSecret
	}
	req.GrantType = "authorization_code"
	err = this.validator.Struct(req)
	if err != nil {
		return nil, err
	}
	params := req.toUrlValues()
	apiName := req.ApiName()
	apiUrl := fmt.Sprintf("%s%s?%s", Oauth2ApiUrl, apiName, params.Encode())
	respBody, err := internal.HttpDo(this.ctx, defaultReqTimeout, this.httpCli, "GET", apiUrl, "")
	if err != nil {
		return nil, err
	}
	resp = new(UserAccessTokenResp)
	err = json.Unmarshal(respBody, resp)
	return
}

func (this *client) UserRefreshToken(req *UserRefreshTokenReq) (resp *UserAccessTokenResp, err error) {
	if len(req.AppId) == 0 {
		req.AppId = this.appId
	}
	req.GrantType = "refresh_token"
	err = this.validator.Struct(req)
	if err != nil {
		return nil, err
	}
	params := req.toUrlValues()
	apiName := req.ApiName()
	apiUrl := fmt.Sprintf("%s%s?%s", Oauth2ApiUrl, apiName, params.Encode())
	respBody, err := internal.HttpDo(this.ctx, defaultReqTimeout, this.httpCli, "GET", apiUrl, "")
	if err != nil {
		return nil, err
	}
	resp = new(UserAccessTokenResp)
	err = json.Unmarshal(respBody, resp)
	return
}

func (this *client) VerifySign(data []byte) error {
	key := this.apiKey
	key, err := this.getSignKey("")
	if err != nil {
		return err //errors.Wrap(err, "get sandbox sign key failed :")
	}
	if err = this.verifyRespSign(data, key); err != nil {
		return err
	}
	return nil
}
