package algorithms66

func plusOne(digits []int) []int {
	var isPlus bool
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			isPlus = false
			break
		}
		digits[i] = 0
		isPlus = true
	}
	if isPlus {
		digits = append([]int{1}, digits...)
	}
	return digits
}
