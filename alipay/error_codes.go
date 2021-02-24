package alipay

//goland:noinspection ALL
const (
	// https://docs.open.alipay.com/common/105806
	CodeSuccess            = "10000" //接口调用成功
	CodeServiceUnAvailable = "20000" //服务不可用
	CodeUnAuthorized       = "20001" //授权权限不足
	CodeMissArgument       = "40001" //缺少必选参数
	CodeInvalidArgument    = "40002" //非法的参数
	CodeBusinessFailed     = "40004" //业务处理失败
	CodeMissPermission     = "40006" //权限不足

	SubCodeIspUnknowError = "isp.unknow-error" //服务暂不可用（业务系统不可用）
	SubCodeAopUnknowError = "aop.unknow-error" //服务暂不可用（网关自身的未知错误）

	SubCodeAopInvalidAuthToken = "aop.invalid-auth-token" //无效的访问令牌
)

//noinspection ALL
const (
	ACQ_SYSTEM_ERROR                    = "ACQ.SYSTEM_ERROR"                    //接口返回错误  请立即调用查询订单API，查询当前订单的状态，并根据订单状态决定下一步的操作
	ACQ_INVALID_PARAMETER               = "ACQ.INVALID_PARAMETER"               //参数无效  检查请求参数，修改后重新发起请求
	ACQ_ACCESS_FORBIDDEN                = "ACQ.ACCESS_FORBIDDEN"                //无权限使用接口  联系支付宝小二进行签约
	ACQ_EXIST_FORBIDDEN_WORD            = "ACQ.EXIST_FORBIDDEN_WORD"            //订单信息中包含违禁词  修改订单信息后，重新发起请求
	ACQ_PARTNER_ERROR                   = "ACQ.PARTNER_ERROR"                   //应用APP_ID填写错误  联系支付宝小二，确认APP_ID的状态
	ACQ_TOTAL_FEE_EXCEED                = "ACQ.TOTAL_FEE_EXCEED"                //订单总金额超过限额  修改订单金额再发起请求
	ACQ_CONTEXT_INCONSISTENT            = "ACQ.CONTEXT_INCONSISTENT"            //交易信息被篡改  更换商家订单号后，重新发起请求
	ACQ_TRADE_HAS_SUCCESS               = "ACQ.TRADE_HAS_SUCCESS"               //交易已被支付  确认该笔交易信息是否为当前买家的，如果是则认为交易付款成功，如果不是则更换商家订单号后，重新发起请求
	ACQ_TRADE_HAS_CLOSE                 = "ACQ.TRADE_HAS_CLOSE"                 //交易已经关闭  更换商家订单号后，重新发起请求
	ACQ_BUYER_SELLER_EQUAL              = "ACQ.BUYER_SELLER_EQUAL"              //买卖家不能相同  更换买家重新付款
	ACQ_TRADE_BUYER_NOT_MATCH           = "ACQ.TRADE_BUYER_NOT_MATCH"           //交易买家不匹配  更换商家订单号后，重新发起请求
	ACQ_BUYER_ENABLE_STATUS_FORBID      = "ACQ.BUYER_ENABLE_STATUS_FORBID"      //买家状态非法  用户联系支付宝小二，确认买家状态为什么非法
	ACQ_SELLER_BEEN_BLOCKED             = "ACQ.SELLER_BEEN_BLOCKED"             //商家账号被冻结  联系支付宝小二，解冻账号
	ACQ_ERROR_BUYER_CERTIFY_LEVEL_LIMIT = "ACQ.ERROR_BUYER_CERTIFY_LEVEL_LIMIT" //买家未通过人行认证  让用户联系支付宝小二并更换其它付款方式
	ACQ_SUB_MERCHANT_CREATE_FAIL        = "ACQ.SUB_MERCHANT_CREATE_FAIL"        //二级商户创建失败  检查上送的二级商户信息是否有效
	ACQ_SUB_MERCHANT_TYPE_INVALID       = "ACQ.SUB_MERCHANT_TYPE_INVALID"       //二级商户类型非法  检查上传的二级商户类型是否有效

	ACQ_PAYMENT_AUTH_CODE_INVALID         = "ACQ.PAYMENT_AUTH_CODE_INVALID"         //支付授权码无效  用户刷新条码后，重新扫码发起请求
	ACQ_BUYER_BALANCE_NOT_ENOUGH          = "ACQ.BUYER_BALANCE_NOT_ENOUGH"          //买家余额不足  买家绑定新的银行卡或者支付宝余额有钱后再发起支付
	ACQ_BUYER_BANKCARD_BALANCE_NOT_ENOUGH = "ACQ.BUYER_BANKCARD_BALANCE_NOT_ENOUGH" //用户银行卡余额不足  建议买家更换支付宝进行支付或者更换其它付款方式
	ACQ_ERROR_BALANCE_PAYMENT_DISABLE     = "ACQ.ERROR_BALANCE_PAYMENT_DISABLE"     //余额支付功能关闭  用户打开余额支付开关后，再重新进行支付
	ACQ_PULL_MOBILE_CASHIER_FAIL          = "ACQ.PULL_MOBILE_CASHIER_FAIL"          //唤起移动收银台失败  用户刷新条码后，重新扫码发起请求
	ACQ_MOBILE_PAYMENT_SWITCH_OFF         = "ACQ.MOBILE_PAYMENT_SWITCH_OFF"         //用户的无线支付开关关闭  用户在PC上打开无线支付开关后，再重新发起支付
	ACQ_PAYMENT_FAIL                      = "ACQ.PAYMENT_FAIL"                      //支付失败  用户刷新条码后，重新发起请求，如果重试一次后仍未成功，更换其它方式付款
	ACQ_PAYMENT_REQUEST_HAS_RISK          = "ACQ.PAYMENT_REQUEST_HAS_RISK"          //支付有风险  更换其它付款方式
	ACQ_NO_PAYMENT_INSTRUMENTS_AVAILABLE  = "ACQ.NO_PAYMENT_INSTRUMENTS_AVAILABLE"  //没用可用的支付工具  更换其它付款方式
	ACQ_USER_FACE_PAYMENT_SWITCH_OFF      = "ACQ.USER_FACE_PAYMENT_SWITCH_OFF"      //用户当面付付款开关关闭  让用户在手机上打开当面付付款开关
	ACQ_AGREEMENT_NOT_EXIST               = "ACQ.AGREEMENT_NOT_EXIST"               //用户协议不存在  确认代扣业务传入的协议号对应的协议是否已解约
	ACQ_AGREEMENT_INVALID                 = "ACQ.AGREEMENT_INVALID"                 //用户协议失效  代扣业务传入的协议号对应的用户协议已经失效，需要用户重新签约
	ACQ_AGREEMENT_STATUS_NOT_NORMAL       = "ACQ.AGREEMENT_STATUS_NOT_NORMAL"       //用户协议状态非NORMAL  代扣业务用户协议状态非正常状态，需要用户解约后重新签约
	ACQ_MERCHANT_AGREEMENT_NOT_EXIST      = "ACQ.MERCHANT_AGREEMENT_NOT_EXIST"      //商户协议不存在  确认商户与支付宝是否已签约
	ACQ_MERCHANT_AGREEMENT_INVALID        = "ACQ.MERCHANT_AGREEMENT_INVALID"        //商户协议已失效  商户与支付宝合同已失效，需要重新签约
	ACQ_MERCHANT_STATUS_NOT_NORMAL        = "ACQ.MERCHANT_STATUS_NOT_NORMAL"        //商户协议状态非正常状态  商户与支付宝的合同非正常状态，需要重新签商户合同
	ACQ_CARD_USER_NOT_MATCH               = "ACQ.CARD_USER_NOT_MATCH"               //脱机记录用户信息不匹配  请检查传入的进展出站记录是否正确
	ACQ_CARD_TYPE_ERROR                   = "ACQ.CARD_TYPE_ERROR"                   //卡类型错误  检查传入的卡类型
	ACQ_CERT_EXPIRED                      = "ACQ.CERT_EXPIRED"                      //凭证过期  凭证已经过期
	ACQ_AMOUNT_OR_CURRENCY_ERROR          = "ACQ.AMOUNT_OR_CURRENCY_ERROR"          //订单金额或币种信息错误  检查订单传入的金额信息是否有误，或者是不是当前币种未签约
	ACQ_CURRENCY_NOT_SUPPORT              = "ACQ.CURRENCY_NOT_SUPPORT"              //订单币种不支持  请检查是否签约对应的币种
	ACQ_MERCHANT_UNSUPPORT_ADVANCE        = "ACQ.MERCHANT_UNSUPPORT_ADVANCE"        //先享后付2.0准入失败,商户不支持垫资支付产品  先享后付2.0准入失败,商户不支持垫资支付产品
	ACQ_BUYER_UNSUPPORT_ADVANCE           = "ACQ.BUYER_UNSUPPORT_ADVANCE"           //先享后付2.0准入失败,买家不满足垫资条件  先享后付2.0准入失败,买家不满足垫资条件
	ACQ_ORDER_UNSUPPORT_ADVANCE           = "ACQ.ORDER_UNSUPPORT_ADVANCE"           //订单不支持先享后付垫资  订单不支持先享后付垫资
	ACQ_CYCLE_PAY_DATE_NOT_MATCH          = "ACQ.CYCLE_PAY_DATE_NOT_MATCH"          //扣款日期不在签约时的允许范围之内  对于周期扣款产品，签约时会约定扣款的周期。如果发起扣款的日期不符合约定的周期，则不允许扣款。请重新检查扣款日期，在符合约定的日期发起扣款。
	ACQ_CYCLE_PAY_SINGLE_FEE_EXCEED       = "ACQ.CYCLE_PAY_SINGLE_FEE_EXCEED"       //周期扣款的单笔金额超过签约时限制  对于周期扣款产品，签约时会约定单笔扣款的最大金额。如果发起扣款的金额大于约定上限，则不允许扣款。请在允许的金额范围内扣款。
	ACQ_CYCLE_PAY_TOTAL_FEE_EXCEED        = "ACQ.CYCLE_PAY_TOTAL_FEE_EXCEED"        //周期扣款的累计金额超过签约时限制  对于周期扣款产品，签约时可以约定多次扣款的累计金额限制。如果发起扣款的累计金额大于约定上限，则不允许扣款。请在允许的金额范围内扣款。
	ACQ_CYCLE_PAY_TOTAL_TIMES_EXCEED      = "ACQ.CYCLE_PAY_TOTAL_TIMES_EXCEED"      //周期扣款的总次数超过签约时限制  对于周期扣款产品，签约时可以约定多次扣款的总次数限制。如果发起扣款的总次数大于约定上限，则不允许扣款。请在允许的次数范围内扣款
	ACQ_SECONDARY_MERCHANT_STATUS_ERROR   = "ACQ.SECONDARY_MERCHANT_STATUS_ERROR"   //商户状态异常  请联系对应的服务商咨询

	ACQ_BUYER_PAYMENT_AMOUNT_DAY_LIMIT_ERROR   = "ACQ.BUYER_PAYMENT_AMOUNT_DAY_LIMIT_ERROR"   //买家付款日限额超限  更换买家进行支付
	ACQ_BEYOND_PAY_RESTRICTION                 = "ACQ.BEYOND_PAY_RESTRICTION"                 //商户收款额度超限  联系支付宝小二提高限额
	ACQ_BEYOND_PER_RECEIPT_RESTRICTION         = "ACQ.BEYOND_PER_RECEIPT_RESTRICTION"         //商户收款金额超过月限额  联系支付宝小二提高限额
	ACQ_BUYER_PAYMENT_AMOUNT_MONTH_LIMIT_ERROR = "ACQ.BUYER_PAYMENT_AMOUNT_MONTH_LIMIT_ERROR" //买家付款月额度超限  让买家更换账号后，重新付款或者更换其它付款方式
	ACQ_INVALID_STORE_ID                       = "ACQ.INVALID_STORE_ID"                       //商户门店编号无效  检查传入的门店编号是否符合规则

	AQC_SYSTEM_ERROR = "AQC.SYSTEM_ERROR" //系统错误  请使用相同的参数再次调用

	ACQ_SELLER_BALANCE_NOT_ENOUGH   = "ACQ.SELLER_BALANCE_NOT_ENOUGH"   //卖家余额不足  商户支付宝账户充值后重新发起退款即可
	ACQ_REFUND_AMT_NOT_EQUAL_TOTAL  = "ACQ.REFUND_AMT_NOT_EQUAL_TOTAL"  //退款金额超限  检查退款金额是否正确，重新修改请求后，重新发起退款
	ACQ_REASON_TRADE_BEEN_FREEZEN   = "ACQ.REASON_TRADE_BEEN_FREEZEN"   //请求退款的交易被冻结  联系支付宝小二，确认该笔交易的具体情况
	ACQ_TRADE_NOT_EXIST             = "ACQ.TRADE_NOT_EXIST"             //交易不存在  检查请求中的交易号和商户订单号是否正确，确认后重新发起
	ACQ_TRADE_HAS_FINISHED          = "ACQ.TRADE_HAS_FINISHED"          //交易已完结  该交易已完结，不允许进行退款，确认请求的退款的交易信息是否正确
	ACQ_TRADE_STATUS_ERROR          = "ACQ.TRADE_STATUS_ERROR"          //交易状态非法  查询交易，确认交易是否已经付款
	ACQ_DISCORDANT_REPEAT_REQUEST   = "ACQ.DISCORDANT_REPEAT_REQUEST"   //不一致的请求  检查该退款号是否已退过款或更换退款号重新发起请求
	ACQ_REASON_TRADE_REFUND_FEE_ERR = "ACQ.REASON_TRADE_REFUND_FEE_ERR" //退款金额无效  检查退款请求的金额是否正确
	ACQ_TRADE_NOT_ALLOW_REFUND      = "ACQ.TRADE_NOT_ALLOW_REFUND"      //当前交易不允许退款  检查当前交易的状态是否为交易成功状态以及签约的退款属性是否允许退款，确认后，重新发起请求
	ACQ_REFUND_FEE_ERROR            = "ACQ.REFUND_FEE_ERROR"            //交易退款金额有误  请检查传入的退款金额是否正确

	ILLEGAL_ARGUMENT                 = "ILLEGAL_ARGUMENT"                 //授权失败，预授权冻结参数异常或参数缺失，请顾客刷新付款码后重新收款  检查请求参数，修改后重新发起请求
	EXIST_FORBIDDEN_WORD             = "EXIST_FORBIDDEN_WORD"             //授权失败，订单信息中包含违禁词  修改订单信息后，重新发起请求
	ACCESS_FORBIDDEN                 = "ACCESS_FORBIDDEN"                 //授权失败，本商户没有权限使用该产品，建议顾客使用其他方式付款  未签约条码支付或者合同已到期
	UNIQUE_VIOLATION                 = "UNIQUE_VIOLATION"                 //授权失败，商户订单号重复，请收银员取消本笔订单并重新授权  更换商户的授权资金订单号后，重新发起请求
	PAYER_USER_STATUS_LIMIT          = "PAYER_USER_STATUS_LIMIT"          //授权失败，顾客账户暂时无法支付，建议顾客使用其他方式付款  买家支付宝账户受限，请登录支付宝认证升级，详情咨询95188
	PAYER_NOT_EXIST                  = "PAYER_NOT_EXIST"                  //授权失败，获取顾客账户信息失败，请顾客刷新付款码后重新收款  用户刷新条码后，重新扫码发起请求
	PAYMENT_AUTH_CODE_INVALID        = "PAYMENT_AUTH_CODE_INVALID"        //授权失败，获取顾客账户信息失败，请顾客刷新付款码后重新收款，如再次授权失败  用户刷新条码后，重新扫码发起请求
	MONEY_NOT_ENOUGH                 = "MONEY_NOT_ENOUGH"                 //授权失败，顾客余额不足，建议顾客充值完成后再进行付款  买家绑定新的银行卡或者支付宝余额有钱后再发起支付
	ORDER_ALREADY_CLOSED             = "ORDER_ALREADY_CLOSED"             //授权失败，本笔授权订单已关闭  更换商户授权资金订单号后，重新发起请求
	FREEZE_ALREADY_SUCCESS           = "FREEZE_ALREADY_SUCCESS"           //授权失败，授权订单已经冻结成功，请勿重复授权  确认该笔预授权信息是否为当前付款方的，如果是则认为授权成功，如果不是则更换商家授权资金订单号后，重新发起请求
	ERROR_BALANCE_PAYMENT_DISABLE    = "ERROR_BALANCE_PAYMENT_DISABLE"    //授权失败，顾客余额支付功能开关关闭，请用户打开余额支付功能开关  用户打开余额支付开关后，再重新进行支付
	PULL_MOBILE_CASHIER_FAIL         = "PULL_MOBILE_CASHIER_FAIL"         //授权失败，顾客手机唤起收银台失败，请顾客检查手机网络，刷新付款码后重新预授权，并让顾客在付款码页面等待确认  用户检查手机网络，刷新条码后，重新扫码发起请求
	USER_FACE_PAYMENT_SWITCH_OFF     = "USER_FACE_PAYMENT_SWITCH_OFF"     //授权失败，顾客当面付付款开关关闭，请用户在手机上打开当面付付款开关  让用户在手机上打开当面付付款开关
	SYSTEM_ERROR                     = "SYSTEM_ERROR"                     //系统错误  请立即调用查询订单API，查询当前订单的状态，并根据订单状态决定下一步的操作
	ORDER_ALREADY_FINISH             = "ORDER_ALREADY_FINISH"             //授权失败，本笔授权订单已经完结，无法再进行资金操作  更换商家授权资金订单号后，重新发起请求
	PAYEE_NOT_EXIST                  = "PAYEE_NOT_EXIST"                  //授权失败，收款方账号不存在  确认该收款方账号是注册过的支付宝账号
	PAYEE_USER_STATUS_LIMIT          = "PAYEE_USER_STATUS_LIMIT"          //授权失败，收款方账号异常  卖家支付宝账户受限，请登录支付宝认证升级，详情咨询95188
	PAYER_PAYEE_EQUAL                = "PAYER_PAYEE_EQUAL"                //授权失败，收付款方信息不能相同  请商家基于业务诉求更换付款方或收款方信息
	NO_PAYMENT_INSTRUMENTS_AVAILABLE = "NO_PAYMENT_INSTRUMENTS_AVAILABLE" //授权失败，用户没用可用的支付工具  请用户更换其它付款方式
	CLIENT_VERSION_NOT_MATCH         = "CLIENT_VERSION_NOT_MATCH"         //授权失败，顾客手机支付宝客户端版本过低，请更新到最新版本  请用户更新到最新版本的手机支付宝客户端

	CURRENCY_VERIFICATION_FAIL       = "CURRENCY_VERIFICATION_FAIL"       //币种校验失败  确认标价币种、结算币种正确后重新发起请求
	SECONDARY_MERCHANT_STATUS_ERROR  = "SECONDARY_MERCHANT_STATUS_ERROR"  //商户状态异常  请商户与支付宝客服联系确认
	USER_IDENTITY_INFO_VALIDATE_FAIL = "USER_IDENTITY_INFO_VALIDATE_FAIL" //用户实名信息校验失败  用户实名信息校验失败，请确认当前用户与支付宝用户是否匹配。
	USER_ACCOUNT_VALIDATE_FAIL       = "USER_ACCOUNT_VALIDATE_FAIL"       //用户账号校验失败  用户账号信息校验失败，请确认当前用户与支付宝用户是否匹配。
	SUB_MERCHANT_ORG_ID_ERROR        = "SUB_MERCHANT_ORG_ID_ERROR"        //间联模式下，传入的二级商户的机构id为空或错误  间联模式下，传入正确的机构id
	SUB_MERCHANT_LEVEL_ERROR         = "SUB_MERCHANT_LEVEL_ERROR"         //间联商户等级校验错误  间联商户等级校验错误，请提高间联商户的等级
	SUB_MERCHANT_NO_PERMISSION       = "SUB_MERCHANT_NO_PERMISSION"       //间联商户无权使用该产品  商户暂时不符合产品的使用场景
	MERCHANT_STATUS_ERROR            = "MERCHANT_STATUS_ERROR"            //商户状态错误  商户状态异常，请联系支付宝核实

	REQUEST_AMOUNT_EXCEED = "REQUEST_AMOUNT_EXCEED" //请求解冻金额超限  更改解冻金额，重新发起请求
	ILLEGAL_STATUS        = "ILLEGAL_STATUS"        //订单状态非法  查询该笔授权操作信息，确认用户资金授权冻结成功
	BIZ_ERROR             = "BIZ_ERROR"             //业务异常，  商户自行确认该笔预授权订单是否被用于其他业务，或者联系支付宝客服

	AUTH_ORDER_NOT_EXIST     = "AUTH_ORDER_NOT_EXIST"     //支付宝资金授权订单不存在  检查传入参数中的支付宝资金授权订单号或商户授权订单号，修改后重新发起请求
	AUTH_OPERATION_NOT_EXIST = "AUTH_OPERATION_NOT_EXIST" //支付宝资金操作流水不存在  检查传入参数中的支付宝的授权资金操作流水号或商户的授权资金操作流水号，修改后重新发起请求
)
