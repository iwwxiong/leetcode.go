package algorithms32

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	var max int = 0
	var stack []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, string(s[i]))
		} else if string(s[i]) == ")" {
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				if max < i-len(stack)+1 {
					max = i - len(stack) + 1
				}
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}
	return max
}
