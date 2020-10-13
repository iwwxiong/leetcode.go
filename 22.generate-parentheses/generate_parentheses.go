package algorithms22

func generate(res *[]string, current string, left, right, n int) {
	if len(current) == n*2 {
		*res = append(*res, current)
		return
	}
	if left < n {
		generate(res, current+"(", left+1, right, n)
	}
	if right < left {
		generate(res, current+")", left, right+1, n)
	}
}

func generateParenthesis(n int) []string {
	var res []string
	generate(&res, "", 0, 0, n)
	return res
}
