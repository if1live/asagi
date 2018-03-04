package commands

import (
	"fmt"
	"net/url"
	"regexp"

	"gopkg.in/telegram-bot-api.v4"
)

type TweetLink struct {
}

func NewTweetLink(update tgbotapi.Update) Command {
	return &TweetLink{}
}
func (c *TweetLink) Name() string {
	return "%tweetlink"
}
func (c *TweetLink) Desc() string {
	return "save tweet"
}
func (c *TweetLink) IsMatch(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}

	text := update.Message.Text
	uris := getEntityTexts(text, "url", *update.Message.Entities)
	for _, uri := range uris {
		if isTwitterStatusURL(uri) {
			return true
		}
	}
	return false
}

func (c *TweetLink) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	text := update.Message.Text
	uris := getEntityTexts(text, "url", *update.Message.Entities)
	for _, uri := range uris {
		if !isTwitterStatusURL(uri) {
			continue
		}

		m := fmt.Sprintf("LINK: %s", uri)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, m)
		bot.Send(msg)
	}
}

func isTwitterStatusURL(uri string) bool {
	parsed, err := url.Parse(uri)
	if err != nil {
		return false
	}

	// twitter status
	if parsed.Host == "twitter.com" {
		reStatus := regexp.MustCompile(`/.+/status/\d+`)
		if reStatus.MatchString(uri) {
			return true
		}
	}
	return false
}
