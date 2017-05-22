package tweet_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/danbondd/tweet"
)

func tmpFile(json string) (*os.File, error) {
	data := []byte(json)
	tmpfile, err := ioutil.TempFile("", "config")
	if err != nil {
		return nil, err
	}

	if _, err := tmpfile.Write(data); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}

func TestNewConfig(t *testing.T) {
	t.Run("TestInvalidConfigFile", func(t *testing.T) {
		_, err := tweet.NewConfig("", "invalid.file")
		if err == nil {
			t.Error("expected error opening invalid file, got: nil")
		}
	})

	t.Run("TestJSONUnmarshalError", func(t *testing.T) {
		tmpfile, err := tmpFile("invalid JSON")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		_, err = tweet.NewConfig("", tmpfile.Name())
		if err == nil {
			t.Error("expected error unmarshalling JSON from config file, got: nil")
		}
	})

	t.Run("TestSuccess", func(t *testing.T) {
		tmpfile, err := tmpFile(`{
    			"ConsumerKey"      : "CONSUMER_KEY",
   			"ConsumerSecret"   : "CONSUMER_SECRET",
    			"AccessToken"      : "ACCESS_TOKEN",
    			"AccessTokenSecret": "ACCESS_TOKEN_SECRET"
		}`)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())

		_, err = tweet.NewConfig("", tmpfile.Name())
		if err != nil {
			t.Errorf("expected error: nil, got: %v", err)
		}
	})
}
