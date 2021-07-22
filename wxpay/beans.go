package wxpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type BaseReq struct {
	NotifyURL string `xml:"-" json:"-"` //接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
}

func (this *BaseReq) NotifyUrl() string {
	return this.NotifyURL
}

type CodeMsg struct {
	ReturnCode string `xml:"return_code" json:"return_code"`   //返回状态码: SUCCESS/FAIL	此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg" json:"return_msg"`     //返回信息: 当return_code为FAIL时返回信息为错误原因 ，例如	签名失败	参数格式校验错误
	ResultCode string `xml:"result_code" json:"result_code"`   //业务结果: SUCCESS/FAIL
	ErrCode    string `xml:"err_code" json:"err_code"`         //错误代码: 当result_code为FAIL时返回错误代码，详细参见下文错误列表
	ErrCodeDes string `xml:"err_code_des" json:"err_code_des"` //错误代码描述: 当result_code为FAIL时返回错误描述，详细参见下文错误列表
}

type BaseResp struct {
	AppId    string `xml:"appid" json:"appid"`           //公众账号ID/服务商的APPID: 调用接口提交的公众账号ID
	MchId    string `xml:"mch_id" json:"mch_id"`         //商户号: 调用接口提交的商户号
	SubAppId string `xml:"sub_appid" json:"sub_app_id"`  //子商户公众账号ID,微信分配的子商户公众账号ID
	SubMchId string `xml:"sub_mch_id" json:"sub_mch_id"` //子商户号: 微信支付分配的子商户号
	NonceStr string `xml:"nonce_str" json:"nonce_str"`   //随机字符串: 微信返回的随机字符串
	Sign     string `xml:"sign" json:"sign"`             //签名: 微信返回的签名值，详见签名算法
}

type UserOpenIdRespInfo struct {
	Openid         string `xml:"openid" json:"openid"`                     //用户标识: 用户在商户appid下的唯一标识
	IsSubscribe    string `xml:"is_subscribe" json:"is_subscribe"`         //是否关注公众账号: 用户是否关注公众账号，Y-关注，N-未关注
	SubOpenid      string `xml:"sub_openid" json:"sub_openid"`             //用户子标识: 用户在子商户appid下的唯一标识
	SubIsSubscribe string `xml:"sub_is_subscribe" json:"sub_is_subscribe"` //是否关注子公众账号: 用户是否关注子公众账号，Y-关注，N-未关注（机构商户不返回）
}

func (this *CodeMsg) IsSuccess() bool {
	return this.ReturnCode == CodeSuccess
}

func (this *CodeMsg) IsBusinessSuccess() bool {
	return this.ReturnCode == CodeSuccess && this.ResultCode == CodeSuccess
}

func (this *CodeMsg) IsSystemError() bool {
	return this.ErrCode == SYSTEMERROR
}

func (this *CodeMsg) Error() string {
	e := fmt.Sprintf("[%s] %s", this.ReturnCode, this.ReturnMsg)
	if len(this.ErrCode) > 0 && len(this.ErrCodeDes) > 0 {
		return fmt.Sprintf("%s([%s] %s:%s)", e, this.ResultCode, this.ErrCode, this.ErrCodeDes)
	}
	return e
}

type GoodsDetail struct {
	CostPrice   *int               `json:"cost_price,omitempty"`                  //1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。 2.当订单原价与支付金额不相等，则不享受优惠。 3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	ReceiptId   string             `json:"receipt_id,omitempty"`                  //商家小票ID
	GoodsDetail []*GoodsDetailItem `json:"goods_detail" validate:"required,dive"` //单品信息，使用Json数组格式提交
}

type GoodsDetailItem struct {
	GoodsId      string `json:"goods_id" validate:"required"`      // 是 商品编码 :由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	WxpayGoodsId string `json:"wxpay_goods_id"`                    // 否 微信支付定义的统一商品编号（没有可不传）
	GoodsName    string `json:"goods_name" validate:"required"`    // 否 商品的实际名称
	Quantity     int    `json:"quantity" validate:"required,gt=0"` // 是 商品数量
	Price        int64  `json:"price" validate:"required,gt=0"`    // 是 商品单价，单位为：分。如果商户有优惠，需传输商户优惠后的单价(例如：用户对一笔100元的订单使用了商场发的纸质优惠券100-50，则活动商品的单价应为原单价-50)
}

type SceneInfo struct {
	Id       string `json:"id,omitempty"`        // 否 门店编号，由商户自定义
	Name     string `json:"name,omitempty"`      // 否 门店名称 ，由商户自定义
	AreaCode string `json:"area_code,omitempty"` // 否 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	Address  string `json:"address,omitempty"`   // 否 门店详细地址 ，由商户自定义
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=23_1
type payGetSignKeyReq struct {
	BaseReq
}

func (this *payGetSignKeyReq) ApiName() string {
	return APIPayGetSignKey
}

func (this *payGetSignKeyReq) toUrlValues() url.Values {
	return url.Values{}
}

type payGetSignKeyResp struct {
	CodeMsg
	MchId          string `xml:"mch_id" json:"mch_id"`                   //微信支付分配的微信商户号
	SandboxSignkey string `xml:"sandbox_signkey" json:"sandbox_signkey"` //返回的沙箱密钥
}

type PayOrderOptional struct {
	ProductId    string     `xml:"product_id,omitempty" json:"product_id,omitempty"` // trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	LimitPay     string     `xml:"limit_pay,omitempty" json:"limit_pay,omitempty"`   // 上传此参数no_credit--可限制用户不能使用信用卡支付
	Receipt      string     `xml:"receipt,omitempty" json:"receipt,omitempty"`       // Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	SceneInfo    *SceneInfo `xml:"-" json:"-"`                                       // 该字段常用于线下活动时的场景信息上报，支持上报实际门店信息，商户也可以按需求自己上报相关信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开
	SceneInfoStr string     `xml:"scene_info,omitempty" json:"scene_info,omitempty"`
	GoodsTag     string     `xml:"goods_tag,omitempty" json:"goods_tag,omitempty"` // 订单优惠标记，使用代金券或立减优惠功能时需要的参数，说明详见代金券或立减优惠
	Version      string     `xml:"version,omitempty" json:"version,omitempty"`     // 新增字段，接口版本号，区分原接口，默认填写1.0。入参新增version后，则支付通知接口也将返回单品优惠信息字段promotion_detail，请确保支付通知的签名验证能通过。
}

func (this *PayOrderOptional) toUrlValues() url.Values {
	p := url.Values{}

	p.Set("limit_pay", this.LimitPay)
	p.Set("product_id", this.ProductId)
	p.Set("receipt", this.Receipt)

	if this.SceneInfo != nil {
		if len(this.SceneInfoStr) > 0 {
			p.Set("scene_info", this.SceneInfoStr)
		} else {
			var sceneInfoByte, err = json.Marshal(this.SceneInfo)
			if err == nil {
				this.SceneInfoStr = "{\"store_info\" :" + string(sceneInfoByte) + "}"
				p.Set("scene_info", this.SceneInfoStr)
			}
		}
	}
	p.Set("goods_tag", this.GoodsTag)

	return p
}

// https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_1
type PayUnifiedOrderReq struct {
	BaseReq
	AppId          string
	SubAppId       string
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no" validate:"required"`         // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。详见商户订单号
	TotalFee       int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`          // 订单总金额，单位为分，详见支付金额
	Body           string `xml:"body" json:"body" validate:"required"`                         // 商品简单描述，该字段请按照规范传递，具体请见参数规定
	TradeType      string `xml:"trade_type" json:"trade_type" validate:"required"`             // JSAPI，NATIVE，APP等，说明详见参数规定
	SpBillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip" validate:"required"` // 支持IPV4和IPV6两种格式的IP地址。用户的客户端IPP。
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`                 // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型

	OpenId     string `xml:"openid,omitempty" json:"openid,omitempty"`           // trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
	SubOpenId  string `xml:"sub_openid" json:"sub_openid"`                       //trade_type=JSAPI，此参数必传，用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid。下单前需要调用【网页授权获取用户信息】接口获取到用户的Openid。
	Attach     string `xml:"attach,omitempty" json:"attach,omitempty"`           // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"` // 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"

	TimeStart  string `xml:"time_start,omitempty" json:"time_start,omitempty"`   // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire string `xml:"time_expire,omitempty" json:"time_expire,omitempty"` // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则  注意：最短失效时间间隔必须大于5分钟

	GoodsDetail *GoodsDetail `xml:"-" json:"-"`
	DetailStr   string       `xml:"detail,omitempty" json:"detail,omitempty"` // 商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”

	PayOrderOptional
}

func (this *PayUnifiedOrderReq) ApiName() string {
	return APIPayUnifiedOrder
}

func (this *PayUnifiedOrderReq) toUrlValues() url.Values {
	p := this.PayOrderOptional.toUrlValues()
	if this.AppId != "" {
		p.Set("appid", this.AppId)
	}
	if this.SubAppId != "" {
		p.Set("appid", this.SubAppId)
	}
	if this.SubOpenId != "" {
		p.Set("sub_openid", this.SubOpenId)
	}
	if this.OpenId != "" {
		p.Set("openid", this.OpenId)
	}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("total_fee", fmt.Sprintf("%d", this.TotalFee))
	p.Set("body", this.Body)
	p.Set("trade_type", this.TradeType)
	p.Set("spbill_create_ip", this.SpBillCreateIp)
	p.Set("fee_type", this.FeeType)

	p.Set("time_start", this.TimeStart)
	p.Set("time_expire", this.TimeExpire)
	p.Set("device_info", this.DeviceInfo)
	p.Set("attach", this.Attach)

	if this.GoodsDetail != nil {
		if this.DetailStr != "" {
			var detailInfoByte, err = json.Marshal(this.GoodsDetail)
			if err == nil {
				this.DetailStr = string(detailInfoByte)
				p.Set("scene_info", this.DetailStr)
			}
		} else {
			p.Set("detail", this.DetailStr)
		}
	}
	return p
}

type PayUnifiedOrderResp struct {
	CodeMsg
	BaseResp
	DeviceInfo string `xml:"device_info" json:"device_info"` //设备号: 自定义参数，可以为请求支付的终端设备号等
	PrepayId   string `xml:"prepay_id" json:"prepay_id"`     //预支付交易会话标识: 微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
	TradeType  string `xml:"trade_type" json:"trade_type"`   //交易类型: JSAPI -JSAPI支付	NATIVE -Native支付	APP -APP支付	说明详见参数规定
	CodeURL    string `xml:"code_url" json:"code_url"`       //二维码链接: trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。	注意：code_url的值并非固定，使用时按照URL格式转成二维码即可
}

//https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_10&index=1
type PayMicroPayReq struct {
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no" validate:"required"`         // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。详见商户订单号
	TotalFee       int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`          // 订单总金额，单位为分，详见支付金额
	Body           string `xml:"body" json:"body" validate:"required"`                         // 商品简单描述，该字段请按照规范传递，具体请见参数规定
	SpBillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip" validate:"required"` // 支持IPV4和IPV6两种格式的IP地址。用户的客户端IPP。
	AuthCode       string `xml:"auth_code" json:"auth_code" validate:"required"`               //	扫码支付授权码，设备读取用户微信中的条码或者二维码信息 （注：用户付款码条形码规则：18位纯数字，以10、11、12、13、14、15开头）
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`                 // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型

	Attach     string `xml:"attach,omitempty" json:"attach,omitempty"`           // 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"` // 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"

	TimeStart  string `xml:"time_start,omitempty" json:"time_start,omitempty"`   // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire string `xml:"time_expire,omitempty" json:"time_expire,omitempty"` // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则  注意：最短失效时间间隔必须大于5分钟

	GoodsDetail *GoodsDetail `xml:"-" json:"-"`
	DetailStr   string       `xml:"detail,omitempty" json:"detail,omitempty"` // 商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”

	PayOrderOptional
}

func (this *PayMicroPayReq) NotifyUrl() string {
	return ""
}

func (this *PayMicroPayReq) ApiName() string {
	return APIPayMicroPay
}

func (this *PayMicroPayReq) toUrlValues() url.Values {
	p := this.PayOrderOptional.toUrlValues()
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("total_fee", fmt.Sprintf("%d", this.TotalFee))
	p.Set("body", this.Body)
	p.Set("spbill_create_ip", this.SpBillCreateIp)
	p.Set("auth_code", this.AuthCode)

	p.Set("time_start", this.TimeStart)
	p.Set("time_expire", this.TimeExpire)
	p.Set("device_info", this.DeviceInfo)
	p.Set("attach", this.Attach)

	if this.GoodsDetail != nil {
		if this.DetailStr != "" {
			var detailInfoByte, err = json.Marshal(this.GoodsDetail)
			if err == nil {
				this.DetailStr = string(detailInfoByte)
				p.Set("scene_info", this.DetailStr)
			}
		} else {
			p.Set("detail", this.DetailStr)
		}
	}
	return p
}

type PayMicroPayResp struct {
	CodeMsg
	BaseResp
	UserOpenIdRespInfo
	PayOrderInfo
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_2
type PayOrderQueryReq struct {
	BaseReq
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` //微信的订单号，建议优先使用
}

func (this *PayOrderQueryReq) ApiName() string {
	return APIPayOrderQuery
}

func (this *PayOrderQueryReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	return p
}

type PayOrderInfo struct {
	TransactionId string `xml:"transaction_id" json:"transaction_id"` //微信支付订单号: 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no"`     //商户订单号: 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	Attach        string `xml:"attach" json:"attach"`                 //附加数据: 附加数据，原样返回
	TimeEnd       string `xml:"time_end" json:"time_end"`             //支付完成时间: 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TradeType     string `xml:"trade_type" json:"trade_type"`         //交易类型: 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定

	DeviceInfo string `xml:"device_info" json:"device_info"` //设备号: 自定义参数，可以为请求支付的终端设备号等
	BankType   string `xml:"bank_type" json:"bank_type"`     //付款银行: 银行类型，采用字符串类型的银行标识
	TotalFee   int64  `xml:"total_fee" json:"total_fee"`     //订单总金额，单位为分
	FeeType    string `xml:"fee_type" json:"fee_type"`       //标价币种:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee    int64  `xml:"cash_fee" json:"cash_fee"`       //现金支付金额: 现金支付金额订单现金支付金额，详见支付金额

	CashFeeType        string `xml:"cash_fee_type" json:"cash_fee_type"`               //现金支付币种: 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	SettlementTotalFee int    `xml:"settlement_total_fee" json:"settlement_total_fee"` //应结订单金额: 订当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。

	CouponFee   int64 `xml:"coupon_fee" json:"coupon_fee"`     //代金券金额: “代金券”金额<=订单金额，订单金额-“代金券”金额=现金支付金额，详见支付金额
	CouponCount int   `xml:"coupon_count" json:"coupon_count"` //代金券使用数量: 代金券使用数量
}

type PayOrderQueryResp struct {
	CodeMsg
	BaseResp
	UserOpenIdRespInfo
	PayOrderInfo
	TradeState     string `xml:"trade_state" json:"trade_state"`           //交易状态: SUCCESS—支付成功	REFUND—转入退款	NOTPAY—未支付	CLOSED—已关闭 REVOKED—已撤销（付款码支付） USERPAYING--用户支付中（付款码支付） PAYERROR--支付失败(其他原因，如银行返回失败) 支付状态机请见下单API页面
	TradeStateDesc string `xml:"trade_state_desc" json:"trade_state_desc"` //交易状态描述: 对当前查询订单状态的描述和下一步操作的指引
}

//https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_11&index=3
type PayReverseReq struct {
	BaseReq
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` //微信的订单号，建议优先使用
}

func (this *PayReverseReq) ApiName() string {
	return APIPayReverse
}

func (this *PayReverseReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	return p
}

type PayReverseResp struct {
	CodeMsg
	BaseResp
	Recall string `xml:"recall" json:"recall"` //是否需要继续调用撤销，Y-需要，N-不需要
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_3
type PayCloseOrderReq struct {
	OutTradeNo string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"` //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
}

func (this *PayCloseOrderReq) NotifyUrl() string {
	return ""
}

func (this *PayCloseOrderReq) ApiName() string {
	return APIPayCloseOrder
}

func (this *PayCloseOrderReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	return p
}

type PayCloseOrderResp struct {
	CodeMsg
	BaseResp
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_4
type PayRefundReq struct {
	BaseReq
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`                       //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"`                   //微信的订单号，建议优先使用
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty" validate:"required"` //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	TotalFee      int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`                        //订单总金额，单位为分，只能为整数，详见支付金额
	RefundFee     int64  `xml:"refund_fee" json:"refund_fee" validate:"required,gt=0"`                      //退款总金额，单位为分，只能为整数，可部分退款。详见支付金额
	RefundFeeType string `xml:"refund_fee_type,omitempty" json:"refund_fee_type,omitempty"`                 //退款货币类型，需与支付一致，或者不填。符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	RefundDesc    string `xml:"refund_desc,omitempty" json:"refund_desc,omitempty"`                         //若商户传入，会在下发给用户的退款消息中体现退款原因 注意：若订单退款金额≤1元，且属于部分退款，则不会在退款消息中体现退款原因
	RefundAccount string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`                   //仅针对老资金流商户使用	REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款） REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款
}

func (this *PayRefundReq) ApiName() string {
	return APIPayRefund
}

func (this *PayRefundReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	p.Set("out_refund_no", this.OutRefundNo)
	p.Set("total_fee", fmt.Sprintf("%d", this.TotalFee))
	p.Set("refund_fee", fmt.Sprintf("%d", this.RefundFee))
	p.Set("refund_fee_type", this.RefundFeeType)
	p.Set("refund_desc", this.RefundDesc)
	p.Set("refund_account", this.RefundAccount)
	return p
}

type PayRefundResp struct {
	CodeMsg
	BaseResp
	OutRefundNo         string `xml:"out_refund_no" json:"out_refund_no"`                 //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId            string `xml:"refund_id" json:"refund_id"`                         //微信退款单号
	RefundFee           int64  `xml:"refund_fee" json:"refund_fee"`                       //退款总金额，单位为分，只能为整数，可部分退款。详见支付金额
	SettlementRefundFee int64  `xml:"settlement_refund_fee" json:"settlement_refund_fee"` //去掉非充值代金券退款金额后的退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	TotalFee            int64  `xml:"total_fee" json:"total_fee"`                         //订单总金额，单位为分，只能为整数，详见支付金额
	SettlementTotalFee  int64  `xml:"settlement_total_fee" json:"settlement_total_fee"`   //应结订单金额: 订当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType             string `xml:"fee_type" json:"fee_type"`                           //标价币种:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee             int64  `xml:"cash_fee" json:"cash_fee"`                           //现金支付金额: 现金支付金额订单现金支付金额，详见支付金额
	CashRefundFee       int64  `xml:"cash_refund_fee" json:"cash_refund_fee"`             //现金退款金额，单位为分，只能为整数，详见支付金额
	CouponRefundFee     int64  `xml:"coupon_refund_fee" json:"coupon_refund_fee"`         //代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠
	CouponRefundCount   int    `xml:"coupon_refund_count" json:"coupon_refund_count"`     //退款代金券使用数量
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_5
type PayRefundQueryReq struct {
	BaseReq
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` //微信的订单号，建议优先使用
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"`   //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId      string `xml:"refund_id" json:"refund_id"`                               //微信退款单号
	Offset        int    `xml:"offset" json:"offset"`                                     //偏移量，当部分退款次数超过10次时可使用，表示返回的查询结果从这个偏移量开始取记录
}

func (this *PayRefundQueryReq) ApiName() string {
	return APIPayRefundQuery
}

func (this *PayRefundQueryReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	p.Set("out_refund_no", this.OutRefundNo)
	p.Set("refund_id", this.RefundId)
	if this.Offset > 0 {
		p.Set("offset", fmt.Sprintf("%d", this.Offset))
	}
	return p
}

type PayRefundQueryResp struct {
	CodeMsg
	BaseResp
	TransactionId       string `xml:"transaction_id" json:"transaction_id"`               //微信的订单号
	OutRefundNo         string `xml:"out_refund_no" json:"out_refund_no"`                 //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	TotalFee            int64  `xml:"total_fee" json:"total_fee"`                         //订单总金额，单位为分，只能为整数，详见支付金额
	SettlementRefundFee int64  `xml:"settlement_refund_fee" json:"settlement_refund_fee"` //应结订单金额：当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType             string `xml:"fee_type" json:"fee_type"`                           //标价币种:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee             int64  `xml:"cash_fee" json:"cash_fee"`                           //现金支付金额: 现金支付金额订单现金支付金额，详见支付金额
	RefundCount         int    `xml:"refund_count" json:"refund_count"`                   //当前返回退款笔数
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_6
type PayDownloadBillReq struct {
	BaseReq
	BillDate string `xml:"bill_date" json:"bill_date" validate:"required"` // 下载对账单的日期，格式：20140603
	BillType string `xml:"bill_type,omitempty" json:"bill_type,omitempty"` // ALL，返回当日所有订单信息，默认值；SUCCESS，返回当日成功支付的订单；REFUND，返回当日退款订单；RECHARGE_REFUND，返回当日充值退款订单
	TarType  string `xml:"tar_type,omitempty" json:"tar_type,omitempty"`   // 非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
}

func (this *PayDownloadBillReq) ApiName() string {
	return APIPayDownloadBill
}

func (this *PayDownloadBillReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("bill_date", this.BillDate)
	p.Set("bill_type", this.BillType)
	p.Set("tar_type", this.TarType)
	return p
}

type PayDownloadBillResp struct {
	CodeMsg
	BaseResp
	Data []byte `xml:"-" json:"-"`
}

func (this *PayDownloadBillResp) IsSuccess() bool {
	return this.ReturnCode != CodeFail
}

type GetBrandWCPayRequestReq struct {
	BaseReq
	AppId    string `xml:"app_id" json:"app_id"`
	PrepayId string `xml:"prepay_id" json:"prepay_id" validate:"required"` //统一下单接口返回的prepay_id参数值，提交格式如：prepay_id=***
}

func (this *GetBrandWCPayRequestReq) ApiName() string {
	return ""
}

func (this *GetBrandWCPayRequestReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("prepay_id", this.PrepayId)
	return p
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
type GetBrandWCPayRequestResp struct {
	AppId     string `xml:"appId" json:"appId"`         //商户注册具有支付权限的公众号成功后即可获得
	Timestamp string `xml:"timeStamp" json:"timeStamp"` //当前的时间
	NonceStr  string `xml:"nonceStr" json:"nonceStr"`   //随机字符串，不长于32位。推荐随机数生成算法
	Package   string `xml:"package" json:"package"`     //统一下单接口返回的prepay_id参数值，提交格式如：prepay_id=***
	SignType  string `xml:"signType" json:"signType"`   //签名类型，默认为MD5，支持HMAC-SHA256和MD5。注意此处需与统一下单的签名类型一致
	PaySign   string `xml:"paySign" json:"paySign"`     //签名，详见签名生成算法
}

type PayitilReportReq struct {
	BaseReq
	DeviceInfo   string `xml:"device_info,omitempty" json:"device_info,omitempty"`     // 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	InterfaceUrl string `xml:"interface_url" json:"interface_url" validate:"required"` //接口URL 刷卡支付终端上报统一填：https://api.mch.weixin.qq.com/pay/batchreport/micropay/total
	UserIp       string `xml:"user_ip" json:"user_ip" validate:"required"`             //访问接口IP: 发起接口调用时的机器IP
	Trades       string `xml:"trades" json:"trades" validate:"required"`               //上报数据包: POS机采集的交易信息列表，使用JSON格式的数组，每条交易包含： 1. out_trade_no 商户订单号 2. begin_time 交易开始时间（扫码时间) 3. end_time 交易完成时间 4. state 交易结果 OK   -成功 FAIL -失败 CANCLE-取消 5. err_msg 自定义的错误描述信息 *注意，将JSON数组的文本串放到XML节点中时，需要使用!CDATA[]标签将JSON文本串保护起来
}

func (this *PayitilReportReq) ApiName() string {
	return APIPayitilReport
}

func (this *PayitilReportReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("device_info", this.DeviceInfo)
	p.Set("interface_url", this.InterfaceUrl)
	p.Set("user_ip", this.UserIp)
	p.Set("trades", this.Trades)
	return p
}

type PayitilReportResp struct {
	CodeMsg
}

//Deposit
//https: //pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_0&index=1
type DepositFacePayReq struct {
	BaseReq
	Deposit        string `xml:"deposit" json:"deposit" validate:"required"`                       //是否押金人脸支付，Y-是,N-普通人脸支付
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no" validate:"required"`             // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。详见商户订单号
	TotalFee       int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`              // 订单总金额，单位为分，详见支付金额
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty" validate:"required"` // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
	Body           string `xml:"body" json:"body" validate:"required"`                             // 商品简单描述，该字段请按照规范传递，具体请见参数规定
	SpBillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip" validate:"required"`     // 支持IPV4和IPV6两种格式的IP地址。用户的客户端IPP。
	FaceCode       string `xml:"face_code" json:"face_code" validate:"required"`                   //授权码: 		人脸凭证，用于人脸支付
	PayOrderOptional
}

func (this *DepositFacePayReq) ApiName() string {
	return APIDepositFacePay
}

func (this *DepositFacePayReq) toUrlValues() url.Values {
	p := this.PayOrderOptional.toUrlValues()
	p.Set("deposit", this.Deposit)
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("total_fee", fmt.Sprintf("%d", this.TotalFee))
	p.Set("body", this.Body)
	p.Set("spbill_create_ip", this.SpBillCreateIp)
	p.Set("face_code", this.FaceCode)
	p.Set("fee_type", this.FeeType)
	return p
}

type DepositFacePayResp struct {
	CodeMsg
	BaseResp
	UserOpenIdRespInfo
	PayOrderInfo
}

//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_1&index=2
type DepositMicroPayReq struct {
	BaseReq
	Deposit        string `xml:"deposit,omitempty" json:"deposit,omitempty"`                   //是否押金支付，Y-是,N-普通付款码支付
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no" validate:"required"`         // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|* 且在同一个商户号下唯一。详见商户订单号
	TotalFee       int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`          // 订单总金额，单位为分，详见支付金额
	Body           string `xml:"body" json:"body" validate:"required"`                         // 商品简单描述，该字段请按照规范传递，具体请见参数规定
	SpBillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip" validate:"required"` // 支持IPV4和IPV6两种格式的IP地址。用户的客户端IPP。
	AuthCode       string `xml:"auth_code" json:"auth_code" validate:"required"`
	FeeType        string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`       // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
	TimeStart      string `xml:"time_start,omitempty" json:"time_start,omitempty"`   // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire     string `xml:"time_expire,omitempty" json:"time_expire,omitempty"` // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。注意：最短失效时间间隔需大于1分钟
	Attach         string `xml:"attach,omitempty" json:"attach,omitempty"`           //附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	PayOrderOptional
}

func (this *DepositMicroPayReq) ApiName() string {
	return APIDepositMicroPay
}

func (this *DepositMicroPayReq) toUrlValues() url.Values {
	p := this.PayOrderOptional.toUrlValues()
	p.Set("deposit", this.Deposit)
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("total_fee", fmt.Sprintf("%d", this.TotalFee))
	p.Set("body", this.Body)
	p.Set("spbill_create_ip", this.SpBillCreateIp)
	p.Set("auth_code", this.AuthCode)
	p.Set("fee_type", this.FeeType)
	p.Set("time_start", this.TimeStart)
	p.Set("time_expire", this.TimeExpire)
	p.Set("attach", this.Attach)
	return p
}

type DepositMicroPayResp struct {
	CodeMsg
	BaseResp
	UserOpenIdRespInfo
	PayOrderInfo
}

//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_2&index=3
type DepositOrderQueryReq struct {
	BaseReq
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` //微信的订单号，建议优先使用
}

func (this *DepositOrderQueryReq) ApiName() string {
	return APIDepositOrderQuery
}

func (this *DepositOrderQueryReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	return p
}

type DepositOrderQueryResp struct {
	CodeMsg
	BaseResp
	UserOpenIdRespInfo
	PayOrderInfo
	TradeState      string `xml:"trade_state" json:"trade_state"`           //交易状态: 刷卡支付交易状态： SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭 REVOKED—已撤销(刷卡支付) USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败) 押金支付交易状态： NOTPAY—未支付 USERPAYING--用户支付中 PAYERROR--支付失败 SUCCESS?支付成功，资金冻结中 REVOKED—已撤销 SETTLING—押金消费已受理 SETTLEMENTFAIL ?押金解除冻结失败 CONSUMED—押金消费成功
	TradeStateDesc  string `xml:"trade_state_desc" json:"trade_state_desc"` //交易状态描述: 对当前查询订单状态的描述和下一步操作的指引
	PromotionDetail string `xml:"promotion_detail" json:"promotion_detail"` //营销详情: 营销详情列表，使返回值为Json格式
	ConsumeFee      string `xml:"consume_fee" json:"consume_fee"`           //押金消费金额，用户消费后结算给商户的金额
}

//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_3&index=4
type DepositReverseReq struct {
	BaseReq
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` //微信的订单号，建议优先使用
}

func (this *DepositReverseReq) ApiName() string {
	return APIDepositReverse
}

func (this *DepositReverseReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	return p
}

type DepositReverseResp struct {
	CodeMsg
	BaseResp
}

//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_4&index=5
type DepositConsumeReq struct {
	BaseReq
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 微信的订单号，建议优先使用
	OutTradeNo    string `xml:"out_trade_no,omitempty" json:"out_trade_no,omitempty"`     // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TotalFee      int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`      // 押金总金额: 订单总金额，单位为分，详见支付金额
	ConsumeFee    int64  `xml:"consume_fee" json:"consume_fee" validate:"required,gt=0"`  // 订单总金额，单位为分，只能为整数，详见支付金额
	FeeType       string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`             // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
}

func (this *DepositConsumeReq) ApiName() string {
	return APIDepositReverse
}

func (this *DepositConsumeReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_trade_no", this.OutTradeNo)
	p.Set("transaction_id", this.TransactionId)
	return p
}

type DepositConsumeResp struct {
	CodeMsg
	BaseResp
	TransactionId string `xml:"transaction_id" json:"transaction_id"` // 微信的订单号，建议优先使用
	OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no"`     // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	TotalFee      int64  `xml:"total_fee" json:"total_fee"`           // 押金总金额: 订单总金额，单位为分，详见支付金额
	ConsumeFee    int64  `xml:"consume_fee" json:"consume_fee" `      // 订单总金额，单位为分，只能为整数，详见支付金额
	FeeType       string `xml:"fee_type" json:"fee_type"`             // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
}

//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_5&index=6
type DepositRefundReq struct {
	BaseReq
	TransactionId string `xml:"transaction_id" json:"transaction_id" validate:"required"`                   //微信的订单号，建议优先使用
	OutRefundNo   string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty" validate:"required"` //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	TotalFee      int64  `xml:"total_fee" json:"total_fee" validate:"required,gt=0"`                        //订单总金额，单位为分，只能为整数，详见支付金额
	RefundFee     int64  `xml:"refund_fee" json:"refund_fee" validate:"required,gt=0"`                      //退款总金额，单位为分，只能为整数，可部分退款。详见支付金额
	RefundFeeType string `xml:"refund_fee_type,omitempty" json:"refund_fee_type,omitempty"`                 //退款货币类型，需与支付一致，或者不填。符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	RefundDesc    string `xml:"refund_desc,omitempty" json:"refund_desc,omitempty"`                         //若商户传入，会在下发给用户的退款消息中体现退款原因 注意：若订单退款金额≤1元，且属于部分退款，则不会在退款消息中体现退款原因
	RefundAccount string `xml:"refund_account,omitempty" json:"refund_account,omitempty"`                   //仅针对老资金流商户使用	REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款） REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款
}

func (this *DepositRefundReq) ApiName() string {
	return APIDepositRefund
}

func (this *DepositRefundReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("transaction_id", this.TransactionId)
	p.Set("out_refund_no", this.OutRefundNo)
	p.Set("total_fee", fmt.Sprintf("%d", this.TotalFee))
	p.Set("refund_fee", fmt.Sprintf("%d", this.RefundFee))
	p.Set("refund_fee_type", this.RefundFeeType)
	p.Set("refund_desc", this.RefundDesc)
	p.Set("refund_account", this.RefundAccount)
	return p
}

type DepositRefundResp struct {
	CodeMsg
	BaseResp
	TransactionId       string `xml:"transaction_id" json:"transaction_id"`               //微信的订单号，建议优先使用
	OutTradeNo          string `xml:"out_trade_no" json:"out_trade_no"`                   //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	OutRefundNo         string `xml:"out_refund_no" json:"out_refund_no"`                 //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId            string `xml:"refund_id" json:"refund_id"`                         //微信退款单号
	RefundFee           int64  `xml:"refund_fee" json:"refund_fee"`                       //退款总金额，单位为分，只能为整数，可部分退款。详见支付金额
	SettlementRefundFee int64  `xml:"settlement_refund_fee" json:"settlement_refund_fee"` //去掉非充值代金券退款金额后的退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	TotalFee            int64  `xml:"total_fee" json:"total_fee"`                         //订单总金额，单位为分，只能为整数，详见支付金额
	SettlementTotalFee  int64  `xml:"settlement_total_fee" json:"settlement_total_fee"`   //应结订单金额: 订当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
	FeeType             string `xml:"fee_type" json:"fee_type"`                           //标价币种:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee             int64  `xml:"cash_fee" json:"cash_fee"`                           //现金支付金额: 现金支付金额订单现金支付金额，详见支付金额
	CashRefundFee       int64  `xml:"cash_refund_fee" json:"cash_refund_fee"`             //现金退款金额，单位为分，只能为整数，详见支付金额
	CouponRefundFee     int64  `xml:"coupon_refund_fee" json:"coupon_refund_fee"`         //代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠
}

//https://pay.weixin.qq.com/wiki/doc/api/deposit_sl.php?chapter=27_6&index=7
type DepositRefundQueryReq struct {
	BaseReq
	OutRefundNo string `xml:"out_refund_no,omitempty" json:"out_refund_no,omitempty"` //商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	RefundId    string `xml:"refund_id" json:"refund_id"`                             //微信退款单号
}

func (this *DepositRefundQueryReq) ApiName() string {
	return APIDepositRefundQuery
}

func (this *DepositRefundQueryReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("out_refund_no", this.OutRefundNo)
	p.Set("refund_id", this.RefundId)
	return p
}

type DepositRefundQueryResp struct {
	CodeMsg
	BaseResp
	TransactionId string `xml:"transaction_id" json:"transaction_id"` //微信的订单号
	OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no"`     //商户订单号 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	TotalFee      int64  `xml:"total_fee" json:"total_fee"`           //订单总金额，单位为分，只能为整数，详见支付金额

	SettlementTotalFee int `xml:"settlement_total_fee" json:"settlement_total_fee"` //应结订单金额：当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。

	FeeType       string `xml:"fee_type" json:"fee_type"`             //标价币种:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFee       int64  `xml:"cash_fee" json:"cash_fee"`             //现金支付金额: 现金支付金额订单现金支付金额，详见支付金额
	OutRefundNo   string `xml:"out_refund_no" json:"out_refund_no"`   //商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	RefundId      string `xml:"refund_id" json:"refund_id"`           //微信退款单号
	RefundChannel string `xml:"refund_channel" json:"refund_channel"` //退款渠道: ORIGINAL—原路退款 BALANCE—退回到余额 OTHER_BALANCE—原账户异常退到其他余额账户 OTHER_BANKCARD—原银行卡异常退到其他银行卡
	RefundFee     int64  `xml:"refund_fee" json:"refund_fee"`         //申请退款金额:退款总金额,单位为分,可以做部分退款

	SettlementRefundFee string `xml:"settlement_refund_fee" json:"settlement_refund_fee"` //退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额

	RefundStatus      string `xml:"refund_status" json:"refund_status"`             //退款状态： SUCCESS—退款成功 REFUNDCLOSE—退款关闭。 PROCESSING—退款处理中 CHANGE—退款异常，退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。$n为下标，从0开始编号。
	RefundAccount     string `xml:"refund_account" json:"refund_account"`           //退款资金来源: REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款/基本账户 REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款
	RefundRecvAccout  string `xml:"refund_recv_accout" json:"refund_recv_accout"`   //退款入账账户: 取当前退款单的退款入账方 1）退回银行卡： {银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱: 支付用户零钱 3）退还商户: 商户基本账户 商户结算银行账户 4）退回支付用户零钱通: 支付用户零钱通
	RefundSuccessTime string `xml:"refund_success_time" json:"refund_success_time"` //退款成功时间 //退款成功时间，当退款状态为退款成功时有返回。
}

type UserAccessTokenReq struct {
	BaseReq
	AppId     string `json:"appid" validate:"required"`      //公众号的唯一标识
	AppSecret string `json:"secret" validate:"required"`     //公众号的app secret
	Code      string `json:"code" validate:"required"`       //填写第一步获取的code参数
	GrantType string `json:"grant_type" validate:"required"` //填写为authorization_code
}

func (this *UserAccessTokenReq) ApiName() string {
	return APIAccessToken
}

func (this *UserAccessTokenReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("appid", this.AppId)
	p.Set("secret", this.AppSecret)
	p.Set("code", this.Code)
	p.Set("grant_type", "authorization_code")
	return p
}

type UserRefreshTokenReq struct {
	BaseReq
	AppId        string `json:"appid" validate:"required"`          //公众号的唯一标识
	GrantType    string `json:"grant_type" validate:"required"`     //为refresh_token
	RefreshToken string `json:"refresh_token"  validate:"required"` //用户刷新access_token
}

func (this *UserRefreshTokenReq) ApiName() string {
	return APIRefreshToken
}

func (this *UserRefreshTokenReq) toUrlValues() url.Values {
	p := url.Values{}
	p.Set("appid", this.AppId)
	p.Set("grant_type", this.GrantType)
	p.Set("refresh_token", "refresh_token")
	return p
}

type UserAccessTokenResp struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	OpenId       string `json:"openid"`        //用户唯一标识，请注意，在未关注公众号时，用户访问公众号的网页，也会产生一个用户和公众号唯一的OpenID
	AccessToken  string `json:"access_token"`  //网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int64  `json:"expires_in"`    //access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` //用户刷新access_token
	Scope        string `json:"scope"`         //用户授权的作用域，使用逗号（,）分隔
}

func (this *UserAccessTokenResp) Error() string {
	return fmt.Sprintf("%d:%s", this.ErrCode, this.ErrMsg)
}

//https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_7
type TradeNotification struct {
	CodeMsg
	BaseResp
	UserOpenIdRespInfo
	PayOrderInfo
}
