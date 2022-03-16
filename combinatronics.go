package combinatronics

func Cross(h int, xs [][]int) [][]int {
	res := make([][]int, len(xs))
	for i, x := range xs {
		res[i] = append(x, h)
	}
	return res
}

func Combs(xs []int, n int) [][]int {
	// base case
	if n == 1 {
		res := make([][]int, len(xs))
		for i, v := range xs {
			res[i] = []int{v}
		}
		return res
	}

	// inductive case
	res := make([][]int, 0)
	for i := 0; i < len(xs)-n+1; i++ {
		h := xs[i]
		r := xs[i+1:]
		t := Combs(r, n-1)
		c := Cross(h, t)
		res = append(res, c...)
	}

	return res
}

func Subsets(xs []int, n int) [][]int {
	res := make([][]int, 0)
	for i := 1; i <= n; i++ {
		c := Combs(xs, i)
		res = append(res, c...)
	}
	return res
}

func Dot(xs []int) int {
	res := xs[0]
	for _, x := range xs[1:] {
		res *= x
	}
	return res

}
