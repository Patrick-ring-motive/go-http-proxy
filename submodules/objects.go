package submodules

/**
Extensions for my convenience.
Mostly additions to mirror certain JavaScript syntax.
*/
import (
	. "fmt"
)

var Null Any = nil

func Interfacefy(i interface{}) interface{} {
	return i
}

type ExtendedMap struct {
	Map map[string]Any
}

func (f ExtendedMap) get(key Any) Any {
	return f.Map[ToString(key)]
}
func (f ExtendedMap) set(key Any, value Any) {
	f.Map[ToString(key)] = value
}

type Any interface{}

func Object() map[string]Any {
	obj := (make(map[string]Any))
	return obj
}

var Global = Object()

type AnyObject struct {
	value      Any
	properties map[string]*Any
}

func Let(Any) {}

func ToString(str Any) string {
	return Sprint(str)
}

func Stringify(str Any) string {
	switch s := str.(type) {
	case string:
		return s
	default:
		return ToString(s)
	}
}
