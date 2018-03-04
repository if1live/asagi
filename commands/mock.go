package commands

import "gopkg.in/telegram-bot-api.v4"

type mock struct {
	name string
}

func newMock(name string) Command {
	return &mock{name}
}
func (c *mock) Name() string {
	return c.name
}
func (c *mock) Desc() string {
	return c.name
}
func (c *mock) IsMatch(update tgbotapi.Update) bool {
	return true
}
func (c *mock) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
}