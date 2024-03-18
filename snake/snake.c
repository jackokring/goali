#include "snake.h"

// PyCFunction signature (both the tuple and the dictionary)

PyObject *py_api_stdout(PyObject *self, PyObject *args, PyObject *kwargs) {
	const char *arg = NULL;
	Py_ssize_t len = 0;
	PyArg_ParseTuple(args, "y#", &arg, &len);
	int i = go_api_stdout((void*)arg, (int)len);
	return Py_BuildValue("i", i);
}

PyObject *py_api_stderr(PyObject *self, PyObject *args, PyObject *kwargs) {
	const char *arg = NULL;
	Py_ssize_t len = 0;
	PyArg_ParseTuple(args, "y#", &arg, &len);
	int i = go_api_stderr((void*)arg, (int)len);
	return Py_BuildValue("i", i);
}

PyObject *py_api_stdin(PyObject *self, PyObject *args, PyObject *kwargs) {
	Py_ssize_t len = -1;
	PyArg_ParseTuple(args, "|i", &len);
	void *r = go_api_stdinBuffer((int)len);
	int ok = go_api_stdinLen();
	PyObject *ret = Py_BuildValue("y#", r, ok);
	go_api_free(r);
	return ret;
}