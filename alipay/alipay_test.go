package alipay

import (
	"github.com/guestin/go_pay_cli/internal"
	"github.com/guestin/mob"
	"github.com/skip2/go-qrcode"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

const (
	sandBoxAppId        = "2016101900724251"
	sandBoxPrivateKey   = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDjHd4qw0R3OdlCvFgUizhCQ8oo+70G+dpLaOt/2kENOFflKghwcuUQ+4nWpOn+nUGWWcpeYGrr8CdWNDw2aIhzDAyRsMUUh7ViMMC8HA+Esprdm7UsrP9l6GObRDQyY3xJwD3KYQqD118w7sEudHjohw1kRITDtWLuU3Ffbl4TSV7RVH/O+wDZ7SnsakQ3xAIni5dqDHmw4Dd5i7IUUG31gd59WC39M/lBKk1XSRVovShk9XAUrL0/W9hzydqgAS/9dFFQu4Ln3UxQ5hC8RVQmXBSytmZJzYPlFxScCpVtSWSJfJBvWYqNHmzpsrvVaYAlGOjpMZVXOgmrzlTXmnkdAgMBAAECggEAF67bpfXqw8wCfdUKEkpaOX68K/3kPj/7pXVxaUmnEuvXLoxtiNxSSq0QOJPF2sknN7hxQ9omDChk0bZsuPe5ktWk0eRvCK8GGREgZ/09GQdO4uEDyX7YuxW6nUxFbSO2qDIlv17TK+Bfisi62E9I9GZw3Q3QEmBtypBk+CCYKsZ/iAJ0OKMDmQz5Y/gcR578iWFsPfdsGPxK/vCZKIEzxSQDkCMtfFpWUQn8sqwdkoMYSfgiKilqFyjSVQZoTCtdnh8/EvmZMreB1fF7WnZYcdDATG4A5qvZyobRGkztYgeoaWV9NKTpw/GB80iiem2mduxiVaapstwNrKahk79suQKBgQD8HuwfJI+fXVGKnagalmlym2aW7CUur0LRhE9z6VPWNJu3K4D0clGsMZ3A2DEFlzXRlY1Sr8LuGTyVArX7ab9J0vWzSYepGkCwS5gD+77NrSu0Qxc6mp5wNXtdgjgxBaw/3MEjOtajykURB7yW38RaWFNZGKVn7LmMVKmXhcC5QwKBgQDmnHT0l0TJ2OMnZq0tAwBaCokImyJ6HYuRcKSaILWPDMz0SzhVp6jjOQsF12TFLNSgRDC+iPciqaWJzACRbIMDGqd6gguczo8tnCS4eD1LgXXSpODTQQYnTP3MfGow5jBX9II9jSL9NHRL5Sk/l1EN9Dj3M8dQMlJRtfrtB34uHwKBgQC27bzG7+EhcTUbzT5OZDoIRMbP1HE6CUIDAOwhHveMYUlmQrNjKZsmxC1A4dvXwZn0An2ytAJMfZUeTQQ7ccOwTdemCUDcKkcrYv3eTgdn9jDSrycoh01T/woOk9AviX0sLQEZjbR0zOsF60YjdiJiptl6uM4ytGkAb+FJJmvqPwKBgDBj9Ea+1zhjwoaqDPy8/H3oaAjeRMXLHVZPhLqy0mZKEVfR0OhoXhAQEDgRkputZJCcvn28z97+KjZYEGZzlqo4FZynXThyP6kacroiwPnvGIIzBtpNcrUcesVF9iJ8qvhJ3mp8CzOGpkCmvZkSb3e2H53/x3sUlCGBRj4mrFuRAoGBALwFtuRiIEwNRsd82x5QverXDJGsLls4LTZm3shfoZUtNnkvlWxer9SbxRGb7APsToAessgZ14RrVTgoztTKzedOkHIuR3P9IO5aSZeoC0FZg3vjmiNFhrV5dVuHvXTNn6i7Q35ab8m+eqEcJ9YeihBbavEMU04zUB5NGugtt8K+"
	sandBoxAliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyFYBSmUg64eWzagT+mY17QotdpQIapL+mn6v9KJxOWVWsTq4ChIwDMTtSYNr7i2f5q5xJCVX8c0Fb8otKQlMGsY7fO/7xNLRwtKzWbsi8+gLB59qZKBuQ6lDefQkCO5MrEufLC8euYXGp9J08m5oHscqxfN3ohUIklNEDE0l9DA4knp3rBWAGVunsDrivvbg0/ZfMAcXJXt34J9Zfn7UrU54wMx9jxNp5CNwJKdi91mxYHUOxp5YEBqnSVFjmANXA68YlNR2yc/NgPK5PZaEjPq1LtkkiPEKfbYN3aXTr1CgE2Dy1hh81UW9lvQ9r9dPk5wkaZ/AztZj752Z86qzBQIDAQAB"

	appId        = ""
	privateKey   = ""
	aliPublicKey = ""
)

var cli Client = nil
var gBuyerId = ""
var gSellerId = ""
var isvPid = ""

var extendPrams = &ExtendParams{
	//SysServiceProviderId: isvPid,
}

//noinspection ALL
func init() {
	sandBox := true
	if sandBox {
		gBuyerId = "2088102172292710"
		gSellerId = "2088102180300847"
		cli = DefaultAliPayClient(sandBoxAppId, sandBoxPrivateKey, "RSA2", sandBox)
		cli.LoadAliPayPublicKey(sandBoxAliPublicKey)
	} else {
		gBuyerId = ""
		gSellerId = isvPid
		cli = DefaultAliPayClient(appId, privateKey, "RSA2", sandBox)
		cli.LoadAliPayPublicKey(aliPublicKey)
		//cli.SetAppAuthToken("")
	}
}

func testCreate(t *testing.T, cli Client, outTradeNo string, total string) {
	req := &TradeCreateReq{
		Subject:      "支付测试(JSAPI)",
		OutTradeNo:   outTradeNo,
		TotalAmount:  total,
		BuyerId:      gBuyerId,
		ExtendParams: extendPrams,
	}
	rsp, err := cli.TradeCreate(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !rsp.IsCodeSuccess() {
		t.FailNow()
	}
}

func testPreCreate(t *testing.T, cli Client, outTradeNo string, total string) {
	req := &TradePreCreateReq{
		OutTradeNo:   outTradeNo,
		TotalAmount:  total,
		Subject:      "支付测试(C扫B)",
		ExtendParams: extendPrams,
	}
	resp, err := cli.TradePreCreate(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
	t.Log("qr ", resp.QrCode)
	err = internal.OutQr2Console(resp.QrCode, qrcode.Low, 20)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
}

func testTradePay(t *testing.T, cli Client, outTradeNo string, total string, authCode string) {
	req := &TradePayReq{
		Subject:      "支付测试(B扫C)",
		OutTradeNo:   outTradeNo,
		TotalAmount:  total,
		AuthCode:     authCode,
		ExtendParams: extendPrams,
	}
	resp, err := cli.TradePay(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func testTradeQuery(t *testing.T, cli Client, outTradeNo string) {
	for {
		time.Sleep(time.Second * 2)
		req := &TradeQueryReq{
			OutTradeNo: outTradeNo,
		}
		resp, err := cli.TradeQuery(req)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		if resp.IsCodeSuccess() {
			log.Println("trade status = ", resp.TradeStatus)
			if resp.TradeStatus != TradeWaitBuyerPay {
				break
			}
		}
	}
}

func testTradeRefund(t *testing.T, cli Client, outTradeNo string, refundAmount string) {
	req := &TradeRefundReq{
		OutTradeNo:   outTradeNo,
		RefundAmount: refundAmount,
		OutRequestNo: mob.GenRandomUUID(),
	}
	resp, err := cli.TradeRefund(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func testTradeCancel(t *testing.T, cli Client, outTradeNo string) {
	req := &TradeCancelReq{
		OutTradeNo: outTradeNo,
		TradeNo:    "",
	}
	resp, err := cli.TradeCancel(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
	log.Println("retry flag = ", resp.RetryFlag)
}

func testTradeClose(t *testing.T, cli Client, outTradeNo string) {
	req := &TradeCloseReq{
		OutTradeNo: outTradeNo,
		TradeNo:    "",
	}
	resp, err := cli.TradeClose(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func Test_TradeCreate(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testCreate(t, cli, outTradeNo, "0.01")
	testTradeQuery(t, cli, outTradeNo)
}

func Test_TradeCreateAndRefund(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testCreate(t, cli, outTradeNo, "0.01")
	testTradeQuery(t, cli, outTradeNo)
	testTradeRefund(t, cli, outTradeNo, "0.01")
}

func Test_TradeCreateAndRefundPart(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testCreate(t, cli, outTradeNo, "0.03")
	testTradeQuery(t, cli, outTradeNo)
	testTradeRefund(t, cli, outTradeNo, "0.02")
}

func Test_TradePreCreate(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testPreCreate(t, cli, outTradeNo, "0.01")
	testTradeQuery(t, cli, outTradeNo)
}

func Test_TradeRefund(t *testing.T) {
	outTradeNo := "1eded1cfbd7e49b6b2cb5cbd8d9b242e"
	testTradeRefund(t, cli, outTradeNo, "0.01")
}

func TestCreateAndPay(t *testing.T) {
	//这种方式不行
	outTradeNo := mob.GenRandomUUID()
	authCode := "288787000093878028"
	testCreate(t, cli, outTradeNo, "0.01")
	testTradePay(t, cli, outTradeNo, "0.01", authCode)
	testTradeQuery(t, cli, outTradeNo)
}

func TestPreCreateAndPay(t *testing.T) {
	//这种方式可以
	outTradeNo := mob.GenRandomUUID()
	authCode := "285839657145583283"
	testPreCreate(t, cli, outTradeNo, "0.01")
	testTradePay(t, cli, outTradeNo, "0.01", authCode)
	testTradeQuery(t, cli, outTradeNo)
}

func Test_TradePay(t *testing.T) {
	authCode := "282607754913260093"
	outTradeNo := mob.GenRandomUUID()
	testTradePay(t, cli, outTradeNo, "0.01", authCode)
	testTradeQuery(t, cli, outTradeNo)
}

func Test_TradeCreateAndCancel(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testCreate(t, cli, outTradeNo, "0.01")
	time.Sleep(time.Second * 5)
	testTradeCancel(t, cli, outTradeNo)
	testTradeQuery(t, cli, outTradeNo)
}

func Test_TradeCreateAndClose(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testCreate(t, cli, outTradeNo, "0.01")
	time.Sleep(time.Second * 5)
	testTradeClose(t, cli, outTradeNo)
	testTradeQuery(t, cli, outTradeNo)
}

func Test_TradePreCreateAndCancel(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testPreCreate(t, cli, outTradeNo, "0.01")
	time.Sleep(time.Second * 5)
	testTradeCancel(t, cli, outTradeNo)
}

func Test_TradePreCreateAndClose(t *testing.T) {
	outTradeNo := mob.GenRandomUUID()
	testPreCreate(t, cli, outTradeNo, "0.01")
	time.Sleep(time.Second * 10)
	testTradeClose(t, cli, outTradeNo)
}

func testFundAuthOrderVoucherCreate(t *testing.T, cli Client, outOrderNo, outRequestNo string, amount string) *FundAuthOrderVoucherCreateResp {
	req := &FundAuthOrderVoucherCreateReq{
		OutOrderNo:   outOrderNo,
		OutRequestNo: outRequestNo,
		OrderTitle:   "预授权测试(C扫B)",
		Amount:       amount,
		ProductCode:  ProductCodePreAuth,
		ExtraParam:   "{\"category\":\"HOTEL\"}",
	}
	resp, err := cli.FundAuthOrderVoucherCreate(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
	log.Println("qr ", resp.CodeValue)
	err = internal.OutQr2Console(resp.CodeValue, qrcode.Low, 20)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	return resp
}

func testFundAuthOrderFreeze(t *testing.T, cli Client, authCode, outOrderNo, outRequestNo string, amount string) {
	req := &FundAuthOrderFreezeReq{
		AuthCode:     authCode,
		AuthCodeType: AuthCodeTypeBarCode,
		OutOrderNo:   outOrderNo,
		OutRequestNo: outRequestNo,
		OrderTitle:   "预授权测试(B扫C)",
		Amount:       amount,
		ExtraParam:   "{\"category\":\"HOTEL\"}",
	}
	resp, err := cli.FundAuthOrderFreeze(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func testFundAuthOrderAppFreeze(t *testing.T, cli Client, outOrderNo, outRequestNo string, amount string) {
	req := &FundAuthOrderAppFreezeReq{
		OutOrderNo:   outOrderNo,
		OutRequestNo: outRequestNo,
		OrderTitle:   "预授权测试(JSAPI)",
		Amount:       amount,
		ExtraParam:   "{\"category\":\"HOTEL\"}",
	}
	resp, err := cli.FundAuthOrderAppFreeze(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func testTradePayForFundAuthComplete(t *testing.T, cli Client,
	outTradeNo, total, authNo, buyerId, sellerId string) {
	req := &TradePayReq{
		Subject:         "预授权转支付测试",
		OutTradeNo:      outTradeNo,
		TotalAmount:     total,
		AuthNo:          authNo,
		ProductCode:     ProductCodePreAuth,
		AuthConfirmMode: AuthConfirmModeComplete,
		BuyerId:         buyerId,
		SellerId:        sellerId,
		ExtendParams:    extendPrams,
	}
	resp, err := cli.TradePay(req)
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func testFundQuery(t *testing.T, cli Client, outOrderNo, ouRequestNo string) *FundAuthOrderOperationDetailQueryResp {
	for {
		time.Sleep(time.Second * 2)
		req := &FundAuthOrderOperationDetailQueryReq{
			AuthNo:       "",
			OutOrderNo:   outOrderNo,
			OperationId:  "",
			OutRequestNo: ouRequestNo,
		}
		resp, err := cli.FundAuthOrderOperationDetailQuery(req)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		if resp.IsCodeSuccess() {
			log.Println("status = ", resp.Status)
			if resp.Status != FundStatusInit {
				return resp
			}
		}
	}
}

func testFundAuthOrderUnFreeze(t *testing.T, cli Client, authNo, outRequestNo, amount, remark string) {
	req := &FundAuthOrderUnFreezeReq{
		AuthNo:       authNo,
		OutRequestNo: outRequestNo,
		Amount:       amount,
		Remark:       remark,
		ExtraParam:   "",
	}
	resp, err := cli.FundAuthOrderUnFreeze(req)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	if !resp.IsCodeSuccess() {
		t.FailNow()
	}
}

func Test_FundAuthOrderVoucherCreate(t *testing.T) {
	outOrderNo := mob.GenRandomUUID()
	outRequestNo := mob.GenRandomUUID()
	total := "0.01"
	resp := testFundAuthOrderVoucherCreate(t, cli, outOrderNo, outRequestNo, total)
	if resp == nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 10)
	_ = testFundQuery(t, cli, outOrderNo, outRequestNo)
}

func TestClient_FundAuthOrderUnFreeze(t *testing.T) {
	authNo := "2020020710002001140505194376"
	outOrderNo := "c122dcb14fe44a4d97a5dbcb38292a7a"
	amount := "0.02"
	testFundAuthOrderUnFreeze(t, cli, authNo, outOrderNo, amount, "解冻")
}

func Test_FundAuth2Pay(t *testing.T) {
	authNo := "2020020710002001140505201669"
	outTradeNo := mob.GenRandomUUID()
	total := "0.01"
	buyerId := "2088402588269149"
	testTradePayForFundAuthComplete(t, cli, outTradeNo, total, authNo, buyerId, gSellerId)
}

func TestClient_FundAuthOrderVoucherCreateAndUnFreeze(t *testing.T) {
	outOrderNo := mob.GenRandomUUID()
	outRequestNo := mob.GenRandomUUID()
	total := "0.01"
	resp := testFundAuthOrderVoucherCreate(t, cli, outOrderNo, outRequestNo, total)
	if resp == nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 10)
	fundQueryResp := testFundQuery(t, cli, outOrderNo, outRequestNo)
	if fundQueryResp == nil {
		t.FailNow()
	}
	testFundAuthOrderUnFreeze(t, cli,
		fundQueryResp.AuthNo, outOrderNo,
		fundQueryResp.Amount, "解冻")

}

func TestClient_FundAuthOrderVoucherCreateAndComplete(t *testing.T) {
	outOrderNo := mob.GenRandomUUID()
	outRequestNo := mob.GenRandomUUID()
	total := "0.02"
	resp := testFundAuthOrderVoucherCreate(t, cli, outOrderNo, outRequestNo, total)
	if resp == nil {
		t.FailNow()
	}
	time.Sleep(time.Second * 10)
	fundQueryResp := testFundQuery(t, cli, outOrderNo, outRequestNo)
	if fundQueryResp == nil {
		t.FailNow()
	}
	outTradeNo := mob.GenRandomUUID()
	testTradePayForFundAuthComplete(t, cli,
		outTradeNo, "0.01",
		fundQueryResp.AuthNo,
		fundQueryResp.PayerUserId, gSellerId)

}

func Test_FundAuthOrderFreeze(t *testing.T) {
	outOrderNo := mob.GenRandomUUID()
	outRequestNo := mob.GenRandomUUID()
	authCode := ""
	total := "0.01"
	testFundAuthOrderFreeze(t, cli, authCode, outOrderNo, outRequestNo, total)
	time.Sleep(time.Second * 10)
	_ = testFundQuery(t, cli, outOrderNo, outRequestNo)
}

func Test_FundAuthOrderAppFreeze(t *testing.T) {
	outOrderNo := mob.GenRandomUUID()
	outRequestNo := mob.GenRandomUUID()
	total := "0.01"
	testFundAuthOrderAppFreeze(t, cli, outOrderNo, outRequestNo, total)
	_ = testFundQuery(t, cli, outOrderNo, outRequestNo)
}
