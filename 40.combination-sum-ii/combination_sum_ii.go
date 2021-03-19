package algorithms40

import "leetcode.go/lib"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var temp []int
	candidates = lib.SortIntSlice(candidates)
	backtrack(&result, &temp, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, temp *[]int, candidates []int, target int, start int) {
	if target < 0 {
		return
	} else if target == 0 {
		add := make([]int, len(*temp))
		copy(add, *temp)
		*result = append(*result, add)
	} else {
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			*temp = append(*temp, candidates[i])
			backtrack(result, temp, candidates, target-candidates[i], i+1)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}
