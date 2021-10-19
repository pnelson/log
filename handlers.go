package log

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"
)

// Entry represents a log entry.
type Entry struct {
	Level   Level     `json:"level"`
	Message string    `json:"message"`
	Data    M         `json:"data,omitempty"`
	Time    time.Time `json:"time"`
}

// Handler represents a log formatting handler.
type Handler interface {
	Log(e Entry) error
}

const timeFormat = "15:04:05"

type defaultHandler struct {
	mu    sync.Mutex
	enc   *json.Encoder
	debug bool
}

// NewLogger returns a *Logger that writes JSON encoded logs.
//
// This logger is recommended for production. Entries with log
// level INF are discarded as they are intended to be purely
// informational and machine-actionable events. Entries with log
// level DBG will be logged if debug is true.
func NewLogger(w io.Writer, debug bool) *Logger {
	return New(&defaultHandler{enc: json.NewEncoder(w), debug: debug})
}

// Log implements the Handler interface.
func (h *defaultHandler) Log(e Entry) error {
	if e.Level == INF || (e.Level == DBG && !h.debug) {
		return nil
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.enc.Encode(e)
}

type textHandler struct {
	mu sync.Mutex
	w  io.Writer
}

// NewTextLogger returns a *Logger that writes text formatted logs.
func NewTextLogger(w io.Writer) *Logger {
	return New(&textHandler{w: w})
}

// Log implements the Handler interface.
func (h *textHandler) Log(e Entry) error {
	ts := e.Time.Local().Format(timeFormat)
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Fprintf(h.w, "%s %s %s", e.Level, ts, e.Message)
	for _, k := range e.Data.Keys() {
		fmt.Fprintf(h.w, " %s: %v", k, e.Data[k])
	}
	fmt.Fprintln(h.w)
	return nil
}

type shellHandler struct {
	mu sync.Mutex
	w  io.Writer
}

// NewShellLogger returns a *Logger that writes POSIX shell formatted logs.
func NewShellLogger(w io.Writer) *Logger {
	return New(&shellHandler{w: w})
}

// Log implements the Handler interface.
func (h *shellHandler) Log(e Entry) error {
	ts := e.Time.Local().Format(timeFormat)
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Fprintf(h.w, "\033[1;%dm%s\033[0m \033[1;30m%s \033[1;37m%s\033[0m", e.Level, e.Level, ts, e.Message)
	for _, k := range e.Data.Keys() {
		fmt.Fprintf(h.w, " \033[1;%dm%s:\033[0m %v", e.Level, k, e.Data[k])
	}
	fmt.Fprintln(h.w)
	return nil
}

type minimalShellHandler struct {
	mu sync.Mutex
	w  io.Writer
}

// NewMinimalShellLogger returns a *Logger that writes text formatted logs
// with decorative shell colors but without timestamps and text levels.
func NewMinimalShellLogger(w io.Writer) *Logger {
	return New(&minimalShellHandler{w: w})
}

// Log implements the Handler interface.
func (h *minimalShellHandler) Log(e Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	fmt.Fprintf(h.w, "\033[1;%dm•\033[0m \033[1;37m%s\033[0m", e.Level, e.Message)
	for _, k := range e.Data.Keys() {
		fmt.Fprintf(h.w, " \033[1;%dm%s:\033[0m %v", e.Level, k, e.Data[k])
	}
	fmt.Fprintln(h.w)
	return nil
}

type discardHandler struct{}

// NewDiscardLogger returns a *Logger that discards log entries.
func NewDiscardLogger() *Logger {
	return New(&discardHandler{})
}

// Log implements the Handler interface.
func (h *discardHandler) Log(e Entry) error {
	return nil
}
