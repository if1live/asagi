package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
)

func mainWithNet() {
	http.HandleFunc("/asagi/echo", handler)

	// https://stackoverflow.com/questions/38393772/how-to-detect-if-im-running-in-aws-lambda-environment
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	} else {
		http.ListenAndServe(":8080", nil)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlerGet(w, r)
	case http.MethodPost:
		handlerPost(w, r)
	default:
		handlerPost(w, r)
	}
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	var resp EchoResponse
	msg := r.URL.Query()["q"]
	if msg != nil {
		resp = NewEchoResponse(msg[0], false)
	} else {
		resp = NewEchoResponse("", true)
	}

	writeReponse(w, resp)
}
func handlerPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var resp EchoResponse
	if len(body) > 0 {
		resp = NewEchoResponse(string(body), false)
	} else {
		resp = NewEchoResponse("", true)
	}
	writeReponse(w, resp)
}

func writeReponse(w http.ResponseWriter, r interface{}) {
	w.Header().Set("Content-Type", "applicaiton/json")

	data, _ := json.Marshal(r)
	w.Write(data)
}
