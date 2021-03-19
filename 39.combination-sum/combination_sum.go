package algorithms39

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var temp []int
	backtrack(&result, &temp, candidates, target, 0)
	return result
}

// 回溯
func backtrack(result *[][]int, temp *[]int, nums []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		add := make([]int, len(*temp))
		copy(add, *temp)
		*result = append(*result, add)
		// *result = append(*result, *temp)
	} else {
		for i := start; i < len(nums); i++ {
			*temp = append(*temp, nums[i])
			backtrack(result, temp, nums, remain-nums[i], i)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}
