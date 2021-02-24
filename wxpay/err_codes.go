package wxpay

//goland:noinspection ALL
const (
	SYSTEMERROR           = "SYSTEMERROR"           //接口返回错误  支付结果未知  系统超时请立即调用被扫订单结果查询API，查询当前订单状态，并根据订单的状态决定下一步的操作。
	PARAM_ERROR           = "PARAM_ERROR"           //参数错误  支付确认失败  请求参数未按指引进行填写请根据接口返回的详细信息检查您的程序
	ORDERPAID             = "ORDERPAID"             //订单已支付  支付确认失败  订单号重复请确认该订单号是否重复支付，如果是新单，请使用新订单号提交
	NOAUTH                = "NOAUTH"                //商户无权限  支付确认失败  商户没有开通被扫支付权限请开通商户号权限。请联系产品或商务申请
	AUTHCODEEXPIRE        = "AUTHCODEEXPIRE"        //二维码已过期，请用户在微信上刷新后再试  支付确认失败  用户的条码已经过期请收银员提示用户，请用户在微信上刷新条码，然后请收银员重新扫码。 直接将错误展示给收银员
	NOTENOUGH             = "NOTENOUGH"             //余额不足  支付确认失败  用户的零钱余额不足请收银员提示用户更换当前支付的卡，然后请收银员重新扫码。建议：商户系统返回给收银台的提示为“用户余额不足.提示用户换卡支付”
	NOTSUPORTCARD         = "NOTSUPORTCARD"         //不支持卡类型  支付确认失败  用户使用卡种不支持当前支付形式请用户重新选择卡种 建议：商户系统返回给收银台的提示为“该卡不支持当前支付，提示用户换卡支付或绑新卡支付”
	ORDERCLOSED           = "ORDERCLOSED"           //订单已关闭  支付确认失败  该订单已关商户订单号异常，请重新下单支付
	ORDERREVERSED         = "ORDERREVERSED"         //订单已撤销  支付确认失败  当前订单已经被撤销当前订单状态为“订单已撤销”，请提示用户重新支付
	BANKERROR             = "BANKERROR"             //银行系统异常  支付结果未知  银行端超时请立即调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作。
	USERPAYING            = "USERPAYING"            //用户支付中，需要输入密码  支付结果未知  该笔交易因为业务规则要求，需要用户输入支付密码。等待5秒，然后调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作。
	AUTH_CODE_ERROR       = "AUTH_CODE_ERROR"       //授权码参数错误  支付确认失败  请求参数未按指引进行填写每个二维码仅限使用一次，请刷新再试
	AUTH_CODE_INVALID     = "AUTH_CODE_INVALID"     //授权码检验错误  支付确认失败  收银员扫描的不是微信支付的条码请扫描微信支付被扫条码/二维码
	XML_FORMAT_ERROR      = "XML_FORMAT_ERROR"      //XML格式错误  支付确认失败  XML格式错误请检查XML参数格式是否正确
	REQUIRE_POST_METHOD   = "REQUIRE_POST_METHOD"   //请使用post方法  支付确认失败  未使用post传递参数请检查请求参数是否通过post方法提交
	SIGNERROR             = "SIGNERROR"             //签名错误  支付确认失败  参数签名结果不正确请检查签名参数和方法是否都符合签名算法要求
	LACK_PARAMS           = "LACK_PARAMS"           //缺少参数  支付确认失败  缺少必要的请求参数请检查参数是否齐全
	NOT_UTF8              = "NOT_UTF8"              //编码格式错误  支付确认失败  未使用指定编码格式请使用UTF-8编码格式
	BUYER_MISMATCH        = "BUYER_MISMATCH"        //支付帐号错误  支付确认失败  暂不支持同一笔订单更换支付方请确认支付方是否相同
	APPID_NOT_EXIST       = "APPID_NOT_EXIST"       //APPID不存在  支付确认失败  参数中缺少APPID请检查APPID是否正确
	MCHID_NOT_EXIST       = "MCHID_NOT_EXIST"       //MCHID不存在  支付确认失败  参数中缺少MCHID请检查MCHID是否正确
	OUT_TRADE_NO_USED     = "OUT_TRADE_NO_USED"     //商户订单号重复  支付确认失败  同一笔交易不能多次提交请核实商户订单号是否重复提交
	APPID_MCHID_NOT_MATCH = "APPID_MCHID_NOT_MATCH" //appid和mch_id不匹配  支付确认失败  appid和mch_id不匹配请确认appid和mch_id是否匹配
	INVALID_REQUEST       = "INVALID_REQUEST"       //无效请求  支付确认失败  商户系统异常导致，商户权限异常、重复请求支付、证书错误、频率限制等请确认商户系统是否正常，是否具有相应支付权限，确认证书是否正确，控制频率
	TRADE_ERROR           = "TRADE_ERROR"           //交易错误  支付确认失败  业务错误导致交易失败、用户账号异常、风控、规则限制等

	POST_DATA_EMPTY = "POST_DATA_EMPTY" //post数据为空    post数据不能为空请检查post数据是否为空

	ORDERNOTEXIST = "ORDERNOTEXIST" //此交易订单号不存在    查询系统中不存在此交易订单号该API只能查提交支付交易返回成功的订单，请商户检查需要查询的订单号是否正确

	BIZERR_NEED_RETRY     = "BIZERR_NEED_RETRY"     //退款业务流程错误，需要商户触发重试来解决    并发情况下，业务被拒绝，商户重试即可解决请不要更换商户退款单号，请使用相同参数再次调用API。
	TRADE_OVERDUE         = "TRADE_OVERDUE"         //订单已经超过退款期限    订单已经超过可退款的最大期限(支付后一年内可退款)请选择其他方式自行退款
	ERROR                 = "ERROR"                 //业务错误    申请退款业务发生错误该错误都会返回具体的错误原因，请根据实际返回做相应处理。
	USER_ACCOUNT_ABNORMAL = "USER_ACCOUNT_ABNORMAL" //退款请求失败    用户帐号注销此状态代表退款申请失败，商户可自行处理退款。
	INVALID_REQ_TOO_MUCH  = "INVALID_REQ_TOO_MUCH"  //无效请求过多    连续错误请求数过多被系统短暂屏蔽请检查业务是否正常，确认业务正常后请在1分钟后再来重试
	INVALID_TRANSACTIONID = "INVALID_TRANSACTIONID" //无效transaction_id    请求参数未按指引进行填写请求参数错误，检查原交易号是否存在或发起支付交易接口返回失败
	FREQUENCY_LIMITED     = "FREQUENCY_LIMITED"     //频率限制    2个月之前的订单申请退款有频率限制该笔退款未受理，请降低频率后重试

	REFUNDNOTEXIST = "REFUNDNOTEXIST" //退款订单查询失败    订单号错误或订单状态不正确请检查订单号是否有误以及订单状态是否正确，如：未支付、已支付未退款

)
