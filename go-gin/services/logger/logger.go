package logger

var DefaultLogger Logger

// should i use other interface
// some info:
// https://github.com/go-logr/logr
// https://www.reddit.com/r/golang/comments/6bptkl/why_isnt_there_a_logger_interface_in_the_standard/?rdt=36955
// https://github.com/golang/go/issues/28412
// https://medium.com/@ansujain/building-a-logger-wrapper-in-go-with-support-for-multiple-logging-libraries-48092b826bee
type Logger interface {
	Info(msg string)
	Error(msg string, err error)
}

// The init() function in Go is a special function that is
// automatically called before the program starts.
// It is typically used for initialization tasks,
// such as setting up global variables, configuring packages,
func init() {
	DefaultLogger = CreateZapLog()
	//	DefaultLogger = CreateZeroLog()
}
