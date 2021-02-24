package wxpay

const (
	sandBoxApiUrl    = "https://api.mch.weixin.qq.com/sandboxnew"
	productionApiUrl = "https://api.mch.weixin.qq.com"
	oauth2ApiUrl     = "https://api.weixin.qq.com/sns/oauth2"
)

const (
	CodeSuccess = "SUCCESS" //
	CodeFail    = "FAIL"
)

//goland:noinspection ALL
const (
	APIPayUnifiedOrder = "/pay/unifiedorder"
	APIPayOrderQuery   = "/pay/orderquery"
	APIPayCloseOrder   = "/pay/closeorder"
	APIPayRefund       = "/secapi/pay/refund"
	APIPayRefundQuery  = "/pay/refundquery"
	APIPayDownloadBill = "/pay/downloadbill"
	APIPayGetSignKey   = "/pay/getsignkey"
	APIPayMicroPay     = "/pay/micropay"
	APIPayReverse      = "/pay/reverse"
	APIPayitilReport   = "/payitil/report"

	APIDepositFacePay     = "/deposit/facepay"
	APIDepositMicroPay    = "/deposit/micropay"
	APIDepositOrderQuery  = "/deposit/orderquery"
	APIDepositReverse     = "/deposit/reverse"
	APIDepositConsume     = "/deposit/consume"
	APIDepositRefund      = "/deposit/refund"
	APIDepositRefundQuery = "/deposit/refundquery"

	APIAccessToken  = "/access_token"
	APIRefreshToken = "/refresh_token"
)

//noinspection ALL
const (
	TradeTypeJSAPI  = "JSAPI"
	TradeTypeNative = "NATIVE"
	TradeTypeAPP    = "APP"
)

//noinspection ALL
const (
	TradeSuccess    = "SUCCESS"
	TradeRefund     = "REFUND"
	TradeNotPay     = "NOTPAY"
	TradeClosed     = "CLOSED"
	TradeRevoked    = "REVOKED"
	TradeUserPaying = "USERPAYING"
	TradePayerror   = "PAYERROR"

	TradeSettling       = "SETTLING"       //押金消费已受理
	TradeSettlementFail = "SETTLEMENTFAIL" //押金解除冻结失败
	TradeConsumed       = "CONSUMED"       //押金消费成功

)
