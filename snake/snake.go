package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"os"
	"strconv"

	py "github.com/jackokring/cpy3"
)

func SetIO() {
	// assume consistency of process file descriptors
	py.PyRun_SimpleString("sys.stdin = os.fdopen(" + strconv.Itoa(int(os.Stdin.Fd())) + ")")
	py.PyRun_SimpleString("sys.stout = os.fdopen(" + strconv.Itoa(int(os.Stdout.Fd())) + ")")
	py.PyRun_SimpleString("sys.stderr = os.fdopen(" + strconv.Itoa(int(os.Stderr.Fd())) + ")")
}
