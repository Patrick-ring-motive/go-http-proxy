package submodules

import (
	. "fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	. "unsafe"
  "math/rand"
)

var void = Void(sync.Mutex{})

func HttpServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	ReflectRequest(&responseWriter, request)
}

func CreateRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		request.Header.Add("go-error", err.Error())
	}
	return request
}

func ErrorResponse(res http.ResponseWriter, errString string) {
	http.Error(res, errString, http.StatusInternalServerError)
	Print(errString)
}

func ReflectRequest(responseWriter *http.ResponseWriter, request *http.Request) {
	requestHeaders := request.Header
	for key, val := range requestHeaders {
		for i := 0; i < len(val); i++ {
			(*responseWriter).Header().Add(key, val[i])
		}
	}
	(*responseWriter).WriteHeader(200)
	(*responseWriter).Write([]byte(Sprint(*request)))

}

func TransferRequestHeaders(newRequest *http.Request, oldRequest *http.Request) {
	reqHeaders := oldRequest.Header
	hostOld := oldRequest.Host
	hostNew := newRequest.Host
	for key, val := range reqHeaders {
		for i := 0; i < len(val); i++ {
			newRequest.Header.Add(key, strings.Replace(val[i],
				hostOld,
				hostNew,
				-1))
		}
	}
}

func ProxyResponseHeaders(res *http.ResponseWriter, response *http.Response, hostOld string, hostNew string) {
  resHeaders := (*res).Header()
	responseHeaders := response.Header
	for key, val := range responseHeaders {
		for i := 0; i < len(val); i++ {
      resHeaders.Add(key,
				strings.Replace(val[i],
					hostOld,
					hostNew,
					-1))
		}
	}
  resHeaderMap := HeaderToMap((*res).Header())
  delete(resHeaderMap,"X-Frame-Options")
  delete(resHeaderMap,"Content-Security-Policy")
  resHeaderMap["Access-Control-Allow-Origin"]=[]string{"*"}

	
}

var fetchClient = http.Client{}
var fetchClientInUseBy uintptr = 0
var requestInit = CreateRequest("GET", "/", io.NopCloser(strings.NewReader("")))

func Fetch(request *http.Request) *http.Response {
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
	if rand.Intn(100) == 0 {
		resetFetchClient(fetchClientId)
	}
	return response
}

func FetchURL(url string) *http.Response {
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
	nextClient := http.Client{}
	if (fetchClientInUseBy == id) || (fetchClientInUseBy == 0) {
		fetchClient = nextClient
		fetchClientInUseBy = 0
	}
}

func ProxyFetch(url string, req *http.Request) *http.Response {
	request := CreateRequest(req.Method, url, req.Body)
	TransferRequestHeaders(request, req)
	response := Fetch(request)
	return response
}

func GetResponseBody(response *http.Response) []byte {
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		defer Print("error", err)
    response.StatusCode = 500
    response.Status = err.Error()
		return []byte(response.Status)
	}
	defer func() {
		if r := recover(); r != nil {
			defer Print("Unhandled Exception:\n", r)
      response.StatusCode = 500
      response.Status = Sprint(r)
			bodyBytes = []byte(response.Status)
		}
	}()
	return bodyBytes
}

/******************IO Read All Promise Structures*************************/
func IoReadAll(response *http.Response) []byte {
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

type PromiseIoReadAll struct {
	PromiseChannel chan ([]byte)
	Error          error
	Result         []byte
	Resolved       bool
	Rejected       bool
}

func AsyncIoReadAll(response *http.Response) PromiseIoReadAll {
	promiseChannel := make(chan []byte)
	promise := PromiseIoReadAll{PromiseChannel: promiseChannel, Error: nil, Result: []byte(""), Resolved: false, Rejected: false}
	go GoIoReadAll(response, promise)
	return promise
}

func GoIoReadAll(response *http.Response, promise PromiseIoReadAll) PromiseIoReadAll {
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
    close(promise.PromiseChannel)
		return promise.Result, promise.Error
	}
}


func HeaderToMap(h http.Header) map[string][]string {
  return *(* map[string][]string)(Pointer(&h))
}
/***************IO Read All Thread Structure************************/

/*
type ThreadIoReadAll struct {
	ThreadChannel chan (*http.Response)
	Lock          PromiseIoReadAll
}

var Unlocked = PromiseIoReadAll{PromiseChannel: nil, Error: nil, Result: []byte(""), Resolved: false, Rejected: false}

func UnlockThread(thread ThreadIoReadAll) {
	thread.Lock = Unlocked
}

func NewThreadIoReadAll() ThreadIoReadAll {
	threadChannel := make(chan (*http.Response))
	thread := ThreadIoReadAll{ThreadChannel: threadChannel, Lock: Unlocked}
	go StartThreadIoReadAll(thread)
	return thread
}

var ThreadPoolIoReadAll = []ThreadIoReadAll{}
var GoThreadPoolIoReadAll = goInitializeThreadPool(10)

func initializeThreadPool(numThreads int) []ThreadIoReadAll {
  threadPoolIoReadAll := make([]ThreadIoReadAll, numThreads)
	
	for i := range ThreadPoolIoReadAll {
		threadPoolIoReadAll[i] = NewThreadIoReadAll()
	}
  ThreadPoolIoReadAll = threadPoolIoReadAll
	return ThreadPoolIoReadAll
}

func goInitializeThreadPool(numThreads int) []ThreadIoReadAll {
	go initializeThreadPool(numThreads)
	return ThreadPoolIoReadAll
}

func AquireThread(promise PromiseIoReadAll) ThreadIoReadAll {
  threadPoolIoReadAll := (ThreadPoolIoReadAll)
	for i := range threadPoolIoReadAll {
    thread := threadPoolIoReadAll[i]
		if thread.Lock.PromiseChannel == Unlocked.PromiseChannel {
      thread.Lock = promise
      if thread.Lock.PromiseChannel == promise.PromiseChannel{
			 return thread
      }
		}
	}
  threadChannel := make(chan (*http.Response))
  newThread := ThreadIoReadAll{ThreadChannel: threadChannel, Lock: promise}
  threadPoolIoReadAll=append(threadPoolIoReadAll,newThread)
  ThreadPoolIoReadAll= threadPoolIoReadAll
	return newThread
}

func StartThreadIoReadAll(thread ThreadIoReadAll) {
  for true{
	response := <-thread.ThreadChannel
	GoThreadIoReadAll(response, thread.Lock)
    }
}

func AsyncThreadIoReadAll(response *http.Response) ThreadIoReadAll {
  promiseChannel := make(chan []byte)
	promise := PromiseIoReadAll{PromiseChannel: promiseChannel, Error: nil, Result: []byte(""), Resolved: false, Rejected: false}
	thread := AquireThread(promise)
	thread.ThreadChannel <- response
	return thread
}

func GoThreadIoReadAll(response *http.Response, promise PromiseIoReadAll) PromiseIoReadAll {
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

func AwaitThreadIoReadAll(thread ThreadIoReadAll) ([]byte, error) {
	promise := thread.Lock
	if promise.Resolved || promise.Rejected {
		return promise.Result, promise.Error
	} else {
		defer UnlockThread(thread)
		promise.Result = <-promise.PromiseChannel
		return promise.Result, promise.Error
	}
}
*/