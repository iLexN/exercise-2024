package logger

var DefaultLogger Logger

type Logger interface {
	Info(msg string)
}

func init() {
	//	DefaultLogger = CreateZapLog()
	DefaultLogger = CreateZeroLog()
}
