package handler

import (
	. "handler/api/main"
	"net/http"
)


func OnRequest(responseWriter http.ResponseWriter, request *http.Request){
  HandleRequest(&responseWriter, request)
}

