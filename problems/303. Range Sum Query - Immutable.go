package problems

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	var sum []int
	curr_sum := 0
	for _, num := range nums {
		curr_sum += num
		sum = append(sum, curr_sum)
	}
	return NumArray{arr: sum}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.arr[right]
	} else {
		return this.arr[right] - this.arr[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
