package problems

func findDuplicates(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res = append(res, nums[i])
			}
		}
	}
	return res
}
