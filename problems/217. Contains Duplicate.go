package problems

func containsDuplicate(nums []int) bool {
	mapa := make(map[int]bool)

	for _, c := range nums {
		if _, ok := mapa[c]; !ok {
			mapa[c] = true
		} else {
			return true
		}
	}
	return false
}
