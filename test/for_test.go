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
