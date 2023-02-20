package problems

func romanToInt(s string) int {
	mapa := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := mapa[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if mapa[s[i]] < mapa[s[i+1]] {
			res -= mapa[s[i]]
		} else {
			res += mapa[s[i]]
		}
	}
	return res
}
