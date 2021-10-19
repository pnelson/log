package log

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

var (
	_ Handler = &defaultHandler{}
	_ Handler = &textHandler{}
	_ Handler = &shellHandler{}
	_ Handler = &minimalShellHandler{}
	_ Handler = &discardHandler{}
)

func TestDefaultHandlerLog(t *testing.T) {
	var buf bytes.Buffer
	var entry Entry
	tests := []struct {
		level Level
		err   error
	}{
		{DBG, io.EOF},
		{INF, io.EOF},
		{WRN, nil},
		{ERR, nil},
	}
	for _, tt := range tests {
		logger := NewLogger(&buf, false)
		logger.Log(tt.level, nil, "test")
		err := json.NewDecoder(&buf).Decode(&entry)
		if err != tt.err {
			t.Fatalf("unexpected error: %v", err)
		}
		buf.Reset()
	}
}

func TestDefaultHandlerLogDBG(t *testing.T) {
	var buf bytes.Buffer
	var entry Entry
	logger := NewLogger(&buf, true)
	logger.Log(DBG, nil, "test")
	err := json.NewDecoder(&buf).Decode(&entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
