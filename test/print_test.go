package test

import "testing"

func TestPoint(t *testing.T) {
	num := 1
	var p *int
	p = &num
	t.Log(p)
	t.Log(*p)
	t.Log(&p)
}
