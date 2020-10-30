package gocpy

import "fmt"

// PyObjectPrint : equivalent to PyObject_Print
func PyObjectPrint(pyObj *PyObject) error {

	if exc := PyErr_Occurred(); pyObj == nil && exc != nil {
		return fmt.Errorf("Fail to create python object")
	}
	defer pyObj.DecRef()

	repr, err := pythonRepr(pyObj)

	if err != nil {
		return fmt.Errorf("fail to get representation of object")
	}

	fmt.Print(repr)

	return nil
}

func pythonRepr(pyObj *PyObject) (string, error) {
	if pyObj == nil {
		return "", fmt.Errorf("object is nil")
	}

	goObj := pyObj.Repr()
	if goObj == nil {
		PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}
	defer goObj.DecRef()

	return PyUnicode_AsUTF8(goObj), nil
}
