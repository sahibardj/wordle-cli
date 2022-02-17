build:
	go build -o wordle src/main.go
install:
	make build && mv wordle /usr/local/bin/
