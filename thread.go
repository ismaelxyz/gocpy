package gocpy

/*
#include "Python.h"
*/
import "C"

//PyThreadState : https://docs.python.org/3/c-api/init.html#c.PyThreadState
type PyThreadState C.PyThreadState

//PyGILState is an opaque “handle” to the thread state when PyGILState_Ensure() was called, and must be passed to PyGILState_Release() to ensure Python is left in the same state
type PyGILState C.PyGILState_STATE

/*
Deprecated function which does nothing.

In Python 3.6 and older, this function created the GIL if it didn’t exist.

Changed in version 3.9: The function now does nothing.

Changed in version 3.7: This function is now called by Py_Initialize(), so you don’t have to call it yourself anymore.

Changed in version 3.2: This function cannot be called before Py_Initialize() anymore.

Deprecated since version 3.9, will be removed in version 3.11.


PyEval_InitThreads : https://docs.python.org/3/c-api/init.html#c.PyEval_InitThreads
func PyEval_InitThreads() {
	C.PyEval_InitThreads()
}
*/

/*
    Returns a non-zero value if PyEval_InitThreads() has been called. This function can be called without holding the GIL, and therefore can be used to avoid calls to the locking API when running single-threaded.

    Changed in version 3.7: The GIL is now initialized by Py_Initialize().

    Deprecated since version 3.9, will be removed in version 3.11.


PyEval_ThreadsInitialized : https://docs.python.org/3/c-api/init.html#c.PyEval_ThreadsInitialized
func PyEval_ThreadsInitialized() bool {
	return C.PyEval_ThreadsInitialized() != 0
}
*/

//PyEval_SaveThread : https://docs.python.org/3/c-api/init.html#c.PyEval_SaveThread
func PyEval_SaveThread() *PyThreadState {
	return (*PyThreadState)(C.PyEval_SaveThread())
}

//PyEval_RestoreThread : https://docs.python.org/3/c-api/init.html#c.PyEval_RestoreThread
func PyEval_RestoreThread(tstate *PyThreadState) {
	C.PyEval_RestoreThread((*C.PyThreadState)(tstate))
}

//PyThreadState_Get : https://docs.python.org/3/c-api/init.html#c.PyThreadState_Get
func PyThreadState_Get() *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Get())
}

//PyThreadState_Swap : https://docs.python.org/3/c-api/init.html#c.PyThreadState_Swap
func PyThreadState_Swap(tstate *PyThreadState) *PyThreadState {
	return (*PyThreadState)(C.PyThreadState_Swap((*C.PyThreadState)(tstate)))
}

//PyEval_ReInitThreads : https://docs.python.org/3/c-api/sys.html#c.PyOS_AfterFork_Child
func PyOS_AfterFork_Child() {
	C.PyOS_AfterFork_Child()
}

//PyGILState_Ensure : https://docs.python.org/3/c-api/init.html#c.PyGILState_Ensure
func PyGILState_Ensure() PyGILState {
	return PyGILState(C.PyGILState_Ensure())
}

//PyGILState_Release : https://docs.python.org/3/c-api/init.html#c.PyGILState_Release
func PyGILState_Release(state PyGILState) {
	C.PyGILState_Release(C.PyGILState_STATE(state))
}

//PyGILState_GetThisThreadState : https://docs.python.org/3/c-api/init.html#c.PyGILState_GetThisThreadState
func PyGILState_GetThisThreadState() *PyThreadState {
	return (*PyThreadState)(C.PyGILState_GetThisThreadState())
}

//PyGILState_Check : https://docs.python.org/3/c-api/init.html#c.PyGILState_Check
func PyGILState_Check() bool {
	return C.PyGILState_Check() == 1
}
