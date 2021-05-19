package test

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func returnErr() error {
	return errors.New("newErr")
}

func TestString(t *testing.T) {
	a := "abc"
	b := "def"
	var c bytes.Buffer
	c.WriteString(a)
	c.WriteString(b)
	t.Log(c)
	t.Log(c.String())
	t.Log(reflect.TypeOf(c))

	a = "ssss"
	t.Log(a)

	err := errors.New("")
	err = returnErr()
	t.Log(err)
}
