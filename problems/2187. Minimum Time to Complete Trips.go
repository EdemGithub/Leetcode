package problems

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)
	left, right := 0, time[0]*totalTrips
	for left < right {
		mid := left + (right-left)/2
		trips := 0
		for _, t := range time {
			trips += mid / t
		}
		if trips < totalTrips {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int64(left)
}
