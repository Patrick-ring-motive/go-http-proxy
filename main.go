package main

import (
	"html/template"
	"io/ioutil"
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

func writeClientList() {
	clientTargetList := strings.Join([]string(hostTargetList), "','")
	clientTargetList = " globalThis.hostTargetList = ['" + clientTargetList + "'];"
	injectTemplate, err := ioutil.ReadFile("diverts/groxy/inject-template.js")
	if err != nil {
		console.log(err.Error())
	}
	injectjs := []byte(clientTargetList + string(injectTemplate))
	err = ioutil.WriteFile("diverts/groxy/injects.js", injectjs, 0644)
	if err != nil {
		console.log(err.Error())
	}
	console.log(clientTargetList)
}

func main() {
	go writeClientList()
	divertList := []string{"/js/site.js",
		"/css/styles.css",
		"/static/frontend/frontend.js",
		"/static/frontend/frontend.min.css",
		"/tour/static/css/app.css",
		"/groxy/injects.js",
		"/groxy/injects.css"}
	divertList_length := len(divertList)
	for i := 0; i < divertList_length; i++ {
		http.Handle(divertList[i], http.FileServer(http.Dir("diverts")))
	}
	http.HandleFunc("/", onRequest)
	http.HandleFunc("/search*", onRequest)
	http.ListenAndServe(":0", nil)
	console.log("http server up!")
}

func onRequest(res http.ResponseWriter, req *http.Request) {
	hostTarget := hostTargetList[0]
	hostProxy := req.Host
	if req.URL.Query().Has("hostname") {
		hostTarget = req.URL.Query().Get("hostname")
	}
	req.Host = hostTarget
	pathname := "https://" + hostTarget +
		strings.Replace(
			strings.Replace(
				req.URL.RequestURI(),
				"?hostname="+req.Host, "", -1),
			"&hostname="+req.Host, "", -1)
	response := proxyFetch(pathname, req)

	for i := 0; i < hostTargetList_length; i++ {
		if response.StatusCode < 400 {
			break
		}
		if req.Host == hostTargetList[i] {
			continue
		}
		req.Host = hostTargetList[i]
		pathname = "https://" + hostTargetList[i] +
			strings.Replace(
				strings.Replace(
					req.URL.RequestURI(),
					"?hostname="+req.Host, "", -1),
				"&hostname="+req.Host, "", -1)
		response = proxyFetch(pathname, req)
	}

	bodyPromise := promise(async(ioReadAll), response)

	proxyResponseHeaders(&res, response, hostTarget, hostProxy)

	res.WriteHeader(response.StatusCode)
	body := await(bodyPromise)
	bodyBytes := byteSlicefy(body)
	res.Write(bodyBytes)

	if err != nil {
		erres(res, err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			erres(res, "Unhandled Exception: "+toString(r))
			console.log("Unhandled Exception:\n", r)
		}
	}()
}
