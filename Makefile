GITHASH=$(shell git rev-list -1 HEAD)
DATE=$(shell date +'%FT%TZ%:z')

FLAGS="-X main.GitCommitHash=$(GITHASH) -X main.BuildDate=$(DATE)"

timestamp:
	go build -ldflags $(FLAGS)

.PHONY: install
install: timestamp
	go install -ldflags $(FLAGS)
