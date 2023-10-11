package submodules

import (
	"math/rand"
	"net/http"
	"time"
)

// This class is the laziest imaginable half attempt at generating "random numbers"
// Faster than other randoms but do not use if you need true randomness.

func ChanceServerlessRequest(responseWriter http.ResponseWriter, request *http.Request) {
	ReflectRequest(&responseWriter, request)
}

var chanceNum = rand.Int()

//var chancego = goChancer()
var sleeper time.Duration = 1

func goChancer() int {
	go chancer()
	return 0
}

func chancer() {
	for true {
		chanceNum = rand.Int()
		time.Sleep(sleeper * time.Nanosecond)
		sleeper = (sleeper * 2) % 1000000
	}
}

func Int() int {
	sleeper = 1
	return chanceNum

}

func Intn(n int) int {
	sleeper = 1
	return chanceNum % n

}
