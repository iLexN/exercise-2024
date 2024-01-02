package services

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"go.uber.org/zap"
)

var DefaultLogger Logger

type Logger interface {
	Info(msg string)
}

func init() {
	//	var logger ZeroLog
	//	logger.init()

	var logger ZapLog
	logger.init()

	DefaultLogger = logger
}

type ZeroLog struct {
}

func (l ZeroLog) init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (l ZeroLog) Info(msg string) {
	log.Info().Msg(msg)
}

type ZapLog struct {
	logger *zap.Logger
}

func (l *ZapLog) init() {
	logger, err := zap.NewProduction()
	if err != nil {
		return
	}

	l.logger = logger
}

func (l *ZapLog) Info(msg string) {
	s := l.logger.Sugar()
	s.Info(msg)
}
