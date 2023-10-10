package main

import (
	. "fmt"
	"html/template"
	"io/ioutil"
	. "main/submodules"
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
	injectTemplate, err := ioutil.ReadFile("diverts/groxy/inject-template.js")
	if err != nil {
		Print(err.Error())
	}
	injectjs := []byte(clientTargetList + string(injectTemplate))
	err = ioutil.WriteFile("diverts/groxy/injects.js", injectjs, 0644)
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
		Handle(divertList[i], FileServer(Dir("diverts")))
	}
	HandleFunc("/", onRequest)
	HandleFunc("/search*", onRequest)
	ListenAndServe(":0", nil)
	Print("http server up!")
}

func onRequest(res ResponseWriter, request *Request) {
	hostTarget := hostTargetList[0]
	hostProxy := request.Host
	if request.URL.Query().Has("hostname") {
		hostTarget = request.URL.Query().Get("hostname")
	}
	request.Host = hostTarget
	pathname := "https://" + hostTarget +
		strings.Replace(
			strings.Replace(
				request.URL.RequestURI(),
				"?hostname="+request.Host, "", -1),
			"&hostname="+request.Host, "", -1)
	response := ProxyFetch(pathname, request)

	for i := 0; i < hostTargetList_length; i++ {
		if response.StatusCode < 400 {
			break
		}
		if request.Host == hostTargetList[i] {
			continue
		}
		request.Host = hostTargetList[i]
		pathname = "https://" + hostTargetList[i] +
			strings.Replace(
				strings.Replace(
					request.URL.RequestURI(),
					"?hostname="+request.Host, "", -1),
				"&hostname="+request.Host, "", -1)
		response = ProxyFetch(pathname, request)
	}

	bodyPromise := AsyncIoReadAll(response)
	ProxyResponseHeaders(&res, response, hostTarget, hostProxy)
	res.WriteHeader(response.StatusCode)
	bodyBytes, err := AwaitIoReadAll(bodyPromise)
	res.Write(bodyBytes)

	if err != nil {
		ErrorResponse(res, err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			ErrorResponse(res, "Unhandled Exception")
			Print("Unhandled Exception:\n", r)
		}
	}()
}
