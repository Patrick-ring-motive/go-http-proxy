package submodules

import (
	. "fmt"
	"io"
	. "net/http"
	"strings"
	"sync"
	. "unsafe"
)

var s = Let(sync.Mutex{})


func HttpServerlessRequest(responseWriter ResponseWriter, request *Request){
  ReflectRequest(&responseWriter, request)
}

func CreateRequest(method string, url string, body io.Reader) *Request {
	request, err := NewRequest(method, url, body)
	if err != nil {
		request.Header.Add("go-error", err.Error())
	}
	return request
}

func ErrorResponse(res ResponseWriter, errString string) {
	Error(res, errString, StatusInternalServerError)
	Print(errString)
}

func ReflectRequest(responseWriter *ResponseWriter, request *Request) {
	requestHeaders := request.Header
	for key, val := range requestHeaders {
		for i := 0; i < len(val); i++ {
			(*res).Header().Add(key, val[i])
		}
	}
  	(*responseWriter).WriteHeader(200)
  	(*responseWriter).Write([]bytes(Sprint(*request)))
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

func TransferRequestHeaders(req *Request) {
	reqHeaders := req.Header
	for key, val := range reqHeaders {
		for i := 0; i < len(val); i++ {
			req.Header.Add(key, val[i])
		}
	}
}

func ProxyResponseHeaders(res *ResponseWriter, response *Response, hostTarget string, hostProxy string) {
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

var fetchClient = Client{}
var fetchClientInUseBy uintptr = 0
var requestInit = CreateRequest("GET", "/", io.NopCloser(strings.NewReader("")))

func Fetch(request *Request) *Response {
	fetchClientId := uintptr(Pointer(request))
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
			response.Status = "500 Unhandled Exception: " + Sprint(r)
		}
	}()
	if Intn(100) == 0 {
		resetFetchClient(fetchClientId)
	}
	return response
}

func FetchURL(url string) *Response {
	body := io.NopCloser(strings.NewReader(""))
	request := CreateRequest("GET", url, body)
	return Fetch(request)
}

func releaseFetchClient(id uintptr) {
	if fetchClientInUseBy == id {
		fetchClientInUseBy = 0
	}
}

func resetFetchClient(id uintptr) {
	fetchClient.CloseIdleConnections()
	nextClient := Client{}
	if (fetchClientInUseBy == id) || (fetchClientInUseBy == 0) {
		fetchClient = nextClient
		fetchClientInUseBy = 0
	}
}

func ProxyFetch(url string, req *Request) *Response {
	request := CreateRequest(req.Method, url, req.Body)
	TransferRequestHeaders(request)
	response := Fetch(request)
	return response
}

func IoReadAll(response *Response) []byte {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		defer Print("error", err)
		return []byte(err.Error())
	}
	defer func() {
		if r := recover(); r != nil {
			defer Print("Unhandled Exception:\n", r)
			bodyBytes = []byte("Unhandled Exception")
		}
	}()
	return bodyBytes
}

type ThreadIoReadAll struct {
	ThreadChannel chan ([]byte)
	Lock          *PromiseIoReadAll
}

var Unlocked = &PromiseIoReadAll{PromiseChannel:nil, Error: nil, Result: []byte(""), Resolved: false, Rejected: false}

func NewThreadIoReadAll()ThreadIoReadAll{
  threadChannel := make(chan ([]byte))
  thread := ThreadIoReadAll{ThreadChannel: threadChannel, Lock: Unlocked}
  return thread
}

var ThreadPoolIoReadAll = []ThreadIoReadAll{}

func initializeThreadPool(numThreads int) []ThreadIoReadAll {
	ThreadPoolIoReadAll = make([]ThreadIoReadAll, numThreads)
  for i := range ThreadPoolIoReadAll {
        ThreadPoolIoReadAll[i] = NewThreadIoReadAll()
  }
	return ThreadPoolIoReadAll
}

func goInitializeThreadPool(numThreads int) []ThreadIoReadAll {
	go initializeThreadPool(numThreads)
	return ThreadPoolIoReadAll
}

type PromiseIoReadAll struct {
	PromiseChannel chan ([]byte)
	Error          error
	Result         []byte
	Resolved       bool
	Rejected       bool
}

func AsyncIoReadAll(response *Response) PromiseIoReadAll {
	promiseChannel := make(chan []byte)
	promise := PromiseIoReadAll{PromiseChannel: promiseChannel, Error: nil, Result: []byte(""), Resolved: false, Rejected: false}
	go GoIoReadAllAsync(response, promise)
	return promise
}

func GoIoReadAllAsync(response *Response, promise PromiseIoReadAll) PromiseIoReadAll {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		defer Print("error", err)
		promise.Result = []byte(err.Error())
		promise.Rejected = true
		promise.Error = err
		promise.PromiseChannel <- promise.Result
		return promise
	}
	promise.Result = bodyBytes
	promise.Resolved = true
	promise.PromiseChannel <- promise.Result
	return promise
}

func AwaitIoReadAll(promise PromiseIoReadAll) ([]byte, error) {
	if promise.Resolved || promise.Rejected {
		return promise.Result, promise.Error
	} else {
		promise.Result = <-promise.PromiseChannel
		return promise.Result, promise.Error
	}
}
