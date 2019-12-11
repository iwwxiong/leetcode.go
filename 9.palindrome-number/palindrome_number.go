package algorithms9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var y, temp int
	temp = x
	for temp != 0 {
		y = y*10 + temp%10
		temp = temp / 10
	}
	if x == y {
		return true
	}
	return false
}
