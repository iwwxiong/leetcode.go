package algorithms7

func reverse(x int) int {
	var y = int32(x)
	var result int32
	for y != 0 {
		temp := result*10 + y%10
		if temp/10 != result {
			return 0
		}
		result = temp
		y /= 10
	}
	return int(result)
}
