package log

import (
	"encoding/json"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		in   Level
		want string
	}{
		{INF, `"INF"`},
		{WRN, `"WRN"`},
		{ERR, `"ERR"`},
	}
	for _, tt := range tests {
		b, err := json.Marshal(tt.in)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		have := string(b)
		if have != tt.want {
			t.Fatalf("Level\nhave '%s'\nwant '%s'", have, tt.want)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		in   string
		want Level
	}{
		{`"INF"`, INF},
		{`"WRN"`, WRN},
		{`"ERR"`, ERR},
	}
	for _, tt := range tests {
		var have Level
		err := json.Unmarshal([]byte(tt.in), &have)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if have != tt.want {
			t.Fatalf("Level\nhave '%s'\nwant '%s'", have, tt.want)
		}
	}
}
