package algorithms34

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	var result []int = []int{-1, -1}
	var left, right int = 0, len(nums) - 1
	for left <= right {
		var mid int = (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left == len(nums) || nums[left] != target {
		// 未找到
		return result
	}
	result[0] = left
	// left = 0
	right = len(nums) - 1
	for left <= right {
		var mid int = (left + right) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	result[1] = right
	return result
}
