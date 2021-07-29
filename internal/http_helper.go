package internal

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func HttpDo(parentCtx context.Context,
	timeout time.Duration,
	cli *http.Client,
	method, url string,
	reqBody string,
	customHeader ...map[string]string) (respBody []byte, err error) {
	start := time.Now()
	httpStatusCode := -1
	defer func() {
		stop := time.Now()
		latency := stop.Sub(start)
		GLogger.Debugf("%s %s %d latency : %s"+
			"request：\n%s"+
			"response：\n%s"+
			"error: \n%v",
			method,
			url,
			httpStatusCode,
			latency.String(),
			reqBody,
			strings.Replace(string(respBody), "\n", "", -1),
			err,
		)
	}()
	var reqBodyReader io.Reader = nil
	if len(reqBody) > 0 {
		reqBodyReader = strings.NewReader(reqBody)
	}
	withTimeout, cancelFunc := context.WithTimeout(parentCtx, timeout)
	defer cancelFunc()
	req, err := http.NewRequestWithContext(withTimeout, method, url, reqBodyReader)
	if err != nil {
		return nil, errors.Wrapf(err, " new http '%s' request failed ", method)
	}
	if len(customHeader) > 0 && len(customHeader[0]) > 0 {
		for k, v := range customHeader[0] {
			req.Header.Set(k, v)
		}
	}
	response, err := cli.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, " http '%s' '%s' failed ", method, url)
	}
	respBody, err = ioutil.ReadAll(response.Body)
	httpStatusCode = response.StatusCode
	if err != nil {
		return nil, errors.Wrapf(err, " http read body failed %s", err.Error())
	}
	return
}
