package helpers

import (
	"encoding/json"
	"io"
)

// Decoder specifies the behaviour of a given decoder.
// It only implements the Decode method which reads the next JSON-encoded
// value from its input and stores it in the value pointed to by v.
type Decoder interface {
	Decode(v interface{}) error
}

// JSONDecoder defines a custom decoder function signature.
type JSONDecoder func(r io.Reader) Decoder

// NewJSONDecoder returns a new JSON Decoder.
func NewJSONDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
