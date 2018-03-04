all: build-echo build-bot

build-echo:
	cd echo; env GOOS=linux go build -ldflags="-s -w" -o ../bin/echo
build-bot:
	cd bot; env GOOS=linux go build -ldflags="-s -w" -o ../bin/bot
