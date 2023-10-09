package main

import (
	"io"
	"net/http"
	"strings"
)

type createObject struct {
	xhttp_Request func(string, string, io.Reader) *http.Request
}

var create = createObject{
	xhttp_Request: func(method string, url string, body io.Reader) *http.Request {
		clientRequest, err := http.NewRequest(method, url, body)
		if err != nil {
			clientRequest.Header.Add("go-error", err.Error())
		}
		return clientRequest
	},
}

func ioReadAll(response *http.Response) []byte {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		defer console.log("error", err)
		return []byte(err.Error())
	}
	defer func() {
		if r := recover(); r != nil {
			defer console.log("Unhandled Exception:\n", r)
			bodyBytes = []byte("Unhandled Exception")
		}
	}()
	return bodyBytes
}

func erres(res http.ResponseWriter, errString string) {
	http.Error(res, errString, http.StatusInternalServerError)
	console.log(errString)
}

func transferRequestHeaders(req *http.Request) {
	reqHeaders := req.Header
	for key, val := range reqHeaders {
		for i := 0; i < len(val); i++ {
			req.Header.Add(key, val[i])
		}
	}
}

func proxyResponseHeaders(res *http.ResponseWriter, response *http.Response, hostTarget string, hostProxy string) {
	responseHeaders := response.Header
	for key, val := range responseHeaders {
		for i := 0; i < len(val); i++ {
			(*res).Header().Add(key,
				strings.Replace(val[i],
					hostTarget,
					hostProxy,
					-1))
		}
	}
	(*res).Header().Del("x-frame-options")
	(*res).Header().Del("content-security-policy")
	(*res).Header().Add("access-control-allow-origin", "*")
}

func proxyFetch(url string, req *http.Request) *http.Response {
	request := create.xhttp_Request(req.Method, url, req.Body)
	transferRequestHeaders(request)
	response := fetch(request)
	return response
}
