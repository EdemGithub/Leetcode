package problems

func maxSubArray(nums []int) int {
	max, curr := nums[0], nums[0]
	for _, n := range nums[1:] {
		if sum := curr + n; n > sum {
			curr = n
		} else {
			curr = sum
		}
		if curr > max {
			max = curr
		}
	}
	return max
}
