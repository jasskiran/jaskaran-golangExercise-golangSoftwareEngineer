package robbery

func Robbery(nums []int) int {
	n := len(nums)
	//first := nums[0]
	//last := nums[n-1]
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return 0
	}
	if n == 3 {
		return nums[1]
	}

	return max(rob(nums, 0, n-2), rob(nums, 1, n-1))
}

func rob(nums []int, start, end int) int {
	preMax := nums[start]
	curMax := max(preMax, nums[start+1])
	for i := start + 2; i <= end; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
