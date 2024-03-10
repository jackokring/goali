package snake

//=====================================
//********* Embedded Python ***********
//=====================================

import (
	"os"
	"strconv"

	py "github.com/jackokring/cpy3"
	"github.com/jackokring/goali/filerr"
)

func Run(s string) int {
	return py.PyRun_SimpleString(s)
}

// ummm, needs r/w and map to under-laying files .
//
// ..
func SetIO(r filerr.FilterReader, w filerr.FilterWriter) {
	// assume consistency of process file descriptors
	Run("import sys")
	Run("sys.stdin = os.fdopen(" + strconv.Itoa(int(os.Stdin.Fd())) + ")")
	Run("sys.stout = os.fdopen(" + strconv.Itoa(int(os.Stdout.Fd())) + ")")
	Run("sys.stderr = os.fdopen(" + strconv.Itoa(int(os.Stderr.Fd())) + ")")
}
