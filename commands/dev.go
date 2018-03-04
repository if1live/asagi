package commands

import (
	"gopkg.in/telegram-bot-api.v4"
)

type DevKeyboard struct {
}

func NewDevKeyboard(update tgbotapi.Update) Command {
	return &DevKeyboard{}
}

func (c *DevKeyboard) Name() string {
	return "/devkeyboard"
}
func (c *DevKeyboard) Desc() string {
	return "keyboard test"
}

func (c *DevKeyboard) IsMatch(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	n, _ := splitCommand(update.Message.Text)
	return n == c.Name()
}

func (c *DevKeyboard) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	markup := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("data", "data"),
				tgbotapi.NewInlineKeyboardButtonSwitch("switch", "sw"),
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonURL("URL", "https://libsora.so"),
			},
		},
	}

	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "sample keyboard")
	msg.ReplyMarkup = markup
	bot.Send(msg)

}
