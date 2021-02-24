package wxpay

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"golang.org/x/crypto/pkcs12"
	"net/url"
)

func URLValueToXML(m url.Values) string {
	var xmlBuffer = &bytes.Buffer{}
	xmlBuffer.WriteString("<xml>")
	for key := range m {
		var value = m.Get(key)
		if len(value) > 0 {
			if key == "total_fee" || key == "refund_fee" || key == "execute_time" {
				xmlBuffer.WriteString("<" + key + ">" + value + "</" + key + ">")
			} else {
				xmlBuffer.WriteString("<" + key + "><![CDATA[" + value + "]]></" + key + ">")
			}
		}
	}
	xmlBuffer.WriteString("</xml>")
	return xmlBuffer.String()
}

func P12ToPem(p12 []byte, password string) (cert tls.Certificate, err error) {
	blocks, err := pkcs12.ToPEM(p12, password)
	if err != nil {
		return cert, err
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err = tls.X509KeyPair(pemData, pemData)
	return cert, err
}
