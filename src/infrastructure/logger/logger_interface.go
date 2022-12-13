package logger

// Fields ...
type Fields map[string]interface{}

// LogInterface is the interface used to log messages in different levels
type LogInterface interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Warning(format string, args ...interface{})
	WithFields(fields Fields) LogInterface
	WithError(err error)
}
