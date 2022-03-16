package combinatronics

import "testing"

func TestCross(t *testing.T) {
	xs := [][]int{
		{2, 3},
		{2, 4},
	}
	res := Cross(1, xs)

	t.Log(res)
}

func TestCombs(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	res := Combs(xs, 3)

	if ok := len(res) == 10; !ok {
		t.Error("no coinciden")
	}

	t.Log(res)
}

func TestSubsets(t *testing.T) {
	xs := []int{1, 2, 3}
	res := Subsets(xs, 3)

	// 1,2,3,12,13,23,123

	if ok := len(res) == 7; !ok {
		t.Error("no coinciden")
	}

	t.Log(res)
}
