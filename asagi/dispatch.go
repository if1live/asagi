package asagi

import (
	"github.com/if1live/asagi/commands"
	"gopkg.in/telegram-bot-api.v4"
)

func Dispatch(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
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
