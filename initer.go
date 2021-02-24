package go_pay_cli

import (
	"github.com/guestin/go_pay_cli/internal"
	"github.com/guestin/log"
	"go.uber.org/zap"
)

//goland:noinspection ALL
func SetupLogger(zapLogger *zap.Logger, loggerOpt ...log.Opt) {
	internal.GLogger = log.NewTaggedClassicLogger(zapLogger, "Go-Pay-Cli", loggerOpt...)
}
