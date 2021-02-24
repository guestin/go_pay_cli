package alipay

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"github.com/pkg/errors"
)

func ParseCertificate(b []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("invalid x509 certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid x509 certificate")
	}
	return cert, nil
}

func getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}

func VerifyCertChain(cert *x509.Certificate, rootsPem []byte) error {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(rootsPem)
	_, err := cert.Verify(x509.VerifyOptions{Roots: certPool})
	if err != nil {
		return err
	}
	return nil
}
