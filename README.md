# GoCPy
Go binding with Python

## Description
A binding of Go with Python. That in the future will be part of
[Minecraft](https://github.com/jason-bowen-zheng/Minecraft).

Inspired by [go-python3](https://github.com/DataDog/go-python3) this is a
lightweight binding, tested with Python3.9.0 and adapted to the needs of
project which it will serve.

## Installation
The following steps must be strictly followed!

1) Create a variable in the terminal whose name will be CGO_CFLAGS and its
   value "-I" linked to the path of the source files of Python.

2) Create a variable in the terminal whose name will be CGO_LDFLAGS and its
   value "-L" attached to the path that leads to the files .lib of Python
   followed by: " -lpython -lpython39" (set the spaces and the pyton39
   corresponding to the python version installed on the operating system and
   the source code files).

3) Run the command: go to install -u github.com/ismaelxyz/gocpy


As an example in Windows, steps 1 and 2 would be carried out as follows:

set CGO_CFLAGS=-IC:/Users/user/AppData/Local/Programs/Python/Python39-32/include
set CGO_LDFLAGS=-LC:/Users/user/AppData/Local/Programs/Python/Python39-32/libs -lpython3 -lpython39

Assuming that the Python source code files are in
C:/Users/user/AppData/Local/Programs/Python/Python39-32/include and files
.lib of Python in C:/Users/user/AppData/Local/Programs/Python/Python39-32/libs.

## Use
Before using this package, you must perform steps 1 and 2 of the installation
in the terminal to be used; **no need to repeat them, vast with the execution of
the same once.

## Examples

```Go
package main
import (
   "fmt"

   "https://github.com/ismaelxyz/gocpy"
)

func main() {
   gocpy.Py_Initialize()

   
   if !gocpy.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
   }
   
   pyString := gocpy.PyUnicode_FromString("Hello GoCPy")
   pyList := gocpy.PyList_New(10)
   goInt := gocpy.PyList_Size(pyList)

   fmt.Println("The size of the list is:", goInt)
   gocpy.Py_Finalize()

}
```

```Go
package main
import (

   "https://github.com/ismaelxyz/gocpy"
)

func main() {
   gocpy.Py_Initialize()

   
   if !gocpy.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
   }
   
   pyString := gocpy.PyUnicode_FromString("Hello, I am a Python object.")

   PyObjectPrint(pyString)

   gocpy.Py_Finalize()

}
```


## Note
If you see this message it is because this package is considered in its beta
version and you should take precautions when installing or using it on your
device.