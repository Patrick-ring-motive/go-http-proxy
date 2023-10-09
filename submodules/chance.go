package submodules

import (
	"unsafe"
)

// This class is the laziest imaginable half attempt at generating "random numbers"
// Faster than other randoms but do not use if you need true randomness.

type chanceClass struct {
	Int  func() int
	Intn func(int) int
}

func randomizerInt(n unsafe.Pointer){
  
}

func goRandomizerInt(n unsafe.Pointer){
  go randomizerInt(n)
}

func intifyUnsafe(i Any) int {
	return *(*int)(unsafe.Pointer(&i))
}

func intify(ant Any) int {
	switch i := ant.(type) {
	case int:
		return i
	default:
		return intifyUnsafe(i)
	}
}


var chance = chanceClass{
	Int: func() int {
		n := 1
		xn := unsafe.Pointer(&n)
        goRandomizerInt(xn)
    yn := (int(uintptr(xn))%n)
    return yn
	},
	Intn: func(n int) int {
		xn := unsafe.Pointer(&n)
    goRandomizerInt(xn)
    yn := (int(uintptr(xn))%n)
    return yn
	},
}
