package test

import (
	"bytes"
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

func TestCharacter(t *testing.T) {
	a, b := 'a', '好'
	c, d := "a", "好"
	t.Log(a, b, c, d)
}

func TestPointer(t *testing.T) {
	var p *int
	a := 1
	p = &a
	t.Log("a : ", a)
	t.Log("&a : ", &a)
	t.Log("p : ", p)
	t.Log("&p : ", &p)
	t.Log("*p : ", *p)
	t.Log("-----------------------")
	*p = 5
	t.Log("a : ", a)
	t.Log("&a : ", &a)
	t.Log("p : ", p)
	t.Log("&p : ", &p)
	t.Log("*p : ", *p)
	t.Log("-----------------------")
	p = new(int)
	t.Log("p : ", p)
	t.Log("&p : ", &p)
	t.Log("*p : ", *p)
}

func TestBytesBuffer(t *testing.T) {
	a := "abc"
	b := "def"
	var c bytes.Buffer
	c.WriteString(a)
	c.WriteString(b)
	t.Log(c)
	t.Log(c.String())
	t.Log(reflect.TypeOf(c))
}
