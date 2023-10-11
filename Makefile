.PHONY: run
run: main
	go run handler

main: *.go go.mod
	go build -o handler main.go
	chmod +x ./handler

.PHONY: all
all: main
