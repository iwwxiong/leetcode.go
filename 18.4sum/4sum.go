package algorithms18

import (
	"leetcode.go/lib"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int = [][]int{}
	nums = lib.SortIntSlice(nums)
	var length int = len(nums)
	for i := 0; i < length-3; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			for j := i + 1; j < length-2; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					var hp, tp int = j + 1, length - 1
					var temp int = target - nums[i] - nums[j]
					for hp < tp {
						if nums[hp]+nums[tp] == temp {
							result = append(result, []int{nums[i], nums[j], nums[hp], nums[tp]})
							for hp < tp && nums[hp] == nums[hp+1] {
								hp++
							}
							for hp < tp && nums[tp] == nums[tp-1] {
								tp--
							}
							hp++
							tp--
						} else if nums[hp]+nums[tp] < temp {
							hp++
						} else {
							tp--
						}
					}
				}
			}
		}
	}
	return result
}
