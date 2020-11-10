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

func longestValidParenthesesV2(s string) int {
	var left, right, max int = 0, 0, 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			left++
		} else if string(s[i]) == ")" {
			right++
		}
		if left == right {
			if max < right*2 {
				max = right * 2
			}
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if string(s[i]) == ")" {
			right++
		} else if string(s[i]) == "(" {
			left++
		}
		if left == right {
			if max < left*2 {
				max = left * 2
			}
		} else if right < left {
			left, right = 0, 0
		}
	}
	return max
}
