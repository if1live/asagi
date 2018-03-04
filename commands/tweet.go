package commands

import (
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

// 트위터 공앱에서 트윗을 공유한경우
// 앱에서 메세지가 그대로 보낸다
// 규격이 어느정도 정해진거같으니 이를 이용해서 파싱하면 될듯
// @if1live 님의 트윗을 확인해 보세요.
// https://twitter.com/if1live/status/898755978153181185?s=09
type TweetByShare struct {
}

func NewTweetByShare(update tgbotapi.Update) Command {
	return &TweetByShare{}
}
func (c *TweetByShare) Name() string {
	return "@"
}
func (c *TweetByShare) Desc() string {
	return "save shared tweet"
}
func (c *TweetByShare) IsMatch(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	// 문장 전체를 대조하긴 귀찮아
	// TODO URL 기반으로 분기하는걸 나중에 구현하면 여러 목적으로 쓸수있을듯
	return update.Message.Text[0] == '@'
}

func (c *TweetByShare) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// TODO 어디에 저장하지?
	t := fmt.Sprintf("TODO: %s is not impleted command", c.Name())
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, t)
	bot.Send(msg)
}
