package handler

import (
	. "fmt"
	"html/template"
	. "main/api/_pkg"
	"net/http"
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


func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request){
  HandleRequest(&responseWriter, request)
}


func HandleRequest(responseWriter *http.ResponseWriter, request *http.Request) {
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
	ProxyResponseHeaders(responseWriter, response, hostTarget, hostProxy)
	(*responseWriter).WriteHeader(response.StatusCode)
	bodyBytes, err := AwaitIoReadAll(bodyPromise)
	(*responseWriter).Write(bodyBytes)

	if err != nil {
		ErrorResponse(*responseWriter, err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			ErrorResponse(*responseWriter, "Unhandled Exception")
			Print("Unhandled Exception:\n", r)
		}
	}()
}
