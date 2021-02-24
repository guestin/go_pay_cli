package alipay

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strings"
)

type AopParser interface {
	parse(req Request, resp string) (*responseInternal, error)
	unMarshall(req Request, resp *responseInternal, respContent interface{}) error
}

func getAopParser(format string) AopParser {
	if strings.EqualFold(format, "JSON") {
		return &AopJsonParser{}

	} else if strings.EqualFold(format, "XML") {

	}
	panic(fmt.Sprintf("format type '%s' not support yet", format))
}

type AopJsonParser struct {
}

func (this *AopJsonParser) parse(req Request, resp string) (*responseInternal, error) {
	apiName := req.ApiName()
	rootNodeName := strings.Replace(apiName, ".", "_", -1) + kResponseSuffix
	rootIndex := strings.LastIndex(resp, rootNodeName)
	errorIndex := strings.LastIndex(resp, kErrorResponse)
	nullIndex := strings.LastIndex(resp, kNullResponse)
	item := &responseInternal{}
	err := json.Unmarshal([]byte(resp), item)
	if err != nil {
		return nil, errors.Wrap(err, "parse response failed ")
	}
	if rootIndex > 0 {
		srcData, err := this.parseSignSourceData(resp, rootNodeName, rootIndex)
		if err != nil {
			return nil, err
		}
		item.ContentRaw = srcData
	} else if errorIndex > 0 {
		srcData, err := this.parseSignSourceData(resp, kErrorResponse, errorIndex)
		if err != nil {
			return nil, err
		}
		item.IsErrorResp = true
		item.ContentRaw = srcData
	} else if nullIndex > 0 {
		srcData, err := this.parseSignSourceData(resp, kNullResponse, nullIndex)
		if err != nil {
			return nil, err
		}
		item.ContentRaw = srcData
		item.IsErrorResp = true
	} else {
		return nil, errors.New("empty body")
	}
	if item.IsErrorResp {
		codeMsg := &CodeMsg{}
		if err = this.unMarshall(req, item, codeMsg); err != nil {
			return nil, err
		}
		item.errResp = codeMsg
	}
	return item, nil
}

func (this *AopJsonParser) unMarshall(_ Request, resp *responseInternal, respContent interface{}) error {
	return json.Unmarshal([]byte(resp.ContentRaw), respContent)
}

func (this *AopJsonParser) parseSignSourceData(body, nodeName string, indexOfNode int) (string, error) {
	dataStartIndex := indexOfNode + len(nodeName) + 2
	signIndex := strings.LastIndex(body, "\""+kSignNodeName+"\"")
	certIndex := strings.LastIndex(body, "\""+kCertSNNodeName+"\"")
	dataEndIndex := len(body) - 1
	if signIndex > 0 && certIndex > 0 {
		dataEndIndex = int(math.Min(float64(signIndex), float64(certIndex))) - 1
	} else if certIndex > 0 {
		dataEndIndex = certIndex - 1
	} else if signIndex > 0 {
		dataEndIndex = signIndex - 1
	}
	if dataEndIndex-dataStartIndex <= 0 {
		return "", nil
	}
	content := body[dataStartIndex:dataEndIndex]
	return content, nil
}
