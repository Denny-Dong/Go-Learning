package test

import "testing"

func TestForBreak(t *testing.T) {
	i := 1
LoopTag:
	for {
		for {
			if i > 5 {
				break LoopTag
			}
			t.Log(i)
			i++
		}
	}
}

func TestSwitch(t *testing.T) {
	switch 5 - 3 {
	case 1:
		t.Log("case1")
	case 2:
		t.Log("case2")
	default:
		t.Log("default")
	}
}

func fibonacciNormal(n int) int {
	a := 1
	b := 1
	for i := 2; i < n; i++ {
		c := b
		b = a + b
		a = c
	}
	return b
}

func fibonacciRecursion(n int) (result int) {
	if n <= 2 {
		result = 1
		return result
	} else {
		result = fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
		return result
	}
}
func fibonacciSwitch(n int) (result int) {
	switch n {
	case 1:
		result = 1
		return result
	case 2:
		result = 1
		return result
	default:
		result = fibonacciSwitch(n-1) + fibonacciSwitch(n-2)
		return result
	}

}

func TestFibonacciMultiImplement(t *testing.T) {
	t.Log(fibonacciNormal(10))
	t.Log(fibonacciRecursion(10))
	t.Log(fibonacciSwitch(10))
}
