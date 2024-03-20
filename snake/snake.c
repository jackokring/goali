#include "snake.h"

// PyCFunction signature (both the tuple and the dictionary)

PyObject *py_api_stdout(PyObject *self, PyObject *args, PyObject *kwargs) {
	const char *arg = NULL;
	Py_ssize_t len = 0;
	PyArg_ParseTuple(args, "y#", &arg, &len);
	Py_ssize_t i = go_api_stdout((void*)arg, len);
	return Py_BuildValue("n", i);
}

PyObject *py_api_stderr(PyObject *self, PyObject *args, PyObject *kwargs) {
	const char *arg = NULL;
	Py_ssize_t len = 0;
	PyArg_ParseTuple(args, "y#", &arg, &len);
	Py_ssize_t i = go_api_stderr((void*)arg, len);
	return Py_BuildValue("n", i);
}

PyObject *py_api_stdin(PyObject *self, PyObject *args, PyObject *kwargs) {
	Py_ssize_t len = -1;
	PyArg_ParseTuple(args, "|n", &len);
	void *r = go_api_stdinBuffer(len);
	Py_ssize_t ok = go_api_stdinLen();
	PyObject *ret = Py_BuildValue("y#", r, ok);
	go_api_free(r);
	return ret;
}