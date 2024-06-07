
.PHONY : tools mock docs

mod:
	go mod tidy

bidbot:
	mkdir -p .build
	go build -o .build/bidbot ./cmd/bidbot

sol:
	mkdir -p .build
	go build -o .build/sol ./cmd/sol

bundlebot:
	mkdir -p .build
	go build -o .build/bundlebot ./cmd/bundlebot

all:
	mkdir -p .build
	go build -o .build/bidbot ./cmd/bidbot
	go build -o .build/sol ./cmd/sol
