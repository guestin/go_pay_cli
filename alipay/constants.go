package alipay

const (
	sandBoxApiUrl    = "https://openapi.alipaydev.com/gateway.do"
	productionApiUrl = "https://openapi.alipay.com/gateway.do"
)

//noinspection ALL
const (
	//支付类API
	APITradeCreate    = "alipay.trade.create"
	APITradePreCreate = "alipay.trade.precreate"
	APITradePay       = "alipay.trade.pay"
	APITradeQuery     = "alipay.trade.query"
	APITradeCancel    = "alipay.trade.cancel"
	APITradeRefund    = "alipay.trade.refund"
	APITradeClose     = "alipay.trade.close"

	//资金类API
	APIFundAuthOrderVoucherCreate        = "alipay.fund.auth.order.voucher.create"
	APIFundAuthOrderFreeze               = "alipay.fund.auth.order.freeze"
	APIFundAuthOrderAppFreeze            = "alipay.fund.auth.order.app.freeze"
	APIFundAuthOrderUnFreeze             = "alipay.fund.auth.order.unfreeze"
	APIFundAuthOrderOperationCancel      = "alipay.fund.auth.operation.cancel"
	APIFundAuthOrderOperationDetailQuery = "alipay.fund.auth.operation.detail.query"
	//工具类
	APICertDownload     = "alipay.open.app.alipaycert.download"
	APIOpenOauthToken   = "alipay.open.auth.token.app"
	APISystemOauthToken = "alipay.system.oauth.token"

	//刷脸支付
	APIZolozAuthenticationCustomerFtokenQuery        = "zoloz.authentication.customer.ftoken.query"
	APIZolozAuthenticationCustomerSmilepayInitialize = "zoloz.authentication.customer.smilepay.initialize"
	APIZolozIdentificationCustomerCertifyzhubQuery   = "zoloz.identification.customer.certifyzhub.query"
	//刷脸生活
	APIZolozAuthenticationCustomerSmileLiveInitialize = "zoloz.authentication.customer.smilelive.initialize"
	//APIZolozAuthenticationCustomerFtokenQuery         = "zoloz.authentication.customer.ftoken.query"
)

//noinspection ALL
const (
	DefaultFormat  = "JSON"
	DefaultCharset = "utf-8"
	ApiVersion     = "1.0"
	longTimeFormat = "2006-01-02 15:04:05"
)

//noinspection ALL
const (
	kResponseSuffix   = "_response"
	kErrorResponse    = "error_response"
	kNullResponse     = "null_response"
	kSignNodeName     = "sign"
	kSignTypeNodeName = "sign_type"
	kCertSNNodeName   = "alipay_cert_sn"
	kCertificateEnd   = "-----END CERTIFICATE-----"
)

//noinspection ALL
const (
	TradeWaitBuyerPay = "WAIT_BUYER_PAY"
	TradeClosed       = "TRADE_CLOSED"
	TradeSuccess      = "TRADE_SUCCESS"
	TradeFinished     = "TRADE_FINISHED"

	FundStatusInit    = "INIT"
	FundStatusSuccess = "SUCCESS"
	FundStatusClosed  = "CLOSED"
)

//noinspection ALL
const (
	TradeSceneBarCode      = "bar_code"
	TradeSceneWaveCode     = "wave_code"
	TradeSceneSecurityCode = "security_code" //刷脸支付
)

//noinspection ALL
const (
	ProductCodePreAuth            = "PRE_AUTH"
	ProductCodePreAuthOnline      = "PRE_AUTH_ONLINE"
	ProductCodeOverseaInstoreAuth = "OVERSEAS_INSTORE_AUTH"
)

//noinspection ALL
const (
	AuthConfirmModeComplete    = "COMPLETE"
	AuthConfirmModeNotComplete = "NOT_COMPLETE"
)

//noinspection ALL
const (
	AuthCodeTypeBarCode = "bar_code"
)
