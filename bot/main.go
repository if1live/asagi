package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
	"gopkg.in/telegram-bot-api.v4"
)

var token string
var bot *tgbotapi.BotAPI

func main() {
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Panic("cannot find telegram token")
	}
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/bot", handler)

	// https://stackoverflow.com/questions/38393772/how-to-detect-if-im-running-in-aws-lambda-environment
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		http.Error(w, "can't unmarshal update json", http.StatusBadRequest)
		return
	}

	dispatch(update)

	w.Header().Set("Content-Type", "applicaiton/json")
	type Response struct {
		Ok bool `json:"ok"`
	}
	resp := Response{true}
	data, _ := json.Marshal(resp)
	w.Write(data)
}

func dispatch(update tgbotapi.Update) {
	cmdEcho(update)
}

func cmdEcho(update tgbotapi.Update) {
	fmt.Printf("update_id=%d message.text=%s", update.UpdateID, update.Message.Text)

	go func(update tgbotapi.Update) {
		echoMsg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		echoMsg.ReplyToMessageID = update.Message.MessageID
		bot.Send(echoMsg)
	}(update)
}
