int main()
{

// Initialize the main interpreter
Py_Initialize();
// Initialize and acquire the global interpreter lock
PyEval_InitThreads();
// Release the lock     
PyThreadState *_save;
_save = PyEval_SaveThread();


// create threads
for (int i = 0; i<MAX_THREADS; i++)
{

    // Create the thread to begin execution on its own.

    hThreadArray[i] = CreateThread
    //(...

        MyThreadFunction,       // thread function name
    //...);   // returns the thread identifier 

} // End of main thread creation loop.

  // Wait until all threads have terminated.
WaitForMultipleObjects(MAX_THREADS, hThreadArray, TRUE, INFINITE);

// Close all thread handles and free memory allocations.
// ...


//end python here
//but need to check for GIL here too
//re capture the lock
PyEval_RestoreThread(_save);
//end python interpreter
Py_Finalize();
return 0;
}

//the thread functions
DWORD WINAPI MyThreadFunction(LPVOID lpParam)
{
// Non Pythonic activity
// ...

//create a new interpreter
PyEval_AcquireLock(); // acquire lock on the GIL
PyThreadState* pThreadState = Py_NewInterpreter();
assert(pThreadState != NULL); // check for failure
PyEval_ReleaseThread(pThreadState); // release the GIL


// switch in current interpreter
PyEval_AcquireThread(pThreadState);

//execute python code
PyRun_SimpleString("from time import time,ctime\n" "print\n"
    "print 'Today is',ctime(time())\n");

// release current interpreter
PyEval_ReleaseThread(pThreadState);

//now to end the interpreter
PyEval_AcquireThread(pThreadState); // lock the GIL
Py_EndInterpreter(pThreadState);
PyEval_ReleaseLock(); // release the GIL

// Other non Pythonic activity
return 0;
}