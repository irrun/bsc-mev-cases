
.PHONY : tools mock docs

mod:
	go mod tidy

bidbot:
	go build -o bidbot ./cmd/bidbot

sol:
	go build -o sol ./cmd/sol

all:
	go build -o bidbot ./cmd/bidbot
	go build -o sol ./cmd/sol
