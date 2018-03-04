package commands

import (
	"gopkg.in/telegram-bot-api.v4"
)

type Echo struct {
}

func NewEcho(update tgbotapi.Update) Command {
	return &Echo{}
}
func (c *Echo) Name() string {
	return "/echo"
}
func (c *Echo) Desc() string {
	return "echo message"
}
func (c *Echo) IsMatch(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	n, _ := splitCommand(update.Message.Text)
	return n == c.Name()
}
func (c *Echo) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	_, arg := splitCommand(update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, arg)
	bot.Send(msg)
}
