package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"fmt"

	py "github.com/jackokring/cpy3"
	fe "github.com/jackokring/goali/filerr"
)

func Run(s string) {
	if py.PyRun_SimpleString(s) != 0 {
		Exit()
		fe.Fatal(fmt.Errorf("python exception: %s", s))
	}
}

func Init() {
	py.Py_Initialize()
	// To be usable as a part of snake
	// Must have objects to modify with extras
	Run("import snake")
}

func Exit() {
	py.Py_Finalize()
}
