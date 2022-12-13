package logger

import (
	"github.com/sirupsen/logrus"
)

// LogrusLogger ...
type logrusLogger struct {
	logger *logrus.Logger
}

func newLogrusLogger(jsonFormat bool, level Level) LogInterface {
	log := logrus.New()
	log.SetLevel(convertLevel(level))
	if jsonFormat {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     true,
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			ForceColors:            true,
			DisableLevelTruncation: true,
			TimestampFormat:        "2006-01-02 15:04:05",
			FullTimestamp:          true,
		})
	}

	return &logrusLogger{logger: log}
}

func convertLevel(l Level) logrus.Level {
	switch l {
	case PanicLevel:
		return logrus.PanicLevel
	case FatalLevel:
		return logrus.FatalLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case WarnLevel:
		return logrus.WarnLevel
	case DebugLevel:
		return logrus.DebugLevel
	case TraceLevel:
		return logrus.TraceLevel
	default:
		// Default info level
		return logrus.InfoLevel
	}
}

// Info ...
func (l *logrusLogger) Info(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Error ...
func (l *logrusLogger) Error(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Debug ...
func (l *logrusLogger) Debug(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Warning ...
func (l *logrusLogger) Warning(format string, args ...interface{}) {
	l.logger.Warningf(format, args...)
}

// WithFields chain method logrus using intermediary struct
// So we have to mimic it here
func (l *logrusLogger) WithFields(fields Fields) LogInterface {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}

	return &logrusEntry{entry: l.logger.WithFields(logrusFields)}
}

// WithError ...
func (l *logrusLogger) WithError(err error) {
	l.logger.WithError(err)
}
