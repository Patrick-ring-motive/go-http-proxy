package handler

import (
	"golang.org/x/exp/slices"
	. "handler/api/main"
	"net/http"
	"strings"
)

var repoList = DivertList

func OnRequest(responseWriter http.ResponseWriter, request *http.Request) {
	shortURI := strings.Split(strings.Split(request.URL.RequestURI(), "?")[0], "#")[0]
	if slices.Contains(repoList, shortURI) {
		RepoFetch(&responseWriter, request)
		return
	}
	HandleRequest(&responseWriter, request)

}
