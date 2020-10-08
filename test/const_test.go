package test

import (
	"reflect"
	"testing"
)

func TestConst(t *testing.T) {
	const a = 3.1415
	const b = "Hello World"
	t.Log(reflect.TypeOf(a))
}

const (
	a = iota
	b
	c
	d
	e = 100
)

func TestMultiConst(t *testing.T) {
	t.Log("a : ", a)
	t.Log("b : ", b)
	t.Log("c : ", c)
	t.Log("d : ", d)
	t.Log("e : ", e)
}
