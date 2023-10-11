package handler

import (
	. "fmt"
	"html/template"
  "io/ioutil"
	. "handler/api/main"
	"net/http"
	"strings"
)


func OnRequest(responseWriter http.ResponseWriter, request *http.Request){
  HandleRequest(&responseWriter, request)
}

