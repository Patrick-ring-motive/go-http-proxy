package submodules

import (
	"io"
	"net/http"
	"sync"
)

type PromiseXhttp struct {
	promiseChannel chan ([]byte)
	result         []byte
	resolved       bool
}

func IoReadAllPromise(response *http.Response) PromiseXhttp {
	promiseChannel := make(chan []byte)
	var promiseGroup sync.WaitGroup
	promiseGroup.Add(1)
	px := PromiseXhttp{promiseChannel: promiseChannel, result: []byte(""), resolved: false}
	go ioReadAllAsync(response, px)
	return px
}

func ioReadAllAsync(response *http.Response, px PromiseXhttp) {

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		defer console.log("error", err)
		px.result = []byte(err.Error())
		px.resolved = true
		px.promiseChannel <- px.result
	}
	defer func() {
		if r := recover(); r != nil {
			defer console.log("Unhandled Exception:\n", r)
			px.result = []byte("Unhandled Exception")
			px.resolved = true
			px.promiseChannel <- px.result
		}
	}()
	px.result = bodyBytes
	px.resolved = true
	px.promiseChannel <- px.result
}

///await
func AwaitXhttp(promis PromiseXhttp) []byte {
	if promis.resolved {
		return promis.result
	} else {
		promis.result = <-promis.promiseChannel
		promis.resolved = true
		return promis.result
	}
}
