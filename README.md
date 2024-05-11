# GoCPy
> The Gocpy project was archived on 5/11/2024. I don't see myself working any more on this beautiful project, it seems very noble to me but I have other duties to attend to and this requires a commitment that I don't want to face. It has been a pleasure : )
> 
**Currently supports python-3.9.x+ only.**

Golang bindings for the C-API of CPython3.

This package provides a ``go`` package named "GoCPy" under which most of the
``PyXYZ`` functions and macros of the public C-API of CPython have been
exposed. Theoretically, you should be able use https://docs.python.org/3/c-api
and know what to type in your ``go`` program.

## Deps

Python-dev

# Install

If your operating system is Linux-based and your *Python.h* is located at
/usr/local/include/python3.9 and your *libpython3.9.a* 
(Static python library) located at 
/usr/local/lib/python3.9/config-3.9-x86_64-linux-gnu, run:
```bash
go get -u github.com/ismaelxyz/gocpy
```
and done!, else:

1) Download [GoCPy](https://codeload.github.com/ismaelxyz/gocpy/zip/main)
2) Extract Zip file.
3) Edit: *high_level_layer.go*
  3.1) Change: /usr/local/include/python3.9, by path to *Python.h*
  3.2) Change: /usr/local/lib/python3.9/config-3.9-x86_64-linux-gnu, by path to
  *libpython3.9.a* or *python39.lib* in windows.
4) In the GoCPy path run: go install
5) Done!

NOTES:
* In Windows the default path to Python is C:\Users\[user]\AppData\Local\Programs\Python\Python3.x
* In Linux the default path to Python (include) is /usr/include/python/pythonx.x and 
python lib /usr/lib/pythonx.x.

## Go get

Very simple: `go get -u github.com/ismaelxyz/gocpy`

# API

Some functions mix go code and call to Python function. Those functions will
return and `int` and `error` type. The `int` represent the Python result code
and the `error` represent any issue from the Go layer.

Example:

`func PyRun_AnyFile(filename string)` open `filename` and then call CPython API
function `int PyRun_AnyFile(FILE *fp, const char *filename)`.

Therefore its signature is `func PyRun_AnyFile(filename string) (int, error)`,
the `int` represent the error code from the CPython `PyRun_AnyFile` function
and error will be set if we failed to open `filename`.

If an error is raise before calling th CPython function `int` default to `-1`.

Take a look at some [examples](examples)
