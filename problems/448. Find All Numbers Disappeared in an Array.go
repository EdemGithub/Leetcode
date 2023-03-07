package problems

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		pos := nums[i] - 1
		if nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}
	fmt.Println(nums)
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
