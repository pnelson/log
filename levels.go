package log

import (
	"bytes"
	"errors"
	"strings"
)

// Level represents a log level.
type Level int

// Level constants.
const (
	// DBG represents debug data.
	DBG Level = 30

	// INF represents machine-actionable data.
	INF = 34

	// WRN represents human-observable data.
	WRN = 33

	// ERR represents human-actionable data.
	ERR = 31
)

// levels represents the levels as three character strings.
var levels = [...]string{
	DBG: "DBG",
	INF: "INF",
	WRN: "WRN",
	ERR: "ERR",
}

// reverse represents a lookup table from three
// character strings to Level constants.
var reverse = map[string]Level{
	levels[DBG]: DBG,
	levels[INF]: INF,
	levels[WRN]: WRN,
	levels[ERR]: ERR,
}

// String implements the fmt.Stringer interface.
func (l Level) String() string {
	return levels[l]
}

// MarshalJSON impelements the json.Marshaler interface.
func (l Level) MarshalJSON() ([]byte, error) {
	return []byte(`"` + levels[l] + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Level) UnmarshalJSON(b []byte) error {
	v, ok := reverse[strings.ToUpper(string(bytes.Trim(b, `"`)))]
	if !ok {
		return errors.New("log: invalid level")
	}
	*l = v
	return nil
}
