// Package log implements structured logging helpers.
package log

import (
	"fmt"
	"os"
	"sort"
	"time"
)

// Logger represents a logger.
type Logger struct {
	h Handler
}

// New returns a new *Logger.
func New(h Handler) *Logger {
	return &Logger{h: h}
}

// Log writes the formatted log entry.
func (l *Logger) Log(level Level, data M, format string, args ...interface{}) {
	e := Entry{
		Level:   level,
		Message: fmt.Sprintf(format, args...),
		Data:    data,
		Time:    time.Now().UTC(),
	}
	err := l.h.Log(e)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// M represents the logging data map.
type M map[string]interface{}

// Keys returns the sorted data keys.
func (m M) Keys() []string {
	var names []string
	for k := range m {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}
