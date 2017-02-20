package utils

import (
	"io"
	"os"
)

// OpenFile is a custom type for os.Open().
type OpenFile func(name string) (io.ReadCloser, error)

// FileReader opens the named file for reading.
func FileReader(name string) (io.ReadCloser, error) {
	return os.Open(name)
}
