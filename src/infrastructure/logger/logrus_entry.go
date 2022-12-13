package logger

import (
	"github.com/sirupsen/logrus"
)

// logrusEntry ...
type logrusEntry struct {
	entry *logrus.Entry
}

// Info ...
func (l *logrusEntry) Info(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Error ...
func (l *logrusEntry) Error(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Debug ...
func (l *logrusEntry) Debug(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Warning ...
func (l *logrusEntry) Warning(format string, args ...interface{}) {
	l.entry.Warningf(format, args...)
}

// WithFields chain method logrus using intermediary struct
// So we have to mimic it here
func (l *logrusEntry) WithFields(fields Fields) LogInterface {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}

	// Return self so can be reused multiple WithFields
	return &logrusEntry{entry: l.entry.WithFields(logrusFields)}
}

// WithError ...
func (l *logrusEntry) WithError(err error) {
	l.entry.WithError(err)
}
