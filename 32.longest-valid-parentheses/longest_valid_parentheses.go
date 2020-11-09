package algorithms32

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	var max int = 0
	var stack []int = []int{-1}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if max < i-stack[len(stack)-1] {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}
