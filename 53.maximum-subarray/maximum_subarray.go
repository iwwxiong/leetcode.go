package algorithms53

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max int = -2147483647
	var length = len(nums)
	var dp []int = make([]int, length)
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			dp[j] = dp[j] + nums[j+i-1]
			if max < dp[j] {
				max = dp[j]
			}
		}
	}
	return max
}

func maxNum(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func maxSubArrayV1(nums []int) int {
	var length int = len(nums)
	if length == 1 {
		return nums[0]
	}
	var dp []int = make([]int, length)
	var max int = nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		max = maxNum(max, dp[i])
	}
	return max
}
