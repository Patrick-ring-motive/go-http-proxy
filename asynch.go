package main

import (
	"sync"
)

type Promise struct{
  promiseChannel chan (Any);
  result Any;
  resolved bool;
}


///async
func promise(function func(Any) Any, payload Any) Promise {
	promiseChannel := make(chan Any)
	var promiseGroup sync.WaitGroup;
	promiseGroup.Add(1);
	go asyncRun(function, payload, promiseChannel, &promiseGroup);
	return Promise{promiseChannel:promiseChannel,result:nil,resolved:false};
}

func asyncRun(function func(Any) Any, payload Any, promiseChannel chan Any, promiseGroup *sync.WaitGroup) {
	defer promiseGroup.Done()
	output := function(payload)
	promiseChannel <- output
}

///await
func await(promis Promise) Any {
	promis.result = <- promis.promiseChannel;
  promis.resolved = true;
	return promis.result;
}

var asyncMap = extendedMap{Map: object()}

func asyncRegister(fun Any, afun func(Any) Any) Any {
	asyncMap.set(fun, afun)
	return afun
}
func async(fun Any) func(Any) Any {
	return asyncify(asyncMap.get(fun))
}

func asunc(a Any) Any { return a }

func asyncify(afun Any) func(Any) Any {
	switch af := afun.(type) {
	case func(Any) Any:
		return af
	default:
		return asunc
	}
}

////work

func exampleFunc(param Any) Any {
	return param
}

func exampleFuncAsyncWrapper(payload Any) Any {

	output := exampleFunc(payload)

	return output
}

var exampleFuncRegister = asyncRegister(exampleFunc, exampleFuncAsyncWrapper)
