package main

import (
	. "fmt"
	. "handler/api/main"
	"io/ioutil"
	. "net/http"
	"strings"
)


func writeClientList() {
	clientTargetList := strings.Join([]string(HostTargetList), "','")
	clientTargetList = " globalThis.hostTargetList = ['" + clientTargetList + "'];"
	injectTemplate, err := ioutil.ReadFile("api/groxy/inject-template.js")
	if err != nil {
		Print(err.Error())
	}
	injectjs := []byte(clientTargetList + string(injectTemplate))
	err = ioutil.WriteFile("api/groxy/injects-js.js", injectjs, 0644)
	if err != nil {
		Print(err.Error())
	}
	Print(clientTargetList)
}

func main() {
	go writeClientList()

	for i := 0; i < DivertList_length; i++ {
		Handle(DivertList[i], FileServer(Dir("api")))
	}
	HandleFunc("/", OnRequest)
	HandleFunc("/search*", OnRequest)
	ListenAndServe(":0", nil)
	Print("http server up!")
}

func OnRequest(responseWriter ResponseWriter, request *Request) {
	HandleRequest(&responseWriter, request)
}
