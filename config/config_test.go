package config_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/danbondd/tweet/config"
)

type testDecoder struct {
	r io.Reader
}

func (t testDecoder) Decode(v interface{}) error {
	return errors.New("decode error")
}

func mockDecoderFact(r io.Reader) config.Decoder {
	return testDecoder{r}
}

type mockFileReader struct {
	readErr bool
}

func (m mockFileReader) mockOpen(name string) (io.ReadCloser, error) {
	if m.readErr {
		return nil, errors.New("file corrupt")
	}

	b := []byte(`{
        "ConsumerKey"      : "CONSUMER_KEY",
        "ConsumerSecret"   : "CONSUMER_SECRET",
        "AccessToken"      : "ACCESS_TOKEN",
        "AccessTokenSecret": "ACCESS_TOKEN_SECRET"
    }`)
	return ioutil.NopCloser(bytes.NewReader(b)), nil
}

func TestFileOpenError(t *testing.T) {
	mock := mockFileReader{true}
	_, err := config.New(mock.mockOpen, config.JSONDecoderFactory)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestDecodeError(t *testing.T) {
	mock := mockFileReader{false}
	_, err := config.New(mock.mockOpen, mockDecoderFact)
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}

func TestConfigSuccess(t *testing.T) {
	mock := mockFileReader{false}
	_, err := config.New(mock.mockOpen, config.JSONDecoderFactory)
	if err != nil {
		t.Errorf("expected nil, got: %v", err)
	}
}
