package alipay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/ooopSnake/assert.go"
	"github.com/pkg/errors"
	"strings"
)

const (
	kPublicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	kPublicKeySuffix = "-----END PUBLIC KEY-----"

	kPKCS1Prefix = "-----BEGIN RSA PRIVATE KEY-----"
	KPKCS1Suffix = "-----END RSA PRIVATE KEY-----"

	kPKCS8Prefix = "-----BEGIN PRIVATE KEY-----"
	KPKCS8Suffix = "-----END PRIVATE KEY-----"
)

func GetEncryptor(signType string) AsymmetricEncryptor {
	if len(signType) == 0 {
		panic("sign type must not be empty")
	}
	assert.Must(len(signType) > 0, "sign type must not be empty")
	switch signType {
	case "RSA":
		return NewRSAEncryptor()
	case "RSA2":
		return NewRSA2Encryptor()
	}
	panic(fmt.Sprintf("sign type '%s' not support yet", signType))
}

type AsymmetricEncryptor interface {
	Sign(data []byte, privateKey *rsa.PrivateKey) (sign []byte, err error)
	Verify(data []byte, publicKey *rsa.PublicKey, sign string) error
	Encrypt(plain []byte, publicKey string) (cipher []byte, err error)
	Decrypt(cipher []byte, privateKey string) (plain []byte, err error)
}

type RSAEncryptor struct {
	hash crypto.Hash
}

func NewRSAEncryptor() AsymmetricEncryptor {
	return &RSAEncryptor{hash: crypto.SHA1}
}

func NewRSA2Encryptor() AsymmetricEncryptor {
	return &RSAEncryptor{hash: crypto.SHA256}
}

func (this *RSAEncryptor) packageData(rawData []byte, packageSize int) (r [][]byte) {
	var src = make([]byte, len(rawData))
	copy(src, rawData)

	r = make([][]byte, 0)
	if len(src) <= packageSize {
		return append(r, src)
	}
	for len(src) > 0 {
		var p = src[:packageSize]
		r = append(r, p)
		src = src[packageSize:]
		if len(src) <= packageSize {
			r = append(r, src)
			break
		}
	}
	return r
}

func (this *RSAEncryptor) Sign(data []byte, privateKey *rsa.PrivateKey) (sign []byte, err error) {
	hash := this.hash
	var h = hash.New()
	h.Write(data)
	var hashed = h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
}

func (this *RSAEncryptor) Verify(data []byte, publicKey *rsa.PublicKey, sign string) error {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	hash := this.hash
	var h = hash.New()
	h.Write(data)
	var hashed = h.Sum(nil)
	return rsa.VerifyPKCS1v15(publicKey, hash, hashed, signBytes)
}

func (this *RSAEncryptor) Encrypt(plain []byte, publicKey string) (cipher []byte, err error) {
	pubKey, err := parserPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	var data = this.packageData(plain, pubKey.N.BitLen()/8-11)
	cipher = make([]byte, 0, 0)
	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, pubKey, d)
		if e != nil {
			return nil, e
		}
		cipher = append(cipher, c...)
	}
	return cipher, nil
}

func (this *RSAEncryptor) Decrypt(cipher []byte, privateKey string) (plain []byte, err error) {
	key, err := parserPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	var data = this.packageData(cipher, key.PublicKey.N.BitLen()/8)
	plain = make([]byte, 0, 0)
	for _, d := range data {
		var p, e = rsa.DecryptPKCS1v15(rand.Reader, key, d)
		if e != nil {
			return nil, e
		}
		plain = append(plain, p...)
	}
	return plain, nil
}

func formatKey(raw, prefix, suffix string, lineCount int) []byte {
	if raw == "" {
		return nil
	}
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)

	var sl = len(raw)
	var c = sl / lineCount
	if sl%lineCount > 0 {
		c = c + 1
	}
	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * lineCount
		var e = b + lineCount
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}

func formatPublicKey(raw string) []byte {
	return formatKey(raw, kPublicKeyPrefix, kPublicKeySuffix, 64)
}

func formatPKCS1PrivateKey(raw string) []byte {
	raw = strings.Replace(raw, kPKCS8Prefix, "", 1)
	raw = strings.Replace(raw, KPKCS8Suffix, "", 1)
	return formatKey(raw, kPKCS1Prefix, KPKCS1Suffix, 64)
}

func formatPKCS8PrivateKey(raw string) []byte {
	raw = strings.Replace(raw, kPKCS1Prefix, "", 1)
	raw = strings.Replace(raw, KPKCS1Suffix, "", 1)
	return formatKey(raw, kPKCS8Prefix, KPKCS8Suffix, 64)
}

func parsePKCS1PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("invalid PKCS1 private key")
	}
	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid PKCS1 private key")
	}
	return key, err
}

func parsePKCS8PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, errors.New("invalid PKCS8 private key")
	}
	rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid PKCS8 private key")
	}
	key, ok := rawKey.(*rsa.PrivateKey)
	if ok == false {
		return nil, errors.New("invalid PKCS8 private key")
	}
	return key, err
}

func parserPrivateKey(raw string) (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	key, err := parsePKCS1PrivateKey(formatPKCS1PrivateKey(raw))
	if err != nil {
		key, err = parsePKCS8PrivateKey(formatPKCS8PrivateKey(raw))
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func parserPublicKey(raw string) (key *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(formatPublicKey(raw))
	if block == nil {
		return nil, errors.Wrap(err, "invalid rsa public key")
	}
	var pub interface{}
	pub, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid rsa public key")
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.Wrap(err, "invalid rsa public key")
	}
	return pubKey, err
}
