package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/telegram-bot-api.v4"
)

func Test_Help_Register(t *testing.T) {
	c1 := newMock("a")
	c2 := newMock("b")
	c3 := newMock("c")

	update := tgbotapi.Update{}
	help := NewHelp(update)

	help.Register(c2)
	help.Register(c1)
	help.Register(c3)

	text := help.makeHelpText()
	assert.True(t, 'a' == text[0])
	assert.True(t, 'c' == text[len(text)-1])
}
