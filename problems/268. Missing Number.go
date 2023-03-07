package problems

func missingNumber(nums []int) int {
	l := len(nums)
	mapa := make(map[int]bool)

	for _, c := range nums {
		mapa[c] = true
	}
	for i := 0; i < l; i++ {
		if _, ok := mapa[i]; !ok {
			return i
		}
	}
	return l
}
