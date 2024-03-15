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

char *go_api_snake1(char*, char*);

// PyCFunction signature (both the tuple and the dictionary)
PyObject *py_api_snake1(PyObject *self, PyObject *args, PyObject *kwargs) {
	char *arg1 = NULL, *arg2 = NULL;
	PyArg_ParseTuple(args, "ss", &arg1, &arg2);
	char *r = go_api_snake1(arg1, arg2);
	return PyUnicode_FromString(r);
}

*/
import "C"
