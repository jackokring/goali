package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"fmt"
	"runtime"
	"unsafe"

	py "github.com/jackokring/cpy3"
	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

/* // double wrapped prototypes of functions available from python
#cgo pkg-config: python3-embed

#define PY_SSIZE_T_CLEAN
#include <Python.h>

PyObject *py_api_snake1(PyObject *self, PyObject *args, PyObject *kwargs);
*/
import "C"

type Command struct {
	clit.PyFile
	// umm, target execute function?
	// How many trailing files does you have?
	// I mean I thought of a ADT to rebuild "words"
	// but the time cost per letter was abysmal.
	clit.StreamFilter // embedded type .. => .
}

func (c *Command) Help() string {
	return `An embedded python script interpreter.`
}

func (c *Command) Run(p *clit.Globals) error {
	// unicorn command hook
	fe.SetGlobals(p)

	return nil
}

func Fatal(e error) {
	if e != nil {
		Exit()
	}
	fe.Fatal(e)
}

func Run(s string) {
	Init()
	if py.PyRun_SimpleString(s) != 0 {
		Fatal(fmt.Errorf("python exception: %s", s))
	}
}

func RunFile(f clit.InputFile) {
	Init()
	//if f.Expand {
	//fe.Fatal(fmt.Errorf("flag -e not allowed: %s", f.InputFile))
	// ignore flag as may be present for data files
	//}
	code, err := py.PyRun_AnyFile(f.InputFile)
	Fatal(err)
	if code != 0 {
		Fatal(fmt.Errorf("python exception in file: %s", f.InputFile))
	}
}

// The master imported object
var snake *py.PyObject

// The thread state
var state *py.PyThreadState

// Python is initialized
var initialized bool

func Init() {
	if initialized {
		return
	}
	initialized = true
	py.Py_Initialize()
	// To be usable as a part of snake
	// Must have objects to modify with extras
	//Run("import snake")
	snake = py.PyImport_ImportModule("snake")
	if snake == nil {
		Fatal(fmt.Errorf("snake module not available to import"))
	}
	state = py.PyEval_SaveThread()
}

// Call a python function.
//
// Supply gil as true for multithreading call.
// Supply gil as false to use the global initial thread.
func Call(name string, args *py.PyObject, kwargs *py.PyObject, gil bool) *py.PyObject {
	Init()
	f := snake.GetAttrString(name)
	if f == nil {
		Fatal(fmt.Errorf("snake does not contain a global %s", name))
	}
	if !py.PyCallable_Check(f) {
		Fatal(fmt.Errorf("%s is not a global callable", name))
	}
	if args == nil {
		args = py.PyTuple_New(0)
	}
	// kwargs already optimized for a nil -> NULL
	// It's a PyDict_New()
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
	if !initialized {
		return
	}
	py.PyEval_RestoreThread(state)
	py.Py_Finalize()
}

//=====================================
//***** Extended Embedded Python ******
//=====================================

// added py.addModuleCFunctions(module, CFuncPointer) int.
//
// https://docs.python.org/3/extending/embedding.html#extending-embedded-python
//
//func snake1(self *py.PyObject, args *py.PyObject, kwargs *py.PyObject) *py.PyObject {
//	return self
//}
//
// N.B. Only the *self (internal?), *args (tuple), *kwargs (dictionary) API
//
// 1. So link py_api functions to names
// 2. Wrap go_api of go function
// 3. Add py_api to go_api link in cwrap.go
// 4. Add prototype of py_api to top of this file
// 5. Wonder at the C ... rap jokes ...

func AddFunc(name string, function unsafe.Pointer) {
	// remove old before new?
	// not sure if it's needed but ...
	// allows "snake.py" to have dummy mypy functions
	Init()
	if snake.DelAttrString(name) != 0 {
		Fatal(fmt.Errorf("%s has no global template in the snake module", name))
	}
	if snake.AddModuleCFunction(name, function) != 0 {
		Fatal(fmt.Errorf("%s couldn't be added to the snake module", name))
	}
}

var files struct {
	fe.FilterReader
	fe.FilterWriter
}

func AddAll(r fe.FilterReader, w fe.FilterWriter) {
	AddFunc("snake1", C.py_api_snake1)
}

//=====================================
//****** Extensions For Python ********
//=====================================

func snake1(a string, b string) string {
	return a + b
}

//export go_api_snake1
func go_api_snake1(a *C.char, b *C.char) *C.char {
	return C.CString(snake1(C.GoString(a), C.GoString(b)))
}

// IO use the reader and writer ...

func stdout(s []byte) int {
	return files.FilterWriter.Write(s)
}

func stderr(s []byte) int {
	fe.Notify(s)
	// fail or true fact
	return len(s)
}

func stdin(size int) []byte {
	if size == -1 {
		r := make([]byte, 0)
		// all the file as one buffer
		for !files.FilterReader.EOF() {
			i := stdin(1024)
			r = append(r, i...) // automatic varadic expansion
		}
		return r
	} else {
		r := make([]byte, size)
		i := files.FilterReader.Read(r)
		return r[:i] // python EOF style
	}
}
