package main

import (
	"reflect"
)

type Vertex struct {
	X, Y int
}

func IsNil[T any](i T) bool {
	return reflect.ValueOf(i).IsNil()
}
