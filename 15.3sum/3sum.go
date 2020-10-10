package algorithms15

import (
	"leetcode.go/lib"
)

func threeSum(nums []int) [][]int {
	var length int = len(nums)
	var result [][]int
	if length < 3 {
		return result
	}
	nums = lib.SortIntSlice(nums)
	for i := 0; i < length-1; i++ {
		// 去重
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			var hp, tp int = i + 1, length - 1 // 头尾指针
			var num = 0 - nums[i]
			for hp < tp {
				if nums[hp]+nums[tp] == num {
					result = append(result, []int{nums[i], nums[hp], nums[tp]})
					for hp < tp && nums[hp] == nums[hp+1] {
						hp++
					}
					for hp < tp && nums[tp] == nums[tp-1] {
						tp--
					}
					hp++
					tp--
				} else if nums[hp]+nums[tp] < num {
					hp++
				} else {
					tp--
				}
			}
		}
	}
	return result
}
