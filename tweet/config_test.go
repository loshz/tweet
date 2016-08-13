package tweet_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/danbondd/tweet/tweet"
)

type testDecoder struct {
	r io.Reader
}

func (t testDecoder) Decode(v interface{}) error {
	return errors.New("decode error")
}

func mockDecoderFactory(r io.Reader) tweet.Decoder {
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
	_, err := tweet.NewConfig(mock.mockOpen, tweet.JSONDecoderFactory)
	if err == nil {
		t.Error("expected file open error, got: nil")
	}
}

func TestDecodeError(t *testing.T) {
	mock := mockFileReader{false}
	_, err := tweet.NewConfig(mock.mockOpen, mockDecoderFactory)
	if err == nil {
		t.Error("expected file decode error, got: nil")
	}
}

func TestConfigSuccess(t *testing.T) {
	mock := mockFileReader{false}
	_, err := tweet.NewConfig(mock.mockOpen, tweet.JSONDecoderFactory)
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}
