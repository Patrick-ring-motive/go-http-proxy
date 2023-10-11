.PHONY: run
run: main
	go run handler

main: *.go go.mod
	bash ./build.sh &
	go build -o handler main.go
	chmod +x ./handler

.PHONY: all
all: main
