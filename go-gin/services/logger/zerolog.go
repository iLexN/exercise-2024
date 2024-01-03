package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ZeroLog struct{}

func (l ZeroLog) Info(msg string) {
	log.Info().Msg(msg)
}

func CreateZeroLog() ZeroLog {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return ZeroLog{}
}
