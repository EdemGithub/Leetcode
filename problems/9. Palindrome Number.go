package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	temp, res := x, 0

	for temp != 0 {
		res = res*10 + temp%10
		temp = temp / 10
	}

	return (x == res)
}
