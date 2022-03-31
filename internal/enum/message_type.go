package enum

import (
	"bytes"
	"encoding/json"
)

type MessageType int

const (
	Success MessageType = iota
	Error
)

var toString = map[MessageType]string{
	Success: "success",
	Error:   "error",
}

var toID = map[string]MessageType{
	"Success": Success,
	"Error":   Error,
}

func (m MessageType) String() string {
	switch m {
	case Success:
		return "success"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}

func (m MessageType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[m])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (m MessageType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	m = toID[j]
	return nil
}
