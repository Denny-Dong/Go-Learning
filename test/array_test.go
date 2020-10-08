package test

import (
	"testing"
)

func TestArray(t *testing.T) {
	var student = [5]string{"a", "b", "c", "d", "e"}
	var teacher = [...]string{"a", "b", "c"}
	t.Log(student)
	t.Log(teacher)
	for index, value := range teacher {
		t.Log(index)
		t.Log(value)
	}
}
