package wxpay

import (
	"fmt"
	"github.com/guestin/go_pay_cli/internal"
	"github.com/guestin/mob"
	"github.com/skip2/go-qrcode"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	appId     = ""
	appSecret = ""

	mchId    = ""
	apiKey   = ""
	subAppId = ""
	subMchId = ""

	//mchId    = ""
	//apiKey   = ""
	//subAppId = ""
	//subMchId = ""
)

var sandBox = false

//noinspection ALL
func init() {
	signType := "HMAC-SHA256"
	//signType:="MD5"
	cli = DefaultWxPayClient(appId, mchId, apiKey, appSecret, subAppId, subMchId, signType, sandBox)
	certFile := fmt.Sprintf("./playground/apiclient_cert_%s.p12", mchId)
	err := cli.LoadCertFromFile(certFile)
	if err != nil {
		panic(err)
	}
}

var cli Client = nil

//type notifyApp struct {
//}
//
//func (*notifyApp) GetName() string {
//	return fmt.Sprintf("%s Notify", logTag)
//}
//func (*notifyApp) GetVersion() string {
//	return "1.0"
//}
//
//func (*notifyApp) Bootstrap() {
//	bootx.Web().Use(middleware.BodyDump(middleware.DumpAll))
//	bootx.Web().POST("notify/wxpay/:tradeNo", func(ctx echo.Context) error {
//		tradeNo := ctx.Param("tradeNo")
//		log.Println("notify : ", tradeNo)
//		return ctx.String(200, "Success")
//	})
//}
//
//func (*notifyApp) Shutdown() {
//
//}
func makeNotifyUrl(outTradeNo string) string {
	return fmt.Sprintf("http://home.hiyun.me:13151/notify/wxpay/%s", outTradeNo)
}
func waitForNotify() {
	go func() {
		//bootx.Bootstrap(&notifyApp{}, bootx.WebConfig{Port: 13151})
	}()
}

func stopWaitForNotify() {
	//bootx.Kill()
}

func testPayUnifiedOrder(t *testing.T, cli Client, outTradeNo string, total int64, tradeType string) *PayUnifiedOrderResp {
	req := &PayUnifiedOrderReq{
		BaseReq:        BaseReq{NotifyURL: makeNotifyUrl(outTradeNo)},
		OutTradeNo:     outTradeNo,
		TotalFee:       total,
		Body:           "测试",
		TradeType:      tradeType,
		SpBillCreateIp: "127.0.0.1",
	}
	resp, err := cli.PayUnifiedOrder(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsBusinessSuccess() {
		log.Println(resp.Error())
		t.FailNow()
	}
	return resp
}

func testPayMicroPay(t *testing.T, cli Client, outTradeNo string, total int64, authCode string) {
	req := &PayMicroPayReq{
		OutTradeNo:     outTradeNo,
		TotalFee:       total,
		Body:           "测试",
		SpBillCreateIp: "127.0.0.1",
		AuthCode:       authCode,
	}
	resp, err := cli.PayMicroPay(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsSuccess() {
		log.Println(resp.Error())
		t.FailNow()
	}
}

func testPayOrderQuery(t *testing.T, cli Client, outTradeNo string) *PayOrderQueryResp {
	for {
		time.Sleep(time.Second * 2)
		req := &PayOrderQueryReq{
			OutTradeNo: outTradeNo,
		}
		resp, err := cli.PayOrderQuery(req)
		assert.NoError(t, err)
		if err != nil {
			t.FailNow()
		}
		if !resp.IsBusinessSuccess() {
			log.Println(resp.Error())
			t.FailNow()
		}
		log.Println("trade state : ", resp.TradeState)
		if resp.TradeState != TradeNotPay && resp.TradeState != TradeUserPaying {
			return resp
		}
	}
}

func testPayRefund(t *testing.T, cli Client, outTradeNo string, total, refund int64) {
	outRefundNo := mob.GenRandomUUID()
	req := &PayRefundReq{
		BaseReq:     BaseReq{NotifyURL: makeNotifyUrl(outTradeNo)},
		OutTradeNo:  outTradeNo,
		OutRefundNo: outRefundNo,
		TotalFee:    total,
		RefundFee:   refund,
	}
	resp, err := cli.PayRefund(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsBusinessSuccess() {
		log.Println(resp.Error())
		t.FailNow()
	}
}

func Test_PayUnifiedOrderJSAPI(t *testing.T) {
	waitForNotify()
	defer stopWaitForNotify()
	outTradeNo := mob.GenRandomUUID()
	total := int64(1)
	if sandBox {
		total = 101
	}
	resp := testPayUnifiedOrder(t, cli, outTradeNo, total, TradeTypeJSAPI)
	if resp == nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 10)
	queryResp := testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	if !sandBox {
		testPayRefund(t, cli, outTradeNo, total, total)
	}
}

func Test_PayUnifiedOrderNative(t *testing.T) {
	waitForNotify()
	defer stopWaitForNotify()
	outTradeNo := mob.GenRandomUUID()
	total := int64(1)
	if sandBox {
		total = 101
	}
	resp := testPayUnifiedOrder(t, cli, outTradeNo, total, TradeTypeNative)
	if resp == nil {
		t.FailNow()
	}
	log.Println("qr : ", resp.CodeURL)
	err := internal.OutQr2Console(resp.CodeURL, qrcode.Low, 20)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 10)
	queryResp := testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	if !sandBox {
		testPayRefund(t, cli, outTradeNo, total, total)
	}
}

func Test_PayUnifiedOrderNativeAndRefund(t *testing.T) {
	waitForNotify()
	defer stopWaitForNotify()
	outTradeNo := mob.GenRandomUUID()
	total := int64(2)
	refund := int64(1)
	if sandBox {
		total = 101
	}
	resp := testPayUnifiedOrder(t, cli, outTradeNo, total, TradeTypeNative)
	if resp == nil {
		t.FailNow()
	}
	log.Println("qr : ", resp.CodeURL)
	err := internal.OutQr2Console(resp.CodeURL, qrcode.Low, 20)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 10)
	queryResp := testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	testPayRefund(t, cli, outTradeNo, total, refund)
}

func Test_PayRefund(t *testing.T) {
	waitForNotify()
	defer stopWaitForNotify()
	outTradeNo := "93ada853ceb8485fa88cd8e6f8996198"
	queryResp := testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	if queryResp.TradeState != TradeSuccess {
		return
	}
	total := queryResp.TotalFee
	testPayRefund(t, cli, outTradeNo, total, total)
	queryResp = testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
}

func Test_PayMicroPay(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	authCode := "134695396766870208"
	total := int64(1)
	if sandBox {
		total = 101
	}
	testPayMicroPay(t, cli, outTradeNo, total, authCode)
	queryResp := testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	if !sandBox {
		testPayRefund(t, cli, outTradeNo, total, total)
	}
}

func Test_PayMicroPayAndRefund(t *testing.T) {
	waitForNotify()
	defer stopWaitForNotify()
	outTradeNo := mob.GenRandomUUID()
	total := int64(2)
	if sandBox {
		total = 101
	}
	authCode := ""
	testPayMicroPay(t, cli, outTradeNo, total, authCode)
	queryResp := testPayOrderQuery(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	if queryResp.TradeState != TradeSuccess {
		t.FailNow()
	}
	testPayRefund(t, cli, outTradeNo, total, 1)
}

func testDepositMicroPay(t *testing.T, cli Client, outTradeNo string, total int64, authCode string) {
	req := &DepositMicroPayReq{
		Deposit:        "Y",
		OutTradeNo:     outTradeNo,
		TotalFee:       total,
		Body:           "222",
		SpBillCreateIp: "127.0.0.1",
		AuthCode:       authCode,
		FeeType:        "",
		TimeStart:      "",
		TimeExpire:     "",
		//SubMchId:       mchId,
		//Attach: "11111111",
	}
	resp, err := cli.DepositMicroPay(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsSuccess() {
		log.Println(resp.Error())
		t.FailNow()
	}
}

func testDepositOrderQuery(t *testing.T, cli Client, outTradeNo string) *DepositOrderQueryResp {
	req := &DepositOrderQueryReq{
		OutTradeNo: outTradeNo,
	}
	resp, err := cli.DepositOrderQuery(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsBusinessSuccess() {
		log.Println(resp.Error())
		t.FailNow()
	}
	return resp
}

func waitDepositPay(t *testing.T, cli Client, outTradeNo string) *DepositOrderQueryResp {
	for {
		time.Sleep(time.Second * 2)
		resp := testDepositOrderQuery(t, cli, outTradeNo)
		log.Println("trade state : ", resp.TradeState)
		if resp.TradeState != TradeNotPay && resp.TradeState != TradeUserPaying {
			return resp
		}
	}

}
func testDepositReverse(t *testing.T, cli Client, outTradeNo string) {
	req := &DepositReverseReq{
		OutTradeNo: outTradeNo,
	}
	resp, err := cli.DepositReverse(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsBusinessSuccess() {
		log.Println(resp.Error())
		t.FailNow()
	}
}

func TestClient_DepositMicroPay(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	//outTradeNo:="123456"
	authCode := "134605804298907140"
	total := int64(1)
	if sandBox {
		total = 101
	}
	testDepositMicroPay(t, cli, outTradeNo, total, authCode)
	queryResp := waitDepositPay(t, cli, outTradeNo)
	if queryResp == nil {
		t.FailNow()
	}
	if !sandBox {
		testDepositReverse(t, cli, outTradeNo)
	}
}
func TestClient_DepositOrderQuery(t *testing.T) {
	outTradeNo := "1234567"
	resp := testDepositOrderQuery(t, cli, outTradeNo)
	if !assert.NotNil(t, resp) {
		t.FailNow()
	}
}
