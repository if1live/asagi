package commands

import (
	"fmt"
	"sort"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

type Help struct {
	commands []Command
}

func NewHelp(update tgbotapi.Update) *Help {
	return &Help{
		commands: []Command{},
	}
}

func (c *Help) Name() string {
	return "/help"
}
func (c *Help) Desc() string {
	return "show help"
}

func (c *Help) Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	helpText := c.makeHelpText()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
	msg.ParseMode = tgbotapi.ModeMarkdown
	bot.Send(msg)
}

func (c *Help) IsMatch(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}
	n, _ := splitCommand(update.Message.Text)
	return n == c.Name()
}

func (c *Help) Register(cmd Command) bool {
	c.commands = append(c.commands, cmd)
	sort.Slice(c.commands, func(i, j int) bool {
		return c.commands[i].Name() < c.commands[j].Name()
	})

	return true
}

func (c *Help) makeHelpText() string {
	lines := make([]string, len(c.commands))
	for i, cmd := range c.commands {
		lines[i] = fmt.Sprintf("%s - %s", cmd.Name(), cmd.Desc())
	}
	return strings.Join(lines, "\n")
}
