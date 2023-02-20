package problems

func searchInsert(nums []int, target int) int {
	res := len(nums)
	for i, val := range nums {
		if val == target {
			return i
		} else if val > target {
			return i
		}
	}
	return res
}
