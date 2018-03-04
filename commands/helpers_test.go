package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_utf8substr(t *testing.T) {
	text := "@if1live 님의 트윗을 확인해 보세요. https://twitter.com/if1live/status/898755978153181185?s=09한글"
	expected := "https://twitter.com/if1live/status/898755978153181185?s=09"
	offset := 25
	length := 58

	actual := utf8substr(text, offset, length)
	assert.Equal(t, expected, actual)
}
