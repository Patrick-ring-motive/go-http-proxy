package main

import (
	. "fmt"
	"html/template"
  "io/ioutil"
	. "main/api"
	. "net/http"
	"strings"
)

var tmpl, err = template.ParseGlob("templates/*")
var hostTargetList = []string{
	"go.dev",
	"pkg.go.dev",
	"golang.org",
	"learn.go.dev",
	"play.golang.org",
	"proxy.golang.org",
	"sum.golang.org",
	"index.golang.org",
	"tour.golang.org",
	"blog.golang.org"}
var hostTargetList_length = len(hostTargetList)

func writeClientList() {
	clientTargetList := strings.Join([]string(hostTargetList), "','")
	clientTargetList = " globalThis.hostTargetList = ['" + clientTargetList + "'];"
	injectTemplate, err := ioutil.ReadFile("api/groxy/inject-template.js")
	if err != nil {
		Print(err.Error())
	}
	injectjs := []byte(clientTargetList + string(injectTemplate))
	err = ioutil.WriteFile("api/groxy/injects.js", injectjs, 0644)
	if err != nil {
		Print(err.Error())
	}
	Print(clientTargetList)
}

func main() {
	go writeClientList()
	divertList := []string{"/js/site.js",
		"/css/styles.css",
		"/static/frontend/frontend.js",
		"/static/frontend/frontend.min.css",
		"/tour/static/css/app.css",
		"/groxy/injects.js",
		"/groxy/injects.css",
    "/sw.js"}
	divertList_length := len(divertList)
	for i := 0; i < divertList_length; i++ {
		Handle(divertList[i], FileServer(Dir("api")))
	}
	HandleFunc("/", onRequest)
	HandleFunc("/search*", onRequest)
	ListenAndServe(":0", nil)
	Print("http server up!")
}

func onRequest(responseWriter ResponseWriter, request *Request){
  HandleRequest(&responseWriter, request)
}

