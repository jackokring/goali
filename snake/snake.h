#define PY_SSIZE_T_CLEAN
#include <Python.h>

extern int go_api_stdout(void*, int);
extern int go_api_stderr(void*, int);
extern void* go_api_stdinBuffer(int);
extern int go_api_stdinLen();
extern void go_api_free(void*);

// PyCFunction signature (both the tuple and the dictionary)

extern PyObject *py_api_stdout(PyObject *self, PyObject *args, PyObject *kwargs);
extern PyObject *py_api_stderr(PyObject *self, PyObject *args, PyObject *kwargs);
extern PyObject *py_api_stdin(PyObject *self, PyObject *args, PyObject *kwargs);