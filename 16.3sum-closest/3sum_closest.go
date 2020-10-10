package algorithms16

import (
	"leetcode.go/lib"
)

func abs(num int) int {
	if num < 0 {
		return 0 - num
	} else {
		return num
	}
}

func threeSumClosest(nums []int, target int) int {
	nums = lib.SortIntSlice(nums)
	var max int = 100000000000
	var sum int
	var length int = len(nums)
	for i := 0; i <= length-1; i++ {
		var hp, tp int = i + 1, length - 1
		for hp < tp {
			if abs(nums[i]+nums[hp]+nums[tp]-target) < max {
				sum = nums[i] + nums[hp] + nums[tp]
				max = abs(sum - target)
			}
			if nums[i]+nums[hp]+nums[tp] > target {
				tp--
			} else {
				hp++
			}
		}
	}
	return sum
}
