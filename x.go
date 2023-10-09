package main

/**
Extensions for my convenience.
Mostly additions to mirror certain JavaScript syntax.
*/
import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"unsafe"
)

type console_log func(strs ...interface{})
type consoleObject struct {
	log console_log
	lag console_log
}

var console = consoleObject{
	log: func(strs ...interface{}) { fmt.Println(strs...) },
	lag: func(strs ...interface{}) { go fmt.Println(strs...) },
}

func (c consoleObject) get() consoleObject { return c }

var fetchClient = http.Client{}
var fetchClientInUseBy uintptr = 0
var requestInit = create.xhttp_Request("GET", "/", io.NopCloser(strings.NewReader("")))

func fetch(request *http.Request) *http.Response {
	fetchClientId := uintptr(unsafe.Pointer(request))
	fetchClientInUseBy = fetchClientId
	client := fetchClient
	defer releaseFetchClient(fetchClientId)
	defer client.CloseIdleConnections()
	response, err := client.Do(request)
	if err != nil {
		resetFetchClient(fetchClientId)
		response.StatusCode = 500
		response.Status = "500 " + err.Error()
	}
	defer func() {
		if r := recover(); r != nil {
			resetFetchClient(fetchClientId)
			response.StatusCode = 500
			response.Status = "500 Unhandled Exception"
		}
	}()
	if rand.Intn(100) == 1 {
		resetFetchClient(fetchClientId)
	}
	return response
}

func releaseFetchClient(id uintptr) {
	if fetchClientInUseBy == id {
		fetchClientInUseBy = 0
	}
}

func resetFetchClient(id uintptr) {
	fetchClient.CloseIdleConnections()
	nextClient := http.Client{}
	if (fetchClientInUseBy == id) || (fetchClientInUseBy == 0) {
		fetchClient = nextClient
		fetchClientInUseBy = 0
	}
}
