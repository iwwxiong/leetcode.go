package algorithms5

func longestPalindrome(s string) string {
	var palindrome string
	if len(s) <= 1 {
		return s
	}
	if len(s) == 2 && s[0] == s[1] {
		return s
	}

	palindrome = string(s[0])

	for i := 0; i < len(s)-1; i++ {
		j := 1
		for i-j >= 0 && i+j <= len(s)-1 {
			if s[i-j] != s[i+j] {
				break
			}
			if 2*j+1 > len(palindrome) {
				palindrome = s[i-j : i+j+1]
			}
			j++
		}
		if s[i] == s[i+1] {
			if len(palindrome) < 2 {
				palindrome = s[i : i+2]
			}
			k := 1
			for i-k >= 0 && i+1+k <= len(s)-1 {
				if s[i-k] != s[i+1+k] {
					break
				}
				if 2*k+2 > len(palindrome) {
					palindrome = s[i-k : i+k+2]
				}
				k++
			}
		}
	}

	return palindrome
}
