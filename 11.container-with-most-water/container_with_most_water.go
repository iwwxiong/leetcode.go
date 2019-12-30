package algorithms11

func maxArea(height []int) int {
	var max, maxTemp int
	left, right := 0, len(height)-1
	for left != right {
		if height[left] > height[right] {
			maxTemp = height[right] * (right - left)
		} else {
			maxTemp = height[left] * (right - left)
		}
		if maxTemp > max {
			max = maxTemp
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
