package problems

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min, max := prices[0], 0
	for _, c := range prices {
		if min > c {
			min = c
		} else if max < c-min {
			max = c - min
		}
	}
	return max
}
