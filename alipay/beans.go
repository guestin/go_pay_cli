package alipay

import (
	"fmt"
	"net/url"
)

//goland:noinspection ALL
type CodeMsg struct {
	Code    string `json:"code" validate:"required"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`

	contentRaw string `json:"-"`
}

func (this *CodeMsg) setCodeMsg(code, msg string) {
	this.Code = code
	this.Msg = msg
}

func (this *CodeMsg) GetCode() string {
	return this.Code
}

func (this *CodeMsg) GetMsg() string {
	return this.Msg
}

func (this *CodeMsg) setSubCodeMsg(subCode, subMsg string) {
	this.SubCode = subCode
	this.SubMsg = subMsg
}

func (this *CodeMsg) GetSubCode() string {
	return this.SubCode
}

func (this *CodeMsg) GetSubMsg() string {
	return this.SubMsg
}

func (this *CodeMsg) setRespContent(content string) {
	this.contentRaw = content
}

func (this *CodeMsg) GetRespContent() string {
	return this.contentRaw
}

func (this *CodeMsg) IsCodeSuccess() bool {
	return this.Code == CodeSuccess
}

func (this *CodeMsg) IsSystemError() bool {
	return this.SubCode == ACQ_SYSTEM_ERROR
}

func (this *CodeMsg) Error() string {
	e := fmt.Sprintf("%s:%s", this.Code, this.Msg)
	if len(this.SubCode) > 0 && len(this.SubMsg) > 0 {
		return fmt.Sprintf("%s(%s:%s)", e, this.SubCode, this.SubMsg)
	}
	return e
}

//region Trade

//region Trade Optional
type GoodsDetailItem struct {
	//region 必传参数
	GoodsId   string `json:"goods_id" validate:"required"`      //商品的编号
	GoodsName string `json:"goods_name" validate:"required"`    //商品名称
	Quantity  int    `json:"quantity" validate:"required,gt=0"` //商品数量
	Price     string `json:"price" validate:"required"`         //商品单价，单位为元
	//endregion
	GoodsCategory  string `json:"goods_category,omitempty"`  //商品类目
	CategoriesTree string `json:"categories_tree,omitempty"` //商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割
	Body           string `json:"body,omitempty"`            //商品描述信息
	ShowURL        string `json:"show_url,omitempty"`        //商品的展示地址
}

type BusinessParams struct {
	CampusCard      string `json:"campus_card,omitempty"`       //校园卡编号
	CardType        string `json:"card_type,omitempty"`         //卡类型
	ActualOrderTime string `json:"actual_order_time,omitempty"` //实际订单时间，在乘车码场景，传入的是用户刷码乘车时间
}

type ExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"` //系统商编号.该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的PID
	IndustryRefluxInfo   string `json:"industry_reflux_info,omitempty"`    //行业数据回流信息, 详见：地铁支付接口参数补充说明
	CardType             string `json:"card_type,omitempty"`               //卡类型
}

type SettleDetailInfo struct {
	TransInType      string  `json:"transInType" validate:"required"` //结算收款方的账户类型。cardAliasNo：结算收款方的银行卡编号; userId：表示是支付宝账号对应的支付宝唯一用户号; loginName：表示是支付宝登录号；
	TransIn          string  `json:"trans_in" validate:"required"`    //结算收款方。当结算收款方类型是cardAliasNo时，本参数为用户在支付宝绑定的卡编号；结算收款方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字；当结算收款方类型是loginName时，本参数为用户的支付宝登录号
	SummaryDimension string  `json:"summary_dimension,omitempty"`     //结算汇总维度，按照这个维度汇总成批次结算，由商户指定。
	SettleEntityId   string  `json:"settle_entity_id,omitempty"`      //结算主体标识。当结算主体类型为SecondMerchant时，为二级商户的SecondMerchantID；当结算主体类型为Store时，为门店的外标。
	SettleEntityType string  `json:"settle_entity_type,omitempty"`    //结算主体类型。	二级商户:SecondMerchant;商户或者直连商户门店:Store
	Amount           float64 `json:"amount" validate:"gt=0"`          //结算的金额，单位为元。目前必须和交易金额相同
	SettlePeriodTime string  `json:"settle_period_time,omitempty"`    //该笔订单的超期自动确认结算时间，到达期限后，将自动确认结算。此字段只在签约账期结算模式时有效。取值范围：1d～365d。d-天。 该参数数值不接受小数点。
}

type SettleInfo struct {
	SettleDetailInfos []*SettleDetailInfo `json:"settle_detail_infos" validate:"required"` //结算详细信息，json数组，目前只支持一条。
	SettlePeriodTime  string              `json:"settle_period_time,omitempty"`            //该笔订单的超期自动确认结算时间，到达期限后，将自动确认结算。此字段只在签约账期结算模式时有效。取值范围：1d～365d。d-天。 该参数数值不接受小数点。
}

type LogisticsDetail struct {
	LogisticsType string `json:"logistics_type,omitempty"` //物流类型,	POST 平邮,	EXPRESS 其他快递,	VIRTUAL 虚拟物品,	EMS EMS,	DIRECT 无需物流。
}

type ReceiverAddressInfo struct {
	Name         string `json:"name,omitempty"`          //收货人的姓名
	Address      string `json:"address,omitempty"`       //收货地址
	Mobile       string `json:"mobile,omitempty"`        //收货人手机号
	Zip          string `json:"zip,omitempty"`           //收货地址邮编
	DivisionCode string `json:"division_code,omitempty"` //中国标准城市区域码
}

//endregion Trade Optional

type BaseReq struct {
	NotifyURL string `json:"-"`
}

func (this *BaseReq) NotifyUrl() string {
	return this.NotifyURL
}

func (this *BaseReq) NeedEncrypt() bool {
	return false
}
func (this *BaseReq) HasBizContent() bool {
	return true
}
func (this *BaseReq) Params() url.Values {
	return url.Values{}
}

//https://docs.open.alipay.com/api_1/alipay.trade.create/
type TradeCreateReq struct {
	BaseReq
	OutTradeNo  string `json:"out_trade_no" validate:"required"` //商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	SellerId    string `json:"seller_id,omitempty"`              //卖家支付宝用户ID。如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount string `json:"total_amount" validate:"required"` //订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]	如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	Subject     string `json:"subject" validate:"required"`      //订单标题
	Body        string `json:"body,omitempty"`                   // 对交易或商品的描述

	BuyerId string `json:"buyer_id" validate:"required"` //买家的支付宝唯一用户号（2088开头的16位纯数字）

	ProductCode string `json:"product_code,omitempty"` //销售产品码。	如果签约的是当面付快捷版，则传OFFLINE_PAYMENT;	其它支付宝当面付产品传FACE_TO_FACE_PAYMENT；	不传默认使用FACE_TO_FACE_PAYMENT；

	GoodsDetails []*GoodsDetailItem `json:"goods_detail,omitempty" validate:"omitempty,required"` //订单包含的商品列表信息，json格式，其它说明详见：“商品明细说明”

	OperatorId string `json:"operator_id,omitempty"` //商户操作员编号
	StoreId    string `json:"store_id,omitempty"`    //商户门店编号
	TerminalId string `json:"terminal_id,omitempty"` //商户机具终端编号

	TimeoutExpress string `json:"timeout_express,omitempty"` //该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。

	ExtendParams *ExtendParams `json:"extend_params,omitempty"` //业务扩展参数

}

func (this *TradeCreateReq) ApiName() string {
	return APITradeCreate
}

type TradeCreateResp struct {
	CodeMsg
	OutTradeNo string `json:"out_trade_no"` //商户订单号
	TradeNo    string `json:"trade_no"`     // 支付宝交易号
}

// https://docs.open.alipay.com/api_1/alipay.trade.precreate
type TradePreCreateReq struct {
	BaseReq
	OutTradeNo  string `json:"out_trade_no" validate:"required"` //商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	TotalAmount string `json:"total_amount" validate:"required"` //订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]	如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	Subject     string `json:"subject" validate:"required"`      //订单标题
	Body        string `json:"body,omitempty"`                   // 对交易或商品的描述

	SellerId     string             `json:"seller_id,omitempty"`                                  //卖家支付宝用户ID。如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	GoodsDetails []*GoodsDetailItem `json:"goods_detail,omitempty" validate:"omitempty,required"` //订单包含的商品列表信息，json格式，其它说明详见：“商品明细说明”

	ProductCode string `json:"product_code,omitempty"` //销售产品码。	如果签约的是当面付快捷版，则传OFFLINE_PAYMENT;	其它支付宝当面付产品传FACE_TO_FACE_PAYMENT；	不传默认使用FACE_TO_FACE_PAYMENT；

	OperatorId string `json:"operator_id,omitempty"` //商户操作员编号
	StoreId    string `json:"store_id,omitempty"`    //商户门店编号
	TerminalId string `json:"terminal_id,omitempty"` //商户机具终端编号

	DisablePayChannels string `json:"disable_pay_channels,omitempty"` //禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔 注，与enable_pay_channels互斥 渠道列表：https://docs.open.alipay.com/common/wifww7
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`  //可用渠道，用户只能在指定渠道范围内支付	当有多个渠道时用“,”分隔	注，与disable_pay_channels互斥

	TimeoutExpress       string `json:"timeout_express,omitempty"`         //该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	QrCodeTimeoutExpress string `json:"qr_code_timeout_express,omitempty"` //该笔订单允许的最晚付款时间，逾期将关闭交易，从生成二维码开始计时。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。

	ExtendParams *ExtendParams `json:"extend_params,omitempty"` //业务扩展参数
}

func (this *TradePreCreateReq) ApiName() string {
	return APITradePreCreate
}

type TradePreCreateResp struct {
	CodeMsg
	OutTradeNo string `json:"out_trade_no"` //商户订单号
	QrCode     string `json:"qr_code"`
}

type PromoParam struct {
	ActualOrderTime string `json:"actual_order_time,omitempty"` //存在延迟扣款这一类的场景，用这个时间表明用户发生交易的时间，比如说，在公交地铁场景，用户刷码出站的时间，和商户上送交易的时间是不一样的。
}

// https://docs.open.alipay.com/api_1/alipay.trade.pay/
type TradePayReq struct {
	BaseReq
	OutTradeNo  string `json:"out_trade_no" validate:"required"` //商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	TotalAmount string `json:"total_amount" validate:"required"` //订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]	如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	Subject     string `json:"subject" validate:"required"`      //订单标题
	Body        string `json:"body,omitempty"`                   // 对交易或商品的描述

	//预授权转支付时必须
	SellerId string `json:"seller_id,omitempty"` //如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	//预授权转支付时,固定值 PRE_AUTH
	ProductCode string `json:"product_code,omitempty"` //销售产品码

	//注：预授权转支付时，填写预授权用户uid，通过预授权冻结接口返回的payer_user_id字段获取
	BuyerId string `json:"buyer_id,omitempty"` //买家的支付宝唯一用户号（2088开头的16位纯数字）

	//以下两个参数付款支付时必须，预授权转支付时非必须
	Scene    string `json:"scene,omitempty"`     //支付场景	条码支付，取值：bar_code	声波支付，取值：wave_code
	AuthCode string `json:"auth_code,omitempty"` //支付授权码，25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准

	//以下两个参数付款支付时非必须，预授权转支付时必须
	AuthNo          string `json:"auth_no,omitempty"`
	AuthConfirmMode string `json:"auth_confirm_mode,omitempty"` //预授权确认模式，授权转交易请求中传入，适用于预授权转交易业务使用，目前只支持PRE_AUTH(预授权产品码) COMPLETE：转交易支付完成结束预授权，解冻剩余金额; NOT_COMPLETE：转交易支付完成不结束预授权，不解冻剩余金

	GoodsDetails []*GoodsDetailItem `json:"goods_detail,omitempty" validate:"omitempty,required"` //订单包含的商品列表信息，json格式，其它说明详见：“商品明细说明”

	OperatorId string `json:"operator_id,omitempty"` //商户操作员编号
	StoreId    string `json:"store_id,omitempty"`    //商户门店编号
	TerminalId string `json:"terminal_id,omitempty"` //商户机具终端编号

	TimeoutExpress string `json:"timeout_express,omitempty"` //该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。

	AdvancePaymentType string `json:"advance_payment_type,omitempty"` //支付模式类型,若值为ENJOY_PAY_V2表示当前交易允许走先享后付2.0垫资
	IsAsyncPay         string `json:"is_async_pay,omitempty"`         //是否异步支付，传入true时，表明本次期望走异步支付，会先将支付请求受理下来，再异步推进。商户可以通过交易的异步通知或者轮询交易的状态来确定最终的交易结果

	ExtendParams *ExtendParams `json:"extend_params,omitempty"` //业务扩展参数
}

func (this *TradePayReq) ApiName() string {
	return APITradePay
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"` //交易使用的资金渠道，
	BankCode    string `json:"bank_code,omitempty"`    //银行卡支付时的银行代码，
	Amount      string `json:"amount,omitempty"`       //支付工具类型所使用的金额，
	RealAmount  string `json:"real_amount,omitempty"`  //渠道实际付款金额，
}

type TradePayResp struct {
	CodeMsg
	TradeNo       string           `json:"trade_no"`                 //支付宝交易号
	OutTradeNo    string           `json:"out_trade_no"`             //商户订单号
	BuyerLogonId  string           `json:"buyer_logon_id"`           //买家支付宝账号
	TotalAmount   string           `json:"total_amount"`             //交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount
	ReceiptAmount string           `json:"receipt_amount"`           //实收金额
	GmtPayment    string           `json:"gmt_payment"`              //交易支付时间
	FundBillList  []*TradeFundBill `json:"fund_bill_list,omitempty"` //交易支付使用的资金渠道
	BuyerUserId   string           `json:"buyer_user_id"`            //买家在支付宝的用户id
}

// https://docs.open.alipay.com/api_1/alipay.trade.query
type TradeQueryReq struct {
	BaseReq
	OutTradeNo   string `json:"out_trade_no,omitempty"`  //订单支付时传入的商户订单号,和支付宝交易号不能同时为空。trade_no,out_trade_no如果同时存在优先取trade_no
	TradeNo      string `json:"trade_no,omitempty"`      // 支付宝交易号，和商户订单号不能同时为空
	QueryOptions string `json:"query_options,omitempty"` //查询选项，商户通过上送该字段来定制查询返回信息
}

func (this *TradeQueryReq) ApiName() string {
	return APITradeQuery
}

type TradeQueryResp struct {
	CodeMsg
	OutTradeNo      string `json:"out_trade_no"`      //商户订单号
	TradeNo         string `json:"trade_no"`          //支付宝交易号
	BuyerLogonId    string `json:"buyer_logon_id"`    //买家支付宝账号
	TradeStatus     string `json:"trade_status"`      //交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
	TotalAmount     string `json:"total_amount"`      //交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount
	TransCurrency   string `json:"trans_currency"`    //标价币种，该参数的值为支付时传入的trans_currency，支持英镑：GBP、港币：HKD、美元：USD、新加坡元：SGD、日元：JPY、加拿大元：CAD、澳元：AUD、欧元：EUR、新西兰元：NZD、韩元：KRW、泰铢：THB、瑞士法郎：CHF、瑞典克朗：SEK、丹麦克朗：DKK、挪威克朗：NOK、马来西亚林吉特：MYR、印尼卢比：IDR、菲律宾比索：PHP、毛里求斯卢比：MUR、以色列新谢克尔：ILS、斯里兰卡卢比：LKR、俄罗斯卢布：RUB、阿联酋迪拉姆：AED、捷克克朗：CZK、南非兰特：ZAR、人民币：CNY、新台币：TWD。当trans_currency 和 settle_currency 不一致时，trans_currency支持人民币：CNY、新台币：TWD
	SettleCurrency  string `json:"settle_currency"`   //订单结算币种，对应支付接口传入的settle_currency，支持英镑：GBP、港币：HKD、美元：USD、新加坡元：SGD、日元：JPY、加拿大元：CAD、澳元：AUD、欧元：EUR、新西兰元：NZD、韩元：KRW、泰铢：THB、瑞士法郎：CHF、瑞典克朗：SEK、丹麦克朗：DKK、挪威克朗：NOK、马来西亚林吉特：MYR、印尼卢比：IDR、菲律宾比索：PHP、毛里求斯卢比：MUR、以色列新谢克尔：ILS、斯里兰卡卢比：LKR、俄罗斯卢布：RUB、阿联酋迪拉姆：AED、捷克克朗：CZK、南非兰特：ZAR
	SettleAmount    string `json:"settle_amount"`     //结算币种订单金额
	PayCurrency     string `json:"pay_currency"`      //订单支付币种
	PayAmount       string `json:"pay_amount"`        //支付币种订单金额
	SettleTransRate string `json:"settle_trans_rate"` //结算币种兑换标价币种汇率
	TransPayRate    string `json:"trans_pay_rate"`    //标价币种兑换支付币种汇率
	BuyerPayAmount  string `json:"buyer_pay_amount"`  //买家实付金额，单位为元，两位小数。该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
	PointAmount     string `json:"point_amount"`      //积分支付的金额，单位为元，两位小数。该金额代表该笔交易中用户使用积分支付的金额，比如集分宝或者支付宝实时优惠等
	InvoiceAmount   string `json:"invoice_amount"`    //交易中用户支付的可开具发票的金额，单位为元，两位小数。该金额代表该笔交易中可以给用户开具发票的金额
	SendPayDate     string `json:"send_pay_date"`     //本次交易打款给卖家的时间
	ReceiptAmount   string `json:"receipt_amount"`    //实收金额，单位为元，两位小数。该金额为本笔交易，商户账户能够实际收到的金额
	StoreId         string `json:"store_id"`          //商户门店编号
	TerminalId      string `json:"terminal_id"`       //商户机具终端编号
	StoreName       string `json:"store_name"`        //请求交易支付中的商户店铺的名称
	BuyerUserId     string `json:"buyer_user_id"`     //买家在支付宝的用户id
}

// https://docs.open.alipay.com/api_1/alipay.trade.refund
type TradeRefundReq struct {
	BaseReq
	OutTradeNo   string `json:"out_trade_no,omitempty"`             //原支付请求的商户订单号,和支付宝交易号不能同时为空
	TradeNo      string `json:"trade_no,omitempty"`                 // 支付宝交易号，和商户订单号不能同时为空
	OutRequestNo string `json:"out_request_no" validate:"required"` //标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	RefundAmount string `json:"refund_amount" validate:"required"`  //需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason string `json:"refund_reason,omitempty"`            //退款的原因说明

	RefundCurrency string `json:"refund_currency,omitempty"` //订单退款币种信息

	OperatorId string `json:"operator_id,omitempty"` //商户操作员编号
	StoreId    string `json:"store_id,omitempty"`    //商户门店编号
	TerminalId string `json:"terminal_id,omitempty"`
}

func (this *TradeRefundReq) ApiName() string {
	return APITradeRefund
}

type TradeRefundResp struct {
	CodeMsg
	TradeNo        string `json:"trade_no"`        //支付宝交易号,当发生交易关闭或交易退款时返回；
	OutTradeNo     string `json:"out_trade_no"`    //商户订单号
	BuyerLogonId   string `json:"buyer_logon_id"`  //用户的登录id
	FundChange     string `json:"fund_change"`     //本次退款是否发生了资金变化
	RefundFee      string `json:"refund_fee"`      //退款总金额
	RefundCurrency string `json:"refund_currency"` //退款币种信息
	GmtRefundPay   string `json:"gmt_refund_pay"`  //退款支付时间
	StoreName      string `json:"store_name"`      //交易在支付时候的门店名称
	BuyerUserId    string `json:"buyer_user_id"`   //买家在支付宝的用户id
}

// https://docs.open.alipay.com/api_1/alipay.trade.cancel
type TradeCancelReq struct {
	BaseReq
	OutTradeNo string `json:"out_trade_no,omitempty"` //原支付请求的商户订单号,和支付宝交易号不能同时为空
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空
}

func (this *TradeCancelReq) ApiName() string {
	return APITradeCancel
}

type TradeCancelResp struct {
	CodeMsg
	TradeNo            string `json:"trade_no"`             //支付宝交易号,当发生交易关闭或交易退款时返回；
	OutTradeNo         string `json:"out_trade_no"`         //商户订单号
	RetryFlag          string `json:"retry_flag"`           //是否需要重试
	Action             string `json:"action"`               //本次撤销触发的交易动作,接口调用成功且交易存在时返回。可能的返回值： close：交易未支付，触发关闭交易动作，无退款； refund：交易已支付，触发交易退款动作； 未返回：未查询到交易，或接口调用失败；
	GmtRefundPay       string `json:"gmt_refund_pay"`       //当撤销产生了退款时，返回退款时间；		默认不返回该信息，需与支付宝约定后配置返回；
	RefundSettlementId string `json:"refund_settlement_id"` //当撤销产生了退款时，返回的退款清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息；
}

// https://docs.open.alipay.com/api_1/alipay.trade.close
type TradeCloseReq struct {
	BaseReq
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号，和商户订单号不能同时为空
	OutTradeNo string `json:"out_trade_no,omitempty"` //原支付请求的商户订单号,和支付宝交易号不能同时为空
	OperatorId string `json:"operator_id,omitempty"`  //卖家端自定义的的操作员 ID
}

func (this *TradeCloseReq) ApiName() string {
	return APITradeClose
}

type TradeCloseResp struct {
	CodeMsg
	TradeNo    string `json:"trade_no"`     //支付宝交易号
	OutTradeNo string `json:"out_trade_no"` //创建交易传入的商户订单号
}

//endregion Trade

//region Fund

type FundAuthOrderOptional struct {
	PayeeUserId       string `json:"payee_user_id,omitempty"`       //收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeLogonId      string `json:"payee_logon_id,omitempty"`      //收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayTimeout        string `json:"pay_timeout,omitempty"`         //该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m
	ExtraParam        string `json:"extra_param,omitempty"`         //业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.间联模式必须传入二级商户ID，key为secondaryMerchantId; 2. 当面资金授权业务对应的类目，key为category，value由支付宝分配，酒店业务传 "HOTEL",若使用信用预授权，则该值必传； 3. 外部商户的门店编号，key为outStoreCode，间联场景下建议传； 4. 外部商户的门店简称，key为outStoreAlias，可选; 5.间联模式必须传入二级商户所属机构id，key为requestOrgId;6.信用服务Id，key为serviceId，信用场景下必传，具体值需要联系芝麻客服。
	TransCurrency     string `json:"trans_currency,omitempty"`      //标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	SettleCurrency    string `json:"settle_currency,omitempty"`     //商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	EnablePayChannels string `json:"enable_pay_channels,omitempty"` //商户可用该参数指定用户可使用的支付渠道，本期支持商户可支持三种支付渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。商户可设置一种支付渠道，也可设置多种支付渠道。
	IdentityParams    string `json:"identity_params,omitempty"`     //用户实名信息参数，包含：姓名+身份证号的hash值、指定用户的uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。	姓名+身份证号hash值使用SHA256摘要方式与UTF8编码,返回十六进制的字符串。	identity_hash和alipay_user_id都是可选的，如果两个都传，则会先校验identity_hash，然后校验alipay_user_id。其中identity_hash的待加密字样如"张三4566498798498498498498"
}

//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.voucher.create
type FundAuthOrderVoucherCreateReq struct {
	BaseReq
	OutOrderNo   string `json:"out_order_no" validate:"required"`   ////商户授权资金订单号，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo string `json:"out_request_no" validate:"required"` //商户本次资金操作的请求流水号，用于标示请求流水的唯一性，需要保证在商户端不重复。
	OrderTitle   string `json:"order_title" validate:"required"`    //业务订单的简单描述，如商品名称等
	Amount       string `json:"amount" validate:"required"`         //需要冻结的金额，单位为：元（人民币），精确到小数点后两位 取值范围：[0.01,100000000.00]

	ProductCode string `json:"product_code,omitempty"` //销售产品码，后续新接入预授权当面付的业务，新当面预授权产品取值PRE_AUTH，境外预授权产品取值OVERSEAS_INSTORE_AUTH。

	PayeeUserId  string `json:"payee_user_id,omitempty"`  //收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeLogonId string `json:"payee_logon_id,omitempty"` //收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。

	PayTimeout string `json:"pay_timeout,omitempty"` //该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m

	ExtraParam string `json:"extra_param,omitempty"` //业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.间联模式必须传入二级商户ID，key为secondaryMerchantId; 2. 当面资金授权业务对应的类目，key为category，value由支付宝分配，酒店业务传 "HOTEL",若使用信用预授权，则该值必传； 3. 外部商户的门店编号，key为outStoreCode，间联场景下建议传； 4. 外部商户的门店简称，key为outStoreAlias，可选; 5.间联模式必须传入二级商户所属机构id，key为requestOrgId;6.信用服务Id，key为serviceId，信用场景下必传，具体值需要联系芝麻客服。

	TransCurrency  string `json:"trans_currency,omitempty"`  //标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	SettleCurrency string `json:"settle_currency,omitempty"` //商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP

	EnablePayChannels string `json:"enable_pay_channels,omitempty"` //商户可用该参数指定用户可使用的支付渠道，本期支持商户可支持三种支付渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。商户可设置一种支付渠道，也可设置多种支付渠道。

	IdentityParams string `json:"identity_params,omitempty"` //用户实名信息参数，包含：姓名+身份证号的hash值、指定用户的uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。	姓名+身份证号hash值使用SHA256摘要方式与UTF8编码,返回十六进制的字符串。	identity_hash和alipay_user_id都是可选的，如果两个都传，则会先校验identity_hash，然后校验alipay_user_id。其中identity_hash的待加密字样如"张三4566498798498498498498"

}

func (this *FundAuthOrderVoucherCreateReq) ApiName() string {
	return APIFundAuthOrderVoucherCreate
}

type FundAuthOrderVoucherCreateResp struct {
	CodeMsg
	OutOrderNo   string `json:"out_order_no"`   //商户的授权资金订单号
	OutRequestNo string `json:"out_request_no"` //商户本次资金操作的请求流水号
	CodeType     string `json:"code_type"`      //码类型，分为 barCode：条形码 (一维码) 和 qrCode:二维码(qrCode) ； 目前发码只支持 qrCode
	CodeValue    string `json:"code_value"`     //当前发码请求生成的二维码码串，商户端可以利用二维码生成工具根据该码串值生成对应的二维码
	CodeUrl      string `json:"code_url"`       //生成的带有支付宝logo的二维码地址，如：http://mobilecodec.alipay.com/show.htm?code=aeparsv2dknkqf3018556a；商户端通过在末尾追加picSize来指定要显示的图片大小，如 显示1280大小的URL:http://mobilecodec.alipay.com/show.htm?code=aeparsv2dknkqf3018556a&picSize=1280；目前支持的大小有：256, 227, 270, 344, 430, 512, 570, 860, 1280, 1546；
}

//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.freeze
type FundAuthOrderFreezeReq struct {
	BaseReq
	AuthCode     string `json:"auth_code" validate:"required"`
	AuthCodeType string `json:"auth_code_type" validate:"required"`
	OutOrderNo   string `json:"out_order_no" validate:"required"`   //商户授权资金订单号 ,不能包含除中文、英文、数字以外的字符，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo string `json:"out_request_no" validate:"required"` //商户本次资金操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OrderTitle   string `json:"order_title" validate:"required"`    //业务订单的简单描述，如商品名称等 长度不超过100个字母或50个汉字
	Amount       string `json:"amount" validate:"required"`         //需要冻结的金额，单位为：元（人民币），精确到小数点后两位 取值范围：[0.01,100000000.00]

	ProductCode string `json:"product_code,omitempty"` //销售产品码，后续新接入预授权当面付的业务，本字段取值固定为PRE_AUTH。

	PayeeUserId  string `json:"payee_user_id,omitempty"`  //收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeLogonId string `json:"payee_logon_id,omitempty"` //收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。

	PayTimeout string `json:"pay_timeout,omitempty"` //该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m

	ExtraParam interface{} `json:"extra_param,omitempty"` //业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.间联模式必须传入二级商户ID，key为secondaryMerchantId; 2. 当面资金授权业务对应的类目，key为category，value由支付宝分配，酒店业务传 "HOTEL",若使用信用预授权，则该值必传； 3. 外部商户的门店编号，key为outStoreCode，间联场景下建议传； 4. 外部商户的门店简称，key为outStoreAlias，可选; 5.间联模式必须传入二级商户所属机构id，key为requestOrgId;6.信用服务Id，key为serviceId，信用场景下必传，具体值需要联系芝麻客服。
	SceneCode  string      `json:"scene_code,omitempty"`
}

func (this *FundAuthOrderFreezeReq) ApiName() string {
	return APIFundAuthOrderFreeze
}

type FundAuthOrderFreezeResp struct {
	CodeMsg
	AuthNo       string `json:"auth_no"`        //支付宝的资金授权订单号
	OutOrderNo   string `json:"out_order_no"`   //商户的授权资金订单号
	OperationId  string `json:"operation_id"`   //支付宝的资金操作流水号
	OutRequestNo string `json:"out_request_no"` //商户本次资金操作的请求流水号
	Amount       string `json:"amount"`         //本次操作冻结的金额，单位为：元（人民币），精确到小数点后两位
	Status       string `json:"status"`         //资金预授权明细的状态	目前支持： INIT：初始	SUCCESS: 成功	CLOSED：关闭
	PayerUserId  string `json:"payer_user_id"`  //付款方支付宝用户号
	PayerLogonId string `json:"payer_logon_id"` //收款方支付宝账号（Email或手机号）
	GmtTrans     string `json:"gmt_trans"`      //资金授权成功时间	格式：YYYY-MM-DD HH:MM:SS
}

//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
type FundAuthOrderAppFreezeReq struct {
	BaseReq
	OutOrderNo   string `json:"out_order_no" validate:"required"`   //商户授权资金订单号 ,不能包含除中文、英文、数字以外的字符，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo string `json:"out_request_no" validate:"required"` //商户本次资金操作的请求流水号，用于标示请求流水的唯一性，需要保证在商户端不重复。
	OrderTitle   string `json:"order_title" validate:"required"`    //业务订单的简单描述，如商品名称等
	Amount       string `json:"amount" validate:"required"`         //需要冻结的金额，单位为：元（人民币），精确到小数点后两位 取值范围：[0.01,100000000.00]

	ProductCode string `json:"product_code"` //销售产品码，后续新接入预授权当面付的业务，新当面预授权产品取值PRE_AUTH，境外预授权产品取值OVERSEAS_INSTORE_AUTH。

	PayeeUserId  string `json:"payee_user_id,omitempty"`  //收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeLogonId string `json:"payee_logon_id,omitempty"` //收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。

	PayTimeout string `json:"pay_timeout,omitempty"` //该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m

	ExtraParam string `json:"extra_param,omitempty"` //业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.间联模式必须传入二级商户ID，key为secondaryMerchantId; 2. 当面资金授权业务对应的类目，key为category，value由支付宝分配，酒店业务传 "HOTEL",若使用信用预授权，则该值必传； 3. 外部商户的门店编号，key为outStoreCode，间联场景下建议传； 4. 外部商户的门店简称，key为outStoreAlias，可选; 5.间联模式必须传入二级商户所属机构id，key为requestOrgId;6.信用服务Id，key为serviceId，信用场景下必传，具体值需要联系芝麻客服。

	SceneCode string `json:"scene_code,omitempty"` //场景码，用于区分预授权不同业务场景。如：当面预授权通用场景（O2O_AUTH_COMMON_SCENE）、支付宝预授权通用场景（ONLINE_AUTH_COMMON_SCENE）、境外当面预授权通用场景（OVERSEAS_O2O_AUTH_COMMON_SCENE）、境外支付预授权通用场景（OVERSEAS_ONLINE_AUTH_COMMON_SCENE）等

	TransCurrency  string `json:"trans_currency,omitempty"`  //标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	SettleCurrency string `json:"settle_currency,omitempty"` //商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP

	EnablePayChannels string `json:"enable_pay_channels,omitempty"` //商户可用该参数指定用户可使用的支付渠道，本期支持商户可支持三种支付渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。商户可设置一种支付渠道，也可设置多种支付渠道。

	IdentityParams string `json:"identity_params,omitempty"` //用户实名信息参数，包含：姓名+身份证号的hash值、指定用户的uid。商户传入用户实名信息参数，支付宝会对比用户在支付宝端的实名信息。	姓名+身份证号hash值使用SHA256摘要方式与UTF8编码,返回十六进制的字符串。	identity_hash和alipay_user_id都是可选的，如果两个都传，则会先校验identity_hash，然后校验alipay_user_id。其中identity_hash的待加密字样如"张三4566498798498498498498"
}

func (this *FundAuthOrderAppFreezeReq) ApiName() string {
	return APIFundAuthOrderAppFreeze
}

type FundAuthOrderAppFreezeResp struct {
	CodeMsg
	AuthNo        string `json:"auth_no"`        //支付宝的资金授权订单号
	OutOrderNo    string `json:"out_order_no"`   //商户的授权资金订单号
	OperationId   string `json:"operation_id"`   //支付宝的资金操作流水号
	OutRequestNo  string `json:"out_request_no"` //商户本次资金操作的请求流水号
	Amount        string `json:"amount" `        //本次操作冻结的金额，单位为：元（人民币），精确到小数点后两位
	Status        string `json:"status"`         //资金预授权明细的状态	目前支持： INIT：初始	SUCCESS: 成功	CLOSED：关闭
	PayerUserId   string `json:"payer_user_id"`  //付款方支付宝用户号
	GmtTrans      string `json:"gmt_trans"`      //资金授权成功时间	格式：YYYY-MM-DD HH:MM:SS
	PreAuthType   string `json:"pre_auth_type"`  //预授权类型，目前支持 CREDIT_AUTH(信用预授权);	商户可根据该标识来判断该笔预授权的类型，当返回值为"CREDIT_AUTH"表明该笔预授权为信用预授权，没有真实冻结资金；当返回值为空或者不为"CREDIT_AUTH"则表明该笔预授权为普通资金预授权，会冻结用户资金。
	CreditAmount  string `json:"credit_amount"`  //本次解冻操作中信用解冻金额，单位为：元（人民币），精确到小数点后两位
	FundAmount    string `json:"fund_amount"`    //本次解冻操作中自有资金解冻金额，单位为：元（人民币），精确到小数点后两位
	TransCurrency string `json:"trans_currency"` //标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP, 人民币：CNY
}

//https://docs.open.alipay.com/api_28/alipay.fund.auth.order.unfreeze
type FundAuthOrderUnFreezeReq struct {
	BaseReq
	AuthNo       string `json:"auth_no" validate:"required"`        //支付支付宝资金授权订单号
	OutRequestNo string `json:"out_request_no" validate:"required"` //商户本次资金操作的请求流水号，同一商户每次不同的资金操作请求，商户请求流水号不能重复
	Amount       string `json:"amount" validate:"required"`         //本次操作解冻的金额，单位为：元（人民币），精确到小数点后两位，取值范围：[0.01,100000000.00]
	Remark       string `json:"remark" validate:"required"`         //商户对本次撤销操作的附言描述
	ExtraParam   string `json:"extra_param,omitempty"`              //解冻扩展信息，json格式；unfreezeBizInfo 目前为芝麻消费字段，支持Key值如下：	"bizComplete":"true" -- 选填：标识本次解冻用户是否履约，如果true信用单会完结为COMPLETE
}

func (this *FundAuthOrderUnFreezeReq) ApiName() string {
	return APIFundAuthOrderUnFreeze
}

type FundAuthOrderUnFreezeResp struct {
	CodeMsg
	AuthNo       string `json:"auth_no"`        //支付宝的资金授权订单号
	OutOrderNo   string `json:"out_order_no"`   //商户的授权资金订单号
	OperationId  string `json:"operation_id"`   //支付宝的资金操作流水号
	OutRequestNo string `json:"out_request_no"` //商户本次资金操作的请求流水号
	Amount       string `json:"amount"`         //本次操作冻结的金额，单位为：元（人民币），精确到小数点后两位
	Status       string `json:"status"`         //资金操作流水的状态	目前支持： INIT：初始	SUCCESS：成功	CLOSED：关闭
	GmtTrans     string `json:"gmt_trans"`      //资金授权成功时间	格式：YYYY-MM-DD HH:MM:SS
	CreditAmount string `json:"credit_amount"`  //本次解冻操作中信用解冻金额，单位为：元（人民币），精确到小数点后两位
	FundAmount   string `json:"fund_amount"`    //本次解冻操作中自有资金解冻金额，单位为：元（人民币），精确到小数点后两位
}

//https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.cancel
type FundAuthOrderOperationCancelReq struct {
	BaseReq
	AuthNo       string `json:"auth_no,omitempty"`          //支付宝授权资金订单号，与商户的授权资金订单号不能同时为空，二者都存在时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。
	OutOrderNo   string `json:"out_order_no,omitempty"`     //商户的授权资金订单号，与支付宝的授权资金订单号不能同时为空，二者都存在时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。
	OperationId  string `json:"operation_id,omitempty"`     //支付宝的授权资金操作流水号，与商户的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。
	OutRequestNo string `json:"out_request_no,omitempty"`   //商户的授权资金操作流水号，与支付宝的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。
	Remark       string `json:"remark" validate:"required"` //商户对本次撤销操作的附言描述
}

func (this *FundAuthOrderOperationCancelReq) ApiName() string {
	return APIFundAuthOrderOperationCancel
}

type FundAuthOrderOperationCancelResp struct {
	CodeMsg
	AuthNo       string `json:"auth_no"`        //支付宝的资金授权订单号
	OutOrderNo   string `json:"out_order_no"`   //商户的授权资金订单号
	OperationId  string `json:"operation_id"`   //支付宝的资金操作流水号
	OutRequestNo string `json:"out_request_no"` //商户本次资金操作的请求流水号
	Action       string `json:"action"`         //本次撤销触发的资金动作	close：关闭冻结明细，无资金解冻	unfreeze：产生了资金解冻
}

//https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.detail.query
type FundAuthOrderOperationDetailQueryReq struct {
	BaseReq
	AuthNo       string `json:"auth_no,omitempty"`        //支付宝授权资金订单号，与商户的授权资金订单号不能同时为空，二者都存在时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。
	OutOrderNo   string `json:"out_order_no,omitempty"`   //商户的授权资金订单号，与支付宝的授权资金订单号不能同时为空，二者都存在时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。
	OperationId  string `json:"operation_id,omitempty"`   //支付宝的授权资金操作流水号，与商户的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。
	OutRequestNo string `json:"out_request_no,omitempty"` //商户的授权资金操作流水号，与支付宝的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。
}

func (this *FundAuthOrderOperationDetailQueryReq) ApiName() string {
	return APIFundAuthOrderOperationDetailQuery
}

type FundAuthOrderOperationDetailQueryResp struct {
	CodeMsg
	AuthNo                  string `json:"auth_no"`                    //支付宝的资金授权订单号
	OutOrderNo              string `json:"out_order_no"`               //商户的授权资金订单号
	TotalFreezeAmount       string `json:"total_freeze_amount"`        //订单累计的冻结金额，单位为：元（人民币）
	RestAmount              string `json:"rest_amount"`                //订单总共剩余的冻结金额，单位为：元（人民币）
	TotalPayAmount          string `json:"total_pay_amount"`           //订单累计用于支付的金额，单位为：元（人民币）
	OrderTitle              string `json:"order_title" `               //业务订单的简单描述，如商品名称等
	PayerLogonId            string `json:"payer_logon_id"`             //付款方支付宝账号（Email或手机号），仅作展示使用，默认会加“*”号处理
	PayerUserId             string `json:"payer_user_id"`              //付款方支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成
	ExtraParam              string `json:"extra_param"`                //商户请求创建预授权订单时传入的扩展参数，仅返回商户自定义的扩展信息（merchantExt）
	OperationId             string `json:"operation_id"`               //支付宝资金操作流水号
	OutRequestNo            string `json:"out_request_no"`             //商户资金操作的请求流水号
	Amount                  string `json:"amount"`                     //该笔资金操作流水opertion_id对应的操作金额，单位为：元（人民币）
	OperationType           string `json:"operation_type"`             //支付宝资金操作类型，	目前支持：	FREEZE：冻结	UNFREEZE：解冻	PAY：支付
	Status                  string `json:"status"`                     //资金操作流水的状态	目前支持： INIT：初始	SUCCESS：成功	CLOSED：关闭
	Remark                  string `json:"remark"`                     //商户对本次操作的附言描述，长度不超过100个字母或50个汉字
	GmtCreate               string `json:"gmt_create"`                 //资金授权单据操作流水创建时间，	格式：YYYY-MM-DD HH:MM:SS
	GmtTrans                string `json:"gmt_trans"`                  //资金授权成功时间	格式：YYYY-MM-DD HH:MM:SS
	PreAuthType             string `json:"pre_auth_type"`              //预授权类型，目前支持 CREDIT_AUTH(信用预授权);	商户可根据该标识来判断该笔预授权的类型，当返回值为"CREDIT_AUTH"表明该笔预授权为信用预授权，没有真实冻结资金；当返回值为空或者不为"CREDIT_AUTH"则表明该笔预授权为普通资金预授权，会冻结用户资金。
	TransCurrency           string `json:"trans_currency"`             //标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP, 人民币：CNY
	TotalFreezeCreditAmount string `json:"total_freeze_credit_amount"` //累计冻结信用金额，单位为：元（人民币），精确到小数点后两位
	TotalFreezeFundAmount   string `json:"total_freeze_fund_amount"`   //累计冻结自有资金金额，单位为：元（人民币），精确到小数点后两位
	TotalPayCreditAmount    string `json:"total_pay_credit_amount"`    //累计支付信用金额，单位为：元（人民币），精确到小数点后两位
	TotalPayFundAmount      string `json:"total_pay_fund_amount"`      //累计支付自有资金金额，单位为：元（人民币），精确到小数点后两位
	RestCreditAmount        string `json:"rest_credit_amount"`         //剩余冻结信用金额，单位为：元（人民币），精确到小数点后两位
	RestFundAmount          string `json:"rest_fund_amount"`           //剩余冻结自有资金金额，单位为：元（人民币），精确到小数点后两位
	CreditAmount            string `json:"credit_amount"`              //该笔资金操作流水opertion_id对应的操作信用金额
	FundAmount              string `json:"fund_amount"`                //该笔资金操作流水opertion_id对应的操作自有资金金额
}

//endregion
//https://docs.open.alipay.com/api_9/alipay.open.app.alipaycert.download
type CertDownloadReq struct {
	BaseReq
	AlipayCertSn string `json:"alipay_cert_sn" validate:"required"` //支付宝公钥证书序列号
}

func (this *CertDownloadReq) ApiName() string {
	return APICertDownload
}

type CertDownloadResp struct {
	CodeMsg
	AlipayCertContent string `json:"alipay_cert_content"` //公钥证书Base64后的字符串
}

//https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
type OpenOauthTokenAppReq struct {
	BaseReq
	GrantType    string `json:"grant_type" validate:"required"` //值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code,omitempty"`                 //授权码，用户对应用授权后得到。
	RefreshToken string `json:"refresh_token,omitempty"`        //刷刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
}

func (this *OpenOauthTokenAppReq) ApiName() string {
	return APIOpenOauthToken
}

type OpenOauthTokenAppResp struct {
	CodeMsg
	UserId          string `json:"user_id"`           //支授权商户的user_id
	AuthAppId       string `json:"auth_app_id"`       //授权商户的appid
	AppAuthToken    string `json:"app_auth_token"`    //应用授权令牌
	AppRefreshToken string `json:"app_refresh_token"` //刷新令牌
	ExpiresIn       int64  `json:"expires_in"`        //应用授权令牌的有效时间（从接口调用时间作为起始时间），单位到秒
	ReExpiresIn     int64  `json:"re_expires_in"`     //刷新令牌的有效时间（从接口调用时间作为起始时间），单位到秒
}

func (this *OpenOauthTokenAppResp) IsCodeSuccess() bool {
	return this.Code == CodeSuccess || this.Code == ""
}

//https://docs.open.alipay.com/api_9/alipay.system.oauth.token
type SystemOauthTokenReq struct {
	BaseReq
	GrantType    string `json:"grant_type" validate:"required"` //值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code,omitempty"`                 //授权码，用户对应用授权后得到。
	RefreshToken string `json:"refresh_token,omitempty"`        //刷刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
}

func (this *SystemOauthTokenReq) ApiName() string {
	return APISystemOauthToken
}

func (this *SystemOauthTokenReq) HasBizContent() bool {
	return false
}

func (this *SystemOauthTokenReq) Params() url.Values {
	p := url.Values{}
	p.Set("grant_type", this.GrantType)
	p.Set("code", this.Code)
	p.Set("refresh_token", this.RefreshToken)
	return p
}

type SystemOauthTokenResp struct {
	CodeMsg
	UserId       string `json:"user_id"`       //支付宝用户的唯一userId
	AccessToken  string `json:"access_token"`  //访问令牌。通过该令牌调用需要授权类接口
	ExpiresIn    int64  `json:"expires_in"`    //访问令牌的有效时间，单位是秒。
	RefreshToken string `json:"refresh_token"` //刷新令牌。通过该令牌可以刷新access_token
	ReExpiresIn  int64  `json:"re_expires_in"` //刷新令牌的有效时间，单位是秒。
}

func (this *SystemOauthTokenResp) IsCodeSuccess() bool {
	return this.Code == CodeSuccess || this.Code == ""
}

//https://docs.open.alipay.com/194/103296/
type TradeNotification struct {
	AppId string `json:"app_id" form:"app_id" query:"app_id"` // 开发者的app_id

	NotifyTime string `json:"notify_time" form:"notify_time" query:"notify_time"` // 通知时间 通知的发送时间。格式为yyyy-MM-dd HH:mm:ss
	NotifyType string `json:"notify_type" form:"notify_type" query:"notify_type"` // 通知类型 trade_status_sync
	NotifyId   string `json:"notify_id" form:"notify_id" query:"notify_id"`       // 通知校验ID
	SignType   string `json:"sign_type" form:"sign_type" query:"sign_type"`       // 签名类型,户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2（如果开发者手动验签，不使用 SDK 验签，可以不传此参数）
	Sign       string `json:"sign" form:"sign" query:"sign"`                      // 签名 请参考异步返回结果的验签（如果开发者手动验签，不使用 SDK 验签，可以不传此参数）

	TradeNo    string `json:"trade_no" form:"trade_no" query:"trade_no"`             // 支付宝交易号 支付宝交易凭证号
	OutTradeNo string `json:"out_trade_no" form:"out_trade_no" query:"out_trade_no"` // 商户订单号 原支付请求的商户订单号
	OutBizNo   string `json:"out_biz_no" form:"out_biz_no" query:"out_biz_no"`       // 商户业务号 商户业务ID，主要是退款通知中返回退款申请的流水号

	BuyerId      string `json:"buyer_id" form:"buyer_id" query:"buyer_id"`                   // 买家支付宝用户号
	BuyerLogonId string `json:"buyer_logon_id" form:"buyer_logon_id" query:"buyer_logon_id"` // 买家支付宝账号
	SellerId     string `json:"seller_id" form:"seller_id" query:"seller_id"`                // 卖家支付宝用户号
	SellerEmail  string `json:"seller_email" form:"seller_email" query:"seller_email"`       // 卖家支付宝账号

	TradeStatus string `json:"trade_status" form:"trade_status" query:"trade_status"` // 交易状态

	TotalAmount    string `json:"total_amount" form:"total_amount" query:"total_amount"`             // 订单金额 本次交易支付的订单金额，单位为人民币（元）
	ReceiptAmount  string `json:"receipt_amount" form:"receipt_amount" query:"receipt_amount"`       // 实收金额 商家在交易中实际收到的款项，单位为元
	InvoiceAmount  string `json:"invoice_amount" form:"invoice_amount" query:"invoice_amount"`       // 开票金额 用户在交易中支付的可开发票的金额
	BuyerPayAmount string `json:"buyer_pay_amount" form:"buyer_pay_amount" query:"buyer_pay_amount"` // 付款金额 用户在交易中支付的金额
	PointAmount    string `json:"point_amount" form:"point_amount" query:"point_amount"`             // 集分宝金额 使用集分宝支付的金额
	RefundFee      string `json:"refund_fee" form:"refund_fee" query:"refund_fee"`                   // 总退款金额 退款通知中，返回总退款金额，单位为元，支持两位小数
	SendBackFee    string `json:"send_back_fee" form:"send_back_fee" query:"send_back_fee"`          //实际退款金额 商户实际退款给用户的金额，单位为元，支持两位小数

	Subject string `json:"subject" form:"subject" query:"subject"` // 订单标题 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来
	Body    string `json:"body" form:"body" query:"body"`          // 商品描述 该订单的备注、描述、明细等。对应请求时的body参数，原样通知回来

	GmtCreate    string `json:"gmt_create" form:"gmt_create" query:"gmt_create"`             // 交易创建时间 该笔交易创建的时间。格式为yyyy-MM-dd HH:mm:ss
	GmtPayment   string `json:"gmt_payment" form:"gmt_payment" query:"gmt_payment"`          // 交易付款时间 该笔交易的买家付款时间。格式为yyyy-MM-dd HH:mm:ss
	GmtRefund    string `json:"gmt_refund" form:"gmt_refund" query:"gmt_refund"`             // 交易退款时间 该笔交易的退款时间。格式为yyyy-MM-dd HH:mm:ss.S
	GmtClose     string `json:"gmt_close" form:"gmt_close" query:"gmt_close"`                // 交易结束时间 该笔交易结束时间。格式为yyyy-MM-dd HH:mm:ss
	FundBillList string `json:"fund_bill_list" form:"fund_bill_list" query:"fund_bill_list"` // 支付金额信息 支付成功的各个渠道金额信息，详见资金明细信息说明

	//region  (预授权通知字段)

	AuthNo       string `json:"auth_no" form:"auth_no" query:"auth_no"`                      //支付宝资金授权订单号
	OutOrderNo   string `json:"out_order_no" form:"out_order_no" query:"out_trade_no"`       // 商户的授权资金订单号
	OutRequestNo string `json:"out_request_no" form:"out_request_no" query:"out_request_no"` //商户资金操作的请求流水号
	OperationId  string `json:"operation_id" form:"operation_id" query:"operation_id"`       //支付宝资金操作流水号

	OperationType string `json:"operation_type" form:"operation_type" query:"operation_type"` //支付宝资金操作类型， 目前支持： FREEZE：冻结 UNFREEZE：解冻 PAY：支付

	Status string `json:"status" form:"status" query:"status"` // 资金操作流水的状态，	目前支持：INIT：初始;SUCCESS：成功;CLOSED：关闭

	Amount string `json:"amount" form:"amount" query:"amount"` //该笔资金操作流水opertion_id对应的操作金额，单位为：元（人民币）

	TotalFreezeAmount string `json:"total_freeze_amount" form:"total_freeze_amount" query:"total_freeze_amount"` //订单累计的冻结金额，单位为：元（人民币）
	RestAmount        string `json:"rest_amount" form:"rest_amount" query:"rest_amount"`                         //订单总共剩余的冻结金额，单位为：元（人民币）
	TotalPayAmount    string `json:"total_pay_amount" form:"total_pay_amount" query:"total_pay_amount"`          //订单累计用于支付的金额，单位为：元（人民币）
	PayerUserId       string `json:"payer_user_id" form:"payer_user_id" query:"payer_user_id"`                   // 付款方支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成
	PayerLogonId      string `json:"payer_logon_id" form:"payer_logon_id" query:"payer_logon_id"`                // 付款方支付宝账号（Email或手机号），仅作展示使用，默认会加“*”号处理

	//endregion
}
type FaceExtInfo struct {
	QueryType string `json:"query_type"` //query_type不填, 返回uid	query_type=1, 返回手机号 query_type=2, 返回图片
}

//https://opendocs.alipay.com/apis/api_46/zoloz.authentication.customer.ftoken.query
type ZolozAuthenticationCustomerFtokenQueryReq struct {
	BaseReq
	Ftoken  string       `json:"ftoken" validate:"required"`   //人脸token
	BizType string       `json:"biz_type" validate:"required"` //1、1：1人脸验证能力	2、1：n人脸搜索能力（支付宝uid入库） 3、1：n人脸搜索能力（支付宝手机号入库） 4、手机号和人脸识别综合能力
	ExtInfo *FaceExtInfo `json:"ext_info,omitempty"`           //人脸产品拓展参数
}

type ZhubUidTelPair struct {
	UserId string `json:"user_id"` //支付宝uid
	Phone  string `json:"phone"`   //	手机号
}

type ZolozAuthenticationCustomerFtokenQueryResp struct {
	CodeMsg
	Uid            string            `json:"uid"`               //支付宝uid
	AuthimgBase64  string            `json:"authimg_base_64"`   //图片base64 encodeString
	UidTelPairList []*ZhubUidTelPair `json:"uid_tel_pair_list"` //用户名和手机号信息返回的列表
}

func (this *ZolozAuthenticationCustomerFtokenQueryReq) ApiName() string {
	return APIZolozAuthenticationCustomerFtokenQuery
}

//https://opendocs.alipay.com/apis/api_46/zoloz.authentication.customer.smilepay.initialize
type ZolozAuthenticationCustomerSmileLiveInitializeReq struct {
	BaseReq
	Zimmetainfo string `json:"zimmetainfo" validate:"required"` //初始化入参,在zolozGetMetaInfo接口返回的metainfo对象中extInfo子对象中加入以下参数
	/*
		参数名称		是否必填	参数说明
		bizType		是		1，表明1：1身份核验场景（姓名＋身份证号）	2，表明1：N刷脸搜索， 支付宝UID入库［暂未上线］	3，表明1：N刷脸搜索， 支付宝手机号入库［暂未上线］	4，表明1：1身份核验场景（手机号）
		certType	否		IDCARD，表明身份证；当bizType=1时，必填
		certNo		否		用户的身份证号；当bizType=1时，必填
		certName	否		用户的姓名（保持和身份证一致）；当bizType=1时，必填
	*/
}

func (this *ZolozAuthenticationCustomerSmileLiveInitializeReq) ApiName() string {
	return APIZolozAuthenticationCustomerSmileLiveInitialize
}

type ZolozAuthenticationCustomerSmileLiveInitializeResp struct {
	CodeMsg
	Result string `json:"result"` //result对象结构如下：
	/*
	   键值					描述
	   retCode				刷脸返回码
	   retMessage			刷脸返回消息
	   zimId				刷脸调用的标识，将作为下一步zolozVerify接口的入參
	   zimInitClientData	刷脸的下发协议数据，将作为下一步zolozVerify接口的入參
	*/
}

//https://opendocs.alipay.com/apis/api_46/zoloz.authentication.customer.smilepay.initialize
type ZolozAuthenticationCustomerSmilepayInitializeReq struct {
	BaseReq
	Zimmetainfo string `json:"zimmetainfo" validate:"required"` //{ "apdidToken": "设备指纹", "appName": "应用名称", "appVersion": "应用版本", "bioMetaInfo": "生物信息如2.3.0:3,-4" }
}

func (this *ZolozAuthenticationCustomerSmilepayInitializeReq) ApiName() string {
	return APIZolozAuthenticationCustomerSmilepayInitialize
}

type ZolozAuthenticationCustomerSmilepayInitializeResp struct {
	CodeMsg
	Result string `json:"result"` //{"retCode":"","retMessage":"","result":{"zimId":" 唤人脸参数","type":"zimInit","zimInitClientData":"下发参数"}}
}

//https://opendocs.alipay.com/apis/api_46/zoloz.identification.customer.certifyzhub.query
//人脸服务的结果查询(一体化)
type ZolozIdentificationCustomerCertifyzhubQueryReq struct {
	BaseReq
	BizId    string `json:"biz_id" validate:"required"`       //业务单据号，用于核对和排查
	ZimId    string `json:"zim_id" validate:"required"`       //zimId，用于查询认证结果
	FaceType int    `json:"face_type" validate:"oneof=0 1 2"` //0：匿名注册 1：匿名认证 2：实名认证
	NeedImg  bool   `json:"need_img"`                         //是否需要返回人脸图片
}

//
func (this *ZolozIdentificationCustomerCertifyzhubQueryReq) ApiName() string {
	return APIZolozIdentificationCustomerCertifyzhubQuery
}

type FaceAttrInfo struct {
	Rect string `json:"rect"` //left,top,width,height 人脸图片中的人脸框的左上点和宽高，图片内坐标，无需脱敏
}

type ZolozIdentificationCustomerCertifyzhubQueryResp struct {
	CodeMsg
	BizId        string       `json:"biz_id"`         //业务单据号，用于核对和排查
	ZimCode      string       `json:"zim_code"`       //人脸服务端返回码
	ZimMsg       string       `json:"zim_msg"`        //人脸服务端返回信息
	ImgStr       string       `json:"img_str"`        //图片字节数组进行Base64编码后的字符串
	FaceAttrInfo FaceAttrInfo `json:"face_attr_info"` //	人脸属性信息，提供对人脸base64图片的额外描述，包括不限于人脸矩形框。目前仅为矩形框，无需脱敏。
}
