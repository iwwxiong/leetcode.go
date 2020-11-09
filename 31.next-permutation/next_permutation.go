package algorithms31

func swap(nums []int, i, j int) {
	var temp int = nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, i int) {
	var j int = len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func nextPermutation(nums []int) []int {
	var length int = len(nums)
	if length <= 1 {
		return nums
	}
	var i int = length - 1
	for i >= 1 && nums[i-1] >= nums[i] {
		i--
	}
	if i <= 0 {
		reverse(nums, 0)
		return nums
	}
	var j int = length - 1
	for j >= i && nums[i-1] >= nums[j] {
		j--
	}
	swap(nums, i-1, j)
	reverse(nums, i)
	return nums
}
