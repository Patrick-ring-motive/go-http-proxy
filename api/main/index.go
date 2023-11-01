package indexgo

import (
	. "fmt"
	. "handler/api/main/src"
	"html/template"
	"net/http"
	"strings"
)

var tmpl, err = template.ParseGlob("templates/*")
var HostTargetList = []string{
	"go.dev",
	"pkg.go.dev",
	"golang.org",
	"learn.go.dev",
	"play.golang.org",
	"proxy.golang.org",
	"sum.golang.org",
	"index.golang.org",
	"tour.golang.org",
	"play.golang.org",
	"blog.golang.org"}

var HostTargetList_length = len(HostTargetList)

var DivertList = []string{"/js/site.js",
	"/css/styles.css",
	"/images/favicon-gopher.png",
	"/images/favicon-gopher.svg",
	"/static/shared/icon/favicon.ico",
	"/static/frontend/frontend.js",
	"/static/frontend/frontend.min.css",
	"/tour/static/css/app.css",
	"/groxy/injects-js.js",
	"/groxy/injects-css.css",
	"/sw.js"}

var DivertList_length = len(DivertList)

func OnServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	HandleRequest(&responseWriter, request)
}

func HandleRequest(responseWriter *http.ResponseWriter, request *http.Request) {
	hostTarget := HostTargetList[0]
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

	for i := 0; i < HostTargetList_length; i++ {
		if response.StatusCode < 400 {
			break
		}
		if request.Host == HostTargetList[i] {
			continue
		}
		request.Host = HostTargetList[i]
		pathname = "https://" + HostTargetList[i] +
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


	defer (*responseWriter).Write(bodyBytes)

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

var repoRoot = "https://raw.githubusercontent.com/Patrick-ring-motive/go-http-proxy/main/api"

func RepoFetch(responseWriter *http.ResponseWriter, request *http.Request) {
	uri := request.URL.RequestURI()
	url := repoRoot + uri
	response := FetchURL(url)

	bodyPromise := AsyncIoReadAll(response)
	contentType := "text/html"
	if strings.Contains(uri, ".css") {
		contentType = "text/css"
	}
	if strings.Contains(uri, ".js") {
		contentType = "text/javascript"
	}
	if strings.Contains(uri, ".css") {
		contentType = "text/css"
	}
	if strings.Contains(uri, ".ico") {
		contentType = "image/x-icon"
	}
	if strings.Contains(uri, ".svg") {
		contentType = "image/svg+xml"
	}
	if strings.Contains(uri, ".png") {
		contentType = "image/png"
	}

  
	(*responseWriter).WriteHeader(response.StatusCode)
	(*responseWriter).Header().Del("x-frame-options")
	(*responseWriter).Header().Del("content-security-policy")
	(*responseWriter).Header().Set("access-control-allow-origin", "*")
	(*responseWriter).Header().Set("content-type", contentType)
	bodyBytes, err := AwaitIoReadAll(bodyPromise)

	defer (*responseWriter).Write(bodyBytes)

	if err != nil {
		ErrorResponse(*responseWriter, err.Error())
		return
	}

}

