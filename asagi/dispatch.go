package asagi

import (
	"github.com/if1live/asagi/commands"
	"gopkg.in/telegram-bot-api.v4"
)

func Dispatch(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message != nil {
		dispatchMessage(update, bot)
	} else if update.CallbackQuery != nil {
		dispatchCallbackQuery(update, bot)
	}
}

func dispatchMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	cmds := []commands.Command{
		commands.NewEcho(update),
		commands.NewTweetByShare(update),
		// dev
		commands.NewDevKeyboard(update),
	}
	help := commands.NewHelp(update)
	for _, cmd := range cmds {
		help.Register(cmd)
	}

	allcmds := append(cmds, help)
	for _, cmd := range allcmds {
		if cmd.IsMatch(update) {
			cmd.Execute(update, bot)
			break
		}
	}
}

func dispatchCallbackQuery(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// TODO 키보드 이외의 것도 지원하도록 구현하기
	cmd := commands.NewDevKeyboard(update)
	cmd.Callback(update, bot)
}
