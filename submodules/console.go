package submodules

import (
	"fmt"
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