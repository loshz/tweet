package config_test

import (
	"errors"
	"os"
	"testing"

	"github.com/danbondd/tweet/config"
)

func mockOpen(name string) (*os.File, error) {
	return nil, errors.New("file corrupt")
}

func TestFileOpenError(t *testing.T) {
	_, err := config.New(mockOpen)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}
