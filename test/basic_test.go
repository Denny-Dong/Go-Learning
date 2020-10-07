package test

import (
	"reflect"
	"testing"
)

func TestVariable(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a)
	t.Log(b)
}

func TestFibonacci(t *testing.T) {
	a := 1
	b := 1
	t.Log(a)
	t.Log(b)
	for i := 0; i < 10; i++ {
		c := a + b
		t.Log(c)
		a = b
		b = c
	}
}

func TestFloat(t *testing.T) {
	a, b := 3.0, 2.0
	t.Log(a / b)
	t.Log(reflect.TypeOf(a))
	t.Log(reflect.TypeOf(b))
}

func TestStringMultiLine(t *testing.T) {
	a := `First Line
Secend Line
Third Line
`
	t.Log(a)
}
