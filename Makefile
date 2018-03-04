build:
	cd echo; env GOOS=linux go build -ldflags="-s -w" -o ../bin/echo
