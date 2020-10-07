package test

import "testing"

func TestVariable(t *testing.T) {

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
