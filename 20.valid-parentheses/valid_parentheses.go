package algorithms20

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var stack []string
	for _, str := range s {
		if string(str) == "(" || string(str) == "[" || string(str) == "{" {
			stack = append(stack, string(str))
		} else {
			if len(stack) == 0 {
				stack = append(stack, string(str))
				break
			}
			if string(str) == ")" {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, string(str))
					break
				}
			} else if string(str) == "]" {
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, string(str))
					break
				}
			} else if string(str) == "}" {
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, string(str))
					break
				}
			}
		}
	}
	return len(stack) == 0
}
