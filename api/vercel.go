package main

import (
	. "fmt"
	"html/template"
  "io/ioutil"
	. "main/api/main"
	"net/http"
	"strings"
)


func OnRequest(responseWriter http.ResponseWriter, request *http.Request){
  HandleRequest(&responseWriter, request)
}

