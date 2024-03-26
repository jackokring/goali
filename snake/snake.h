#define PY_SSIZE_T_CLEAN
#include <Python.h>

extern Py_ssize_t go_api_stdout(void*, Py_ssize_t);
extern Py_ssize_t go_api_stderr(void*, Py_ssize_t);
extern void* go_api_stdinBuffer(Py_ssize_t);
extern Py_ssize_t go_api_stdinLen();
extern void go_api_free(void*);
extern void go_api_action_msg(char*, Py_ssize_t);

// PyCFunction signature (both the tuple and the dictionary)

extern PyObject *py_api_stdout(PyObject *self, PyObject *args, PyObject *kwargs);
extern PyObject *py_api_stderr(PyObject *self, PyObject *args, PyObject *kwargs);
extern PyObject *py_api_stdin(PyObject *self, PyObject *args, PyObject *kwargs);
extern PyObject *py_api_action_msg(PyObject *self, PyObject *args, PyObject *kwargs);