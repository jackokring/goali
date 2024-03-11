package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"fmt"

	py "github.com/jackokring/cpy3"
	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

func Run(s string) {
	if py.PyRun_SimpleString(s) != 0 {
		Exit()
		fe.Fatal(fmt.Errorf("python exception: %s", s))
	}
}

func RunFile(f clit.InputFile) {
	if f.Expand {
		fe.Fatal(fmt.Errorf("flag -e not allowed: %s", f.InputFile))
	}
	code, err := py.PyRun_AnyFile(f.InputFile)
	fe.Fatal(err)
	if code != 0 {
		Exit()
		fe.Fatal(fmt.Errorf("python exception in file: %s", f.InputFile))
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
