package services

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var DefaultLogger Logger

type Logger interface {
	Info(msg string)
}

func init() {
	var logger ZeroLog
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
