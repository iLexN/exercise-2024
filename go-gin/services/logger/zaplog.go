package logger

import (
	"go.uber.org/zap"
	"os"
)

type ZapLog struct {
	logger *zap.Logger
}

func CreateZapLog() ZapLog {
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(2)
	}
	return ZapLog{
		logger: logger,
	}
}

func (l ZapLog) Info(msg string) {
	s := l.logger.Sugar()
	s.Info(msg)
}

func (l ZapLog) Error(msg string, err errror){
	// todo: 
}
