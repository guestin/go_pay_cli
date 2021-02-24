package internal

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func OutQr2Console(content string, level qrcode.RecoveryLevel, size int) error {
	qr, err := qrcode.New(content, level)
	if err != nil {
		return err
	}
	img := qr.Image(size)
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			r32, g32, b32, _ := img.At(x, y).RGBA()
			r, g, b := int(r32>>8), int(g32>>8), int(b32>>8)
			if (r+g+b)/3 > 180 {
				fmt.Print("\033[47;30m  \033[0m")
			} else {
				fmt.Print("\033[40;40m  \033[0m")
			}
		}
		fmt.Println()
	}
	return nil
}
