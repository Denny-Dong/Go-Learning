package test

import (
	"bytes"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	a := "abc"
	b := "def"
	var c bytes.Buffer
	c.WriteString(a)
	c.WriteString(b)
	t.Log(c)
	t.Log(c.String())
	t.Log(reflect.TypeOf(c))
}
