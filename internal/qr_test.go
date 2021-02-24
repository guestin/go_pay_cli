package internal

import (
	"github.com/skip2/go-qrcode"
	"testing"
)

func TestOutQr2Console(t *testing.T) {
	_ = OutQr2Console("test", qrcode.Medium, 20)
}
