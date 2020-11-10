package algorithms33

func search(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	var result int = -1
	for left <= right {
		var mid int = (right + left) / 2
		if nums[mid] == target {
			result = mid
			break
		}
		// 左边有序
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return result
}
