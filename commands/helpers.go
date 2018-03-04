package commands

import (
	"bytes"
	"fmt"
	"strings"
	"text/scanner"

	"gopkg.in/telegram-bot-api.v4"
)

func utf8substr(text string, pos int, length int) string {
	data := []byte(text)
	reader := bytes.NewReader(data)

	var s scanner.Scanner
	s.Init(reader)
	for i := 0; i < pos; i++ {
		s.Next()
	}

	strlist := []string{}
	for i := 0; i < length; i++ {
		r := s.Next()
		s := fmt.Sprintf("%c", r)
		strlist = append(strlist, s)
	}

	return strings.Join(strlist, "")
}

func getEntityText(text, typestr string, ent tgbotapi.MessageEntity) (string, bool) {
	if ent.Type == typestr {
		return utf8substr(text, ent.Offset, ent.Length), true
	}
	return "", false
}

func getEntityTexts(text, typestr string, entities []tgbotapi.MessageEntity) []string {
	retval := []string{}
	for _, ent := range entities {
		if s, ok := getEntityText(text, typestr, ent); ok {
			retval = append(retval, s)
		}
	}
	return retval
}
