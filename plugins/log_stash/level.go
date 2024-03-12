package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
)

func (this Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this Level) String() string {
	var s string
	switch this {
	case DebugLevel:
		s = "debug"
	case InfoLevel:
		s = "info"
	case WarnLevel:
		s = "warn"
	case ErrorLevel:
		s = "error"
	default:
		s = "unknown"
	}
	return s
}
