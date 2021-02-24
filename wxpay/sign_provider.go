package wxpay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ooopSnake/assert.go"
	"github.com/pkg/errors"
	"strings"
)

type SignProvider interface {
	Sign(data []byte, key string) (sign string, err error)
	Verify(data []byte, key string, sign string) error
}

func GetSignProvider(signType string) SignProvider {
	assert.Must(len(signType) > 0, "sign type must not be empty")
	switch signType {
	case "MD5":
		return NewMD5SignProvider()
	case "HMAC-SHA256":
		return NewHmacSha256SignProvider()
	}
	panic(fmt.Sprintf("sign type '%s' not support yet", signType))
}

func NewMD5SignProvider() SignProvider {
	return &md5SignProvider{}
}

type md5SignProvider struct {
}

func (this *md5SignProvider) Sign(data []byte, _ string) (sign string, err error) {
	md5Ctx := md5.New()
	_, err = md5Ctx.Write(data)
	if err != nil {
		return
	}
	return strings.ToUpper(hex.EncodeToString(md5Ctx.Sum(nil))), nil
}

func (this *md5SignProvider) Verify(data []byte, key string, sign string) error {
	signCalc, err := this.Sign(data, key)
	if err != nil {
		return err
	}
	if !strings.EqualFold(sign, signCalc) {
		return errors.New(fmt.Sprintf("verify sign '%s' not eq to calc result '%s'", sign, signCalc))
	}
	return nil
}

func NewHmacSha256SignProvider() SignProvider {
	return &hmacSha256Sin{}
}

type hmacSha256Sin struct {
}

func (this *hmacSha256Sin) Sign(data []byte, key string) (sign string, err error) {
	hash := hmac.New(sha256.New, []byte(key))
	_, err = hash.Write(data)
	if err != nil {
		return
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil))), nil
}

func (this *hmacSha256Sin) Verify(data []byte, key string, sign string) error {
	signCalc, err := this.Sign(data, key)
	if err != nil {
		return err
	}
	if !strings.EqualFold(sign, signCalc) {
		return errors.New(fmt.Sprintf("verify sign '%s' not eq to calc result '%s'", sign, signCalc))
	}
	return nil
}
