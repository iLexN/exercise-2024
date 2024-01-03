package logger

var DefaultLogger Logger

type Logger interface {
	Info(msg string)
}

// The init() function in Go is a special function that is
// automatically called before the program starts.
// It is typically used for initialization tasks,
// such as setting up global variables, configuring packages,
func init() {
	//	DefaultLogger = CreateZapLog()
	DefaultLogger = CreateZeroLog()
}
