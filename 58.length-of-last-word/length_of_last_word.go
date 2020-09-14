package algorithms58

func lengthOfLastWord(s string) int {
	var length int = len(s)
	var n int
	for i := 0; i < length; i++ {
		if n == 0 {
			if s[length-i-1] != ' ' {
				n++
			}
		} else {
			if s[length-i-1] == ' ' {
				break
			}
			n++
		}
	}
	return n
}
