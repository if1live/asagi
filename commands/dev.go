package commands

import (
	"gopkg.in/telegram-bot-api.v4"
)

type DevKeyboard struct {
}

func NewDevKeyboard(update tgbotapi.Update) *DevKeyboard {
	return &DevKeyboard{}
}

func (c *DevKeyboard) Name() string {
	return "/devkeyboard"
}
func (c *DevKeyboard) Desc() string {
	return "keyboard test"
}

func (c *DevKeyboard) makeMarkupRoot() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("sub menu", "/sub"),
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonSwitch("switch", "sw"),
			},
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonURL("URL", "https://libsora.so"),
			},
		},
	}
}
func (c *DevKeyboard) makeMarkupSub() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("root", "/"),
			},
		},
	}
}

func (c *DevKeyboard) IsMatch(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	n, _ := splitCommand(update.Message.Text)
	return n == c.Name()
}

func (c *DevKeyboard) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	markup := c.makeMarkupRoot()
	chatID := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "sample keyboard")
	msg.ReplyMarkup = markup
	bot.Send(msg)
}

func (c *DevKeyboard) moveToRoot(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID
	markup := c.makeMarkupRoot()

	msg := tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, markup)
	bot.Send(msg)
}
func (c *DevKeyboard) moveToSub(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID
	markup := c.makeMarkupSub()
	msg := tgbotapi.NewEditMessageText(chatID, messageID, "sub menu")
	msg.ReplyMarkup = &markup
	bot.Send(msg)
}

func (c *DevKeyboard) Callback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	switch update.CallbackQuery.Data {
	case "/sub":
		c.moveToSub(update, bot)
	case "/":
		c.moveToRoot(update, bot)
	default:
		break
	}
}
