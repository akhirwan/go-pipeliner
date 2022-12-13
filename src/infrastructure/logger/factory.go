package logger

import "log"

// Logger ...
type Logger struct {
	LogInterface
}

// Level ...
type Level uint32

const (
	// Logrus ...
	Logrus int = iota
	// Zap ...
	Zap
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// NewLoggerFactory ...
func NewLoggerFactory(logger int, level Level, jsonFormat bool) LogInterface {
	switch logger {
	case Logrus:
		return newLogrusLogger(jsonFormat, level)
	default:
		log.Fatalf("")
		return nil
	}
}
