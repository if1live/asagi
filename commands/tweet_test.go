package commands

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/telegram-bot-api.v4"
)

func Test_isTwitterStatusURL_valid(t *testing.T) {
	urls := []string{
		"http://twitter.com/if1live/status/898755978153181185",
		"https://twitter.com/if1live/status/898755978153181185",
		"https://twitter.com/if1live/status/898755978153181185?s=09",
	}
	for _, u := range urls {
		assert.True(t, isTwitterStatusURL(u))
	}
}
func Test_isTwitterStatusURL_invalid(t *testing.T) {
	urls := []string{
		"http://google.com",
	}
	for _, u := range urls {
		assert.False(t, isTwitterStatusURL(u))
	}
}

func Test_TweetLink_IsMatch_true(t *testing.T) {
	fpList := []string{
		"../testdata/update_tweet_share.json",
		"../testdata/update_tweet_url.json",
	}
	for _, fp := range fpList {
		data, err := ioutil.ReadFile(fp)
		if err != nil {
			t.Fatalf("ReadFile: %v", err)
		}

		var update tgbotapi.Update
		err = json.Unmarshal(data, &update)
		if err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}

		cmd := NewTweetLink(update)
		assert.True(t, cmd.IsMatch(update))
	}
}
