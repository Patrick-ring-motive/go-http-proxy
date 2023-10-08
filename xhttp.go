package main

import (
	"net/http"
  "unsafe"
  "io"
  "strings"
)

type createObject struct { xhttp_Request func(string,string,io.Reader)(*http.Request)};

var create = createObject{
  xhttp_Request:func (method string, url string, body io.Reader) (*http.Request){
    clientRequest,err := http.NewRequest(method,url,body);
    if(err!=nil){
      clientRequest.Header.Add("go-error",err.Error());
    }
    return clientRequest;
  },
}

func xhttp_ResponsefyUnsafe(x *Any) *http.Response {
	return (*http.Response)(unsafe.Pointer(x))
}

func xhttp_Responsefy(x Any)(*http.Response){
      switch xhr := x.(type) {
        case *http.Response:         
             return xhr;
        case http.Response:         
             return &xhr;
        default:
             return xhttp_ResponsefyUnsafe(&x);
    }
}

func byteSlicefyUnsafe(b Any) []byte {
	return *(*[]byte)(unsafe.Pointer(&b))
}

func byteSlicefy(bite Any)[]byte{
      switch b := bite.(type) {
        case []byte:         
             return b;
        default:
             return byteSlicefyUnsafe(b);
    }
}

func ioReadAll(response *http.Response)[]byte{
  defer response.Body.Close();
  bodyBytes,err := io.ReadAll(response.Body);
  if (err != nil) {
    defer console.log("error", err);
		return []byte(err.Error());	
	}
  defer func() {
    if r := recover(); r != nil {        
        defer console.log("Unhandled Exception:\n", r);
        bodyBytes = []byte(toString(r));	
    }
  }();
  return bodyBytes;
}

func ioReadAllAsyncWrapper(xhr Any)(Any){
  return ioReadAll(xhttp_Responsefy(xhr));
}

var ioReadAllRefister = asyncRegister(ioReadAll,ioReadAllAsyncWrapper);



func erres(res http.ResponseWriter,errString string){
  	http.Error(res, errString, http.StatusInternalServerError);
		console.log(errString);
}

func transferRequestHeaders(req *http.Request){
        reqHeaders := req.Header;
          for key,val := range reqHeaders {
              for i := 0;i < len(val) ;i++ {
                req.Header.Add(key,val[i]);  
              }  
          }  
}

func proxyResponseHeaders(res *http.ResponseWriter,response *http.Response,hostTarget string, hostProxy string){
    responseHeaders := response.Header;
  for key,val := range responseHeaders {
      for i := 0;i < len(val) ;i++ {
        (*res).Header().Add(key,
             strings.Replace(val[i],
                             hostTarget,
                             hostProxy,
                             -1));      
      }  
  }
  (*res).Header().Del("x-frame-options");
  (*res).Header().Del("content-security-policy");
  (*res).Header().Add("access-control-allow-origin","*");
}

func proxyFetch(url string, req *http.Request)  (*http.Response){
  request := create.xhttp_Request(req.Method,url,req.Body);
  transferRequestHeaders(request);
  response := fetch(request);
  return response;
}