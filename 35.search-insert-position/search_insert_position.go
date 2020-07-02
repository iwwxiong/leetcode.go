package algorithms35

func searchInsert(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if target <= nums[i] {
			break
		}
	}
	return i
}
