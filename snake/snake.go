package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"fmt"
	"runtime"

	py "github.com/jackokring/cpy3"
	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

type Command struct {
	clit.PyFile
	clit.StreamFilter // embedded type .. => .
}

func (c *Command) Help() string {
	return `An embedded python script interpreter.`
}

func (c *Command) Run(p *clit.Globals) error {
	// unicorn command hook
	fe.SetGlobals(p)
	fmt.Println(c.InputFile)
	fmt.Println(c.OutputFile)
	return nil
}

func Run(s string) {
	if py.PyRun_SimpleString(s) != 0 {
		Exit()
		fe.Fatal(fmt.Errorf("python exception: %s", s))
	}
}

func RunFile(f clit.InputFile) {
	//if f.Expand {
	//fe.Fatal(fmt.Errorf("flag -e not allowed: %s", f.InputFile))
	// ignore flag as may be present for data files
	//}
	code, err := py.PyRun_AnyFile(f.InputFile)
	fe.Fatal(err)
	if code != 0 {
		Exit()
		fe.Fatal(fmt.Errorf("python exception in file: %s", f.InputFile))
	}
}

// The master imported object
var snake *py.PyObject

// The thread state
var state *py.PyThreadState

func Init() {
	py.Py_Initialize()
	// To be usable as a part of snake
	// Must have objects to modify with extras
	//Run("import snake")
	snake = py.PyImport_ImportModule("snake")
	if snake == nil {
		fe.Fatal(fmt.Errorf("snake module not available to import"))
	}
	state = py.PyEval_SaveThread()
}

// Call a python function.
//
// Supply gil as true for multithreading call.
// Supply gil as false to use the global initial thread.
func Call(name string, args *py.PyObject, kwargs *py.PyObject, gil bool) *py.PyObject {
	f := snake.GetAttrString(name)
	if f == nil {
		fe.Fatal(fmt.Errorf("snake does not contain %s", name))
	}
	if args == nil {
		args = py.PyTuple_New(0)
	}
	if kwargs == nil {
		kwargs = py.PyDict_New()
	}
	var g py.PyGILState
	if gil {
		// this prevents a deadlock style panic sometimes
		// in scheduling interaction
		runtime.LockOSThread()
		g = py.PyGILState_Ensure()
	} else {
		py.PyEval_RestoreThread(state)
	}
	r := f.Call(args, kwargs)
	if gil {
		py.PyGILState_Release(g)
	} else {
		state = py.PyEval_SaveThread()
	}
	return r
}

func Exit() {
	py.PyEval_RestoreThread(state)
	py.Py_Finalize()
}

// call this from python ?????? TODO, needs cpy3 extending
//
// https://docs.python.org/3/extending/embedding.html#extending-embedded-python
func snake1(self *py.PyObject, args *py.PyObject, kwargs *py.PyObject) *py.PyObject {
	return self
}
