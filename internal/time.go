package internal

import (
	"github.com/ooopSnake/assert.go"
	"time"
)

var ZoneAsiaShangHai = time.FixedZone("UTC", int((time.Hour * 8).Seconds()))

func TimeNow() time.Time {
	now := LocalTime(time.Now())
	assert.Must(now.Year() > 2000, "server system time not correct")
	return now
}

func LocalTime(t time.Time) time.Time {
	return t.In(ZoneAsiaShangHai)
}
