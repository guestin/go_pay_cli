package alipay

import (
	"net/url"
	"sort"
	"strings"
)

type Client interface {
	//使用支付宝公钥
	LoadAliPayPublicKey(alipayPublicKey string) error
	// 使用支付宝证书
	LoadCert(appPublicKeyRaw, alipayRootCertRaw, alipayPublicCertRaw []byte) error
	LoadCertFromFile(appPublicKeyFile, alipayRootCertFile, alipayPublicCertFile string) error
	LoadCertFromBase64(appPublicKey, alipayRootCert, alipayPublicCert string) error

	GetAppId() string
	SetAppAuthToken(appAuthToken string)
	GetAppAuthToken() string
	SetSellerId(sellerId string)
	GetSellerId() string
	SetISVPid(isvPid string)
	GetISVPid() string

	Execute(in Request, out Response) error
	SdkExecute(in Request, out Response) error

	// https://docs.open.alipay.com/api_1/alipay.trade.create/
	TradeCreate(req *TradeCreateReq) (resp *TradeCreateResp, err error)
	// https://docs.open.alipay.com/api_1/alipay.trade.precreate
	TradePreCreate(req *TradePreCreateReq) (resp *TradePreCreateResp, err error)
	// https://docs.open.alipay.com/api_1/alipay.trade.pay/
	TradePay(req *TradePayReq) (resp *TradePayResp, err error)
	// https://docs.open.alipay.com/api_1/alipay.trade.query
	TradeQuery(req *TradeQueryReq) (resp *TradeQueryResp, err error)
	// https://docs.open.alipay.com/api_1/alipay.trade.refund
	TradeRefund(req *TradeRefundReq) (resp *TradeRefundResp, err error)
	// https://docs.open.alipay.com/api_1/alipay.trade.cancel
	TradeCancel(req *TradeCancelReq) (resp *TradeCancelResp, err error)
	// https://docs.open.alipay.com/api_1/alipay.trade.close
	TradeClose(req *TradeCloseReq) (resp *TradeCloseResp, err error)

	//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.voucher.create
	FundAuthOrderVoucherCreate(req *FundAuthOrderVoucherCreateReq) (resp *FundAuthOrderVoucherCreateResp, err error)
	//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.freeze
	FundAuthOrderFreeze(req *FundAuthOrderFreezeReq) (resp *FundAuthOrderFreezeResp, err error)
	//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
	FundAuthOrderAppFreeze(req *FundAuthOrderAppFreezeReq) (resp *FundAuthOrderAppFreezeResp, err error)
	FundAuthOrderAppFreezeGetOrderStr(req *FundAuthOrderAppFreezeReq) (resp *FundAuthOrderAppFreezeResp, err error)
	//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.unfreeze
	FundAuthOrderUnFreeze(req *FundAuthOrderUnFreezeReq) (resp *FundAuthOrderUnFreezeResp, err error)
	//https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.cancel
	FundAuthOrderOperationCancel(req *FundAuthOrderOperationCancelReq) (resp *FundAuthOrderOperationCancelResp, err error)
	//https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.detail.query
	FundAuthOrderOperationDetailQuery(req *FundAuthOrderOperationDetailQueryReq) (resp *FundAuthOrderOperationDetailQueryResp, err error)

	//https://docs.open.alipay.com/api_9/alipay.open.app.alipaycert.download
	CertDownload(req *CertDownloadReq) (resp *CertDownloadResp, err error)
	//https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
	OpenOauthTokenApp(req *OpenOauthTokenAppReq) (resp *OpenOauthTokenAppResp, err error)
	//https://docs.open.alipay.com/api_9/alipay.system.oauth.token
	SystemOauthToken(req *SystemOauthTokenReq) (resp *SystemOauthTokenResp, err error)
	//https://opendocs.alipay.com/apis/api_46/zoloz.authentication.customer.smilepay.initialize
	ZolozAuthenticationCustomerSmilepayInitialize(req *ZolozAuthenticationCustomerSmilepayInitializeReq) (resp *ZolozAuthenticationCustomerSmilepayInitializeResp, err error)
	//https://opendocs.alipay.com/apis/api_46/zoloz.identification.customer.certifyzhub.query
	ZolozIdentificationCustomerCertifyzhubQuery(req *ZolozIdentificationCustomerCertifyzhubQueryReq) (resp *ZolozIdentificationCustomerCertifyzhubQueryResp, err error)
	//https://docs.alipay.com/pre-open/20171214171953173616/dpzxo8
	ZolozAuthenticationCustomerSmileLiveInitialize(req *ZolozAuthenticationCustomerSmileLiveInitializeReq) (resp *ZolozAuthenticationCustomerSmileLiveInitializeResp, err error)
	//https://docs.alipay.com/pre-open/20171214171953173616/pkhdn5
	//https://opendocs.alipay.com/apis/api_46/zoloz.authentication.customer.ftoken.query
	ZolozAuthenticationCustomerFtokenQuery(req *ZolozAuthenticationCustomerFtokenQueryReq) (resp *ZolozAuthenticationCustomerFtokenQueryResp, err error)
	VerifySign(param url.Values) error
}

func (this *client) SdkExecute(in Request, out Response) error {
	return this.sdkExecute(in, out)
}

func (this *client) Execute(in Request, out Response) error {
	return this.execute(in, out)
}

func (this *client) TradeCreate(req *TradeCreateReq) (resp *TradeCreateResp, err error) {
	resp = new(TradeCreateResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) TradePreCreate(req *TradePreCreateReq) (resp *TradePreCreateResp, err error) {
	resp = new(TradePreCreateResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) TradePay(req *TradePayReq) (resp *TradePayResp, err error) {
	resp = new(TradePayResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) TradeQuery(req *TradeQueryReq) (resp *TradeQueryResp, err error) {
	resp = new(TradeQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) TradeRefund(req *TradeRefundReq) (resp *TradeRefundResp, err error) {
	resp = new(TradeRefundResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) TradeCancel(req *TradeCancelReq) (resp *TradeCancelResp, err error) {
	resp = new(TradeCancelResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) TradeClose(req *TradeCloseReq) (resp *TradeCloseResp, err error) {
	resp = new(TradeCloseResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) FundAuthOrderVoucherCreate(req *FundAuthOrderVoucherCreateReq) (resp *FundAuthOrderVoucherCreateResp, err error) {
	resp = new(FundAuthOrderVoucherCreateResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) FundAuthOrderFreeze(req *FundAuthOrderFreezeReq) (resp *FundAuthOrderFreezeResp, err error) {
	resp = new(FundAuthOrderFreezeResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) FundAuthOrderAppFreeze(req *FundAuthOrderAppFreezeReq) (resp *FundAuthOrderAppFreezeResp, err error) {
	resp = new(FundAuthOrderAppFreezeResp)
	req.ProductCode = ProductCodePreAuthOnline
	err = this.Execute(req, resp)
	return
}

func (this *client) FundAuthOrderAppFreezeGetOrderStr(req *FundAuthOrderAppFreezeReq) (resp *FundAuthOrderAppFreezeResp, err error) {
	resp = new(FundAuthOrderAppFreezeResp)
	err = this.SdkExecute(req, resp)
	return
}

func (this *client) FundAuthOrderUnFreeze(req *FundAuthOrderUnFreezeReq) (resp *FundAuthOrderUnFreezeResp, err error) {
	resp = new(FundAuthOrderUnFreezeResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) FundAuthOrderOperationCancel(req *FundAuthOrderOperationCancelReq) (resp *FundAuthOrderOperationCancelResp, err error) {
	resp = new(FundAuthOrderOperationCancelResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) FundAuthOrderOperationDetailQuery(req *FundAuthOrderOperationDetailQueryReq) (resp *FundAuthOrderOperationDetailQueryResp, err error) {
	resp = new(FundAuthOrderOperationDetailQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) CertDownload(req *CertDownloadReq) (resp *CertDownloadResp, err error) {
	resp = new(CertDownloadResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) OpenOauthTokenApp(req *OpenOauthTokenAppReq) (resp *OpenOauthTokenAppResp, err error) {
	resp = new(OpenOauthTokenAppResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) SystemOauthToken(req *SystemOauthTokenReq) (resp *SystemOauthTokenResp, err error) {
	resp = new(SystemOauthTokenResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) ZolozAuthenticationCustomerSmilepayInitialize(
	req *ZolozAuthenticationCustomerSmilepayInitializeReq) (resp *ZolozAuthenticationCustomerSmilepayInitializeResp, err error) {
	resp = new(ZolozAuthenticationCustomerSmilepayInitializeResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) ZolozIdentificationCustomerCertifyzhubQuery(
	req *ZolozIdentificationCustomerCertifyzhubQueryReq) (
	resp *ZolozIdentificationCustomerCertifyzhubQueryResp, err error) {
	resp = new(ZolozIdentificationCustomerCertifyzhubQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) ZolozAuthenticationCustomerSmileLiveInitialize(
	req *ZolozAuthenticationCustomerSmileLiveInitializeReq) (
	resp *ZolozAuthenticationCustomerSmileLiveInitializeResp, err error) {
	resp = new(ZolozAuthenticationCustomerSmileLiveInitializeResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) ZolozAuthenticationCustomerFtokenQuery(
	req *ZolozAuthenticationCustomerFtokenQueryReq) (
	resp *ZolozAuthenticationCustomerFtokenQueryResp, err error) {
	resp = new(ZolozAuthenticationCustomerFtokenQueryResp)
	err = this.Execute(req, resp)
	return
}

func (this *client) VerifySign(param url.Values) (err error) {
	var certSN = param.Get(kCertSNNodeName)
	publicKey := this.alipayPublicKey
	if len(certSN) > 0 {
		publicKey, err = this.getAliPublicKey(certSN)
		if err != nil {
			return err
		}
	}
	sign := param.Get(kSignNodeName)
	var paramList = make([]string, 0, 0)
	for key := range param {
		if key == kSignNodeName || key == kSignTypeNodeName {
			continue
		}
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			paramList = append(paramList, key+"="+value)
		}
	}
	sort.Strings(paramList)
	var str = strings.Join(paramList, "&")
	return this.signProvider.Verify([]byte(str), publicKey, sign)
}
