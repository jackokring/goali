package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	py "github.com/jackokring/cpy3"
	clit "github.com/jackokring/goali/clitype"
	fe "github.com/jackokring/goali/filerr"
)

//// double wrapped prototypes of functions available from python
//#cgo pkg-config: python3-embed
//#include "snake.h"
import "C"

// In general some functions accept a "gil" bool.
// Supply "gil" as true for multithreading call.
// Supply "gil" as false to use the global initial thread.

type Command struct {
	clit.PyFile
	// umm, target execute function?
	// How many trailing files does you have?
	// I mean I thought of a ADT to rebuild "words"
	// but the time cost per letter was abysmal.
	clit.StreamFilter // embedded type .. => .
}

func (c *Command) Help() string {
	return `An embedded python script interpreter.
A snake module is added into the global name space.
This can be further imported for shorter names.
Input and output are both redirected.`
}

func (c *Command) Run(p *clit.Globals) error {
	// unicorn command hook
	fe.SetGlobals(p)
	r, w := fe.GetIO(c.StreamFilter)
	fe.Lock.Unlock() // IO unlocked
	AddAll(r, w)
	RunFile(c.PyFile, false) // run global (not threaded)
	Exit()
	return nil
}

// Run some python code
func Run(s string, gil bool) {
	Init()
	g := gilStateDefer(gil)
	defer g()
	rtn := py.PyRun_SimpleString(s)
	if rtn != 0 {
		fe.Fatal(fmt.Errorf("python exception: %s", s))
	}
}

// Run a file of python
func RunFile(f clit.PyFile, gil bool) {
	Init()
	//if f.Expand {
	//fe.fe.Fatal(fmt.Errorf("flag -e not allowed: %s", f.InputFile))
	// ignore flag as may be present for data files
	//}
	g := gilStateDefer(gil)
	defer g()
	code, err := py.PyRun_AnyFile(f.PyFile)
	fe.Fatal(err)
	if code != 0 {
		fe.Fatal(fmt.Errorf("python exception in file: %s", f.PyFile))
	}
}

// The master imported object
var snake *py.PyObject

// The thread state
var state *py.PyThreadState

// Python is initialized
var initialized bool

// Python may be a singleton, but a Mutex might be nice
var lock sync.Mutex

// Initialize python
//
// This is often done automatically. You should not need to call it.
func Init() {
	if !errorLock.TryLock() {
		// of course the implicit assumption is grant if available
		// and not don't grant to throttle processing
		fe.Fatal(fmt.Errorf("there can be only one, low lander! reevaluate initial assumptions"))
	}
	lock.Lock()
	defer lock.Unlock()
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
		fe.Fatal(fmt.Errorf("snake module not available to import"))
	}
	state = py.PyEval_SaveThread()
}

// Create a GIL lock if gil is true else use initial thread.
//
// Call the returned function to release the GIL lock.
func gilStateDefer(gil bool) func() {
	// remove the GIL?
	// Just Make a New Process(TM) ...
	// Sure a mini OS in python would be funny
	// but productive on execution context local storage base offset
	// added in to the addressing modes?
	// Sure software TLB is funny, for some ...
	if gil {
		// this prevents a deadlock style panic sometimes
		// in scheduling interaction
		// something to do with go routine preemption and thread reuse
		runtime.LockOSThread()
	} else {
		// main thread
		py.PyEval_RestoreThread(state)
	}
	g := py.PyGILState_Ensure()
	return func() {
		py.PyGILState_Release(g)
		if gil {
			// matched pair
			runtime.UnlockOSThread()
		} else {
			// main thread
			state = py.PyEval_SaveThread()
		}
	}
}

// Call a python function.
func Call(name string, args *py.PyObject, kwargs *py.PyObject, gil bool) *py.PyObject {
	Init()
	g := gilStateDefer(gil)
	defer g()
	f := snake.GetAttrString(name)
	if f == nil {
		fe.Fatal(fmt.Errorf("snake does not contain a global %s", name))
	}
	if !py.PyCallable_Check(f) {
		fe.Fatal(fmt.Errorf("%s is not a global callable", name))
	}
	if args == nil {
		args = py.PyTuple_New(0)
	}
	// kwargs already optimized for a nil -> NULL
	// It's a PyDict_New()
	r := f.Call(args, kwargs)
	return r
}

var errorLock sync.Mutex

// Exit the python interpreter and ensure it is not initialized again.
func Exit() {
	lock.Lock()
	if !initialized {
		lock.Unlock()
		return
	}
	py.PyEval_RestoreThread(state)
	py.Py_Finalize()
	// sure needs a solid lockout
	//lock.Unlock()
	errorLock.Lock()
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
// 1. Decide on function names and make a go function below here
// 2. Wrap Go_api of go function (for py_api to call) below here
//		It is critical that it is comment marked "//export <name>"
// 3. Add Go_api extern forward prototype at in .h file (in C)
// 4. Add py_api to .h file (in C)
// 5. Wonder at the C ... rap jokes ...

// Add a C function to the python snake module
func AddFunc(name string, function unsafe.Pointer) {
	// remove old before new?
	// not sure if it's needed but ...
	// allows "snake.py" to have dummy mypy functions
	Init()
	if snake.DelAttrString(name) != 0 {
		fe.Fatal(fmt.Errorf("%s has no global template in the snake module", name))
	}
	if snake.AddModuleCFunction(name, function) != 0 {
		fe.Fatal(fmt.Errorf("%s couldn't be added to the snake module", name))
	}
}

type io struct {
	fe.FilterReader
	fe.FilterWriter
}

var files io

// Add all the default functions to the snake module
//
// Also set the IO used for standard IO
func AddAll(r fe.FilterReader, w fe.FilterWriter) {
	files = io{
		FilterReader: r,
		FilterWriter: w,
	}
	AddFunc("Out", C.py_api_stdout)
	AddFunc("Err", C.py_api_stderr)
	AddFunc("In", C.py_api_stdin)
}

//=====================================
//****** Extensions For Python ********
//=====================================

// e.g.	//return C.CString(snake1(C.GoString(a), C.GoString(b)))

// IO use the reader and writer ...

//export go_api_stdout
func go_api_stdout(s unsafe.Pointer, n C.int) C.int {
	return C.int(stdout(C.GoBytes(s, n)))
}

func stdout(s []byte) int {
	return files.FilterWriter.Write(s)
}

//export go_api_stderr
func go_api_stderr(s unsafe.Pointer, n C.int) C.int {
	return C.int(stderr(C.GoBytes(s, n)))
}

func stderr(s []byte) int {
	fe.Notify(s)
	// fail or true fact
	return len(s)
}

//export go_api_stdinBuffer
func go_api_stdinBuffer(s C.int) unsafe.Pointer {
	return C.CBytes(stdin(int(s))) // malloc a buffer
}

//export go_api_stdinLen
func go_api_stdinLen() C.int {
	return C.int(stdinLen)
}

//export go_api_free
func go_api_free(s unsafe.Pointer) {
	C.free(s)
}

var stdinLen int

func stdin(size int) []byte {
	if size == -1 {
		r := make([]byte, 0)
		// all the file as one buffer
		for !files.FilterReader.EOF() {
			i := stdin(1024)
			r = append(r, i...) // automatic varadic expansion
		}
		stdinLen = len(r)
		return r
	} else {
		r := make([]byte, size)
		i := files.FilterReader.Read(r)
		stdinLen = i
		return r[:i] // python EOF style
	}
}
