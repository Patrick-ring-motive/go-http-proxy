package main

/**
Extensions for my convenience. 
Mostly additions to mirror certain JavaScript syntax.
*/
import (
	"fmt"
	"io"
	"net/http"
  "strings"
  "math/rand"
  "unsafe"

)

var null Any = nil;

func interfacefy(i interface{})(interface{}){
  return i;
}

type extendedMap struct{
  Map map[string]Any
}
func(f extendedMap)get(key Any)Any{
  return f.Map[toString(key)]
}
func(f extendedMap)set(key Any,value Any){
  f.Map[toString(key)]=value;
}


type console_log func(strs ...interface{});
type consoleObject struct { log console_log;
                            lag console_log};
var console = consoleObject{
  log:func(strs ...interface{}){fmt.Println(strs...)},
  lag:func(strs ...interface{}){go fmt.Println(strs...)},
}

func(c consoleObject )get()consoleObject{return c;}

type Any interface {};


func object()map[string]Any{
  obj := (make(map[string]Any));
  return obj;
}

var global = object();

type AnyObject struct{
  value Any;
  properties map[string]*Any;
}

func let(Any){}

func toString(str Any)string{
  return fmt.Sprint(str);
}

type String string;

var test = "asdf";

func stringify(str Any)string{
    switch s := str.(type) {
        case string:         
             return s;
        default:
             return toString(s);
    }
}

func(str String)replace(old Any, next Any)String{
 s := strings.Replace(stringify(str),stringify(old),stringify(next),1);
return String(s);
}

func(str String)replaceAll(old Any, next Any)String{
 s := strings.Replace(stringify(str),stringify(old),stringify(next),-1);
return String(s);
}

func(str String)val()string{
  return stringify(str);
}







var fetchClient = http.Client{};
var fetchClientInUseBy uintptr = 0;
var requestInit = create.xhttp_Request("GET","/",io.NopCloser(strings.NewReader("")));
func fetch(request Any)  (*http.Response){
 body := io.NopCloser(strings.NewReader(""));
 clientRequest := requestInit;
   switch requ := request.(type) {
        case string:         
             clientRequest = create.xhttp_Request("GET",requ,body);
        case *http.Request:
             clientRequest = requ;
        default:
             console.log(requ);
             clientRequest = create.xhttp_Request("GET","/",io.NopCloser(strings.NewReader("")));
    }
  fetchClientId := uintptr(unsafe.Pointer(clientRequest));
  fetchClientInUseBy = fetchClientId;
  client := fetchClient;
  defer releaseFetchClient(fetchClientId);
  defer client.CloseIdleConnections();
  response,err := client.Do(clientRequest);
  if (err != nil) {
    resetFetchClient(fetchClientId);
    response.StatusCode=500;
    response.Status="500 "+err.Error();
  }
    defer func() {
    if r := recover(); r != nil {
      resetFetchClient(fetchClientId);
      response.StatusCode=500;
      response.Status="500 Unhandled Exception: "+toString(r);
    }
  }();
  if(rand.Intn(100)==1){resetFetchClient(fetchClientId);}
  return response;
}

func releaseFetchClient(id uintptr){
  if(fetchClientInUseBy == id){
    fetchClientInUseBy = 0;
  }
}


func resetFetchClient(id uintptr){
  fetchClient.CloseIdleConnections();
  nextClient := http.Client{};
  if((fetchClientInUseBy == id)||(fetchClientInUseBy==0)){
    fetchClient = nextClient;
    fetchClientInUseBy=0;
  }
}