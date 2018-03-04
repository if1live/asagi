package commands

import (
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

type Command interface {
	Name() string
	Desc() string

	IsMatch(update tgbotapi.Update) bool
	Execute(update tgbotapi.Update, bot *tgbotapi.BotAPI)
}

// /foo bar spam => '/foo', 'bar spam'
// 명령에 따라서 인자 처리방식이 바뀔 가능성이 있다
// 인자를 텍스트로 남기고 파싱은 각각의 명령이 처리
func splitCommand(text string) (string, string) {
	start := strings.IndexAny(text, "/")
	if start < 0 {
		return "", ""
	}

	end := start + 1
	for end = start + 1; end < len(text); end++ {
		ch := text[end]
		if ch == ' ' || ch == '\t' {
			break
		}
	}
	name := text[start:end]

	arg := text[end:]
	arg = strings.Trim(arg, " \t")

	return name, arg
}
