package snake

//=====================================
//***** Extended Embedded Python ******
//=====================================

// Here we define Python wrappers in C for our "go" functions
// They handle the conversion between Python and C types

// Unfortunately we must use C and not Go to write these functions because:
// 1) the "PyArg_Parse_" functions all use variadic arguments, which are not supported by cgo, and
// 2) the "PyArg_Parse_" functions unpack arguments to pointers, which we cannot implement in Go

// Note: cgo exports cannot be in the same file as cgo preamble functions,
// which is why this file cannot be combined with "snake.go"
// and is also why must forward-declare the "snake" functions in the cgo preamble here

/* // better for the line beginnings
#cgo pkg-config: python3-embed

#define PY_SSIZE_T_CLEAN
#include <Python.h>

int go_api_stdout(const char*, int);
int go_api_stderr(const char*, int);
char *go_api_stdinBuffer(int);
int go_api_stdinLen();
void go_api_free(const char*);

// PyCFunction signature (both the tuple and the dictionary)

PyObject *py_api_stdout(PyObject *self, PyObject *args, PyObject *kwargs) {
	const char *arg = NULL;
	Py_ssize_t len = 0;
	PyArg_ParseTuple(args, "y#", &arg, &len);
	int i = go_api_stdout(arg, (int)len);
	return Py_BuildValue("i", i);
}

PyObject *py_api_stderr(PyObject *self, PyObject *args, PyObject *kwargs) {
	const char *arg = NULL;
	Py_ssize_t len = 0;
	PyArg_ParseTuple(args, "y#", &arg, &len);
	int i = go_api_stderr(arg, (int)len);
	return Py_BuildValue("i", i);
}

PyObject *py_api_stdin(PyObject *self, PyObject *args, PyObject *kwargs) {
	Py_ssize_t len = -1;
	PyArg_ParseTuple(args, "|i", &len);
	char *r = go_api_stdinBuffer((int)len);
	int ok = go_api_stdinLen();
	PyObject *ret = Py_BuildValue("y#", r, ok);
	go_api_free(r);
	return ret;
}
*/
import "C"
