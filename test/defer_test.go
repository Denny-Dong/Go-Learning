package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {

	t.Log("TC Start")
	defer t.Log("Defer Run")
	t.Log(deferReturn())
	t.Log("TC End")
}

func deferReturn() string {
	defer fmt.Println("defer fun run")
	return "Fun Return"
}
