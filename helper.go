package gocpy

//go:generate go run script/variadic.go


/*
#include <stdio.h>
#include "Python.h"*/
import "C"

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}

/*Use this for call functions with args of type *FILE */

// Stderr : Estandar stderr.
func Stderr() *C.FILE {
    return C.stderr
}

// Stdin : Estandar stdin.
func Stdin() *C.FILE {
    return C.stdin
}

// Stdout : Estandar stdout.
func Stdout() *C.FILE {
    return C.stdout
}
