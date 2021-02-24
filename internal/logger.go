package internal

import "fmt"

var GLogger Logger = &_logger{}

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
}

type _logger struct {
}

func (this *_logger) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (this *_logger) Debugf(template string, args ...interface{}) {
	fmt.Printf(template, args...)
}
