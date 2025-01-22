package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	reverse := 0
	for {
		if temp == 0 {
			break
		}

		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}

	return reverse == x
}

func isPalindromeFaster(x int) bool {
	if x < 0 {
		return false
	}

	res := 0
	t := x

	for t != 0 {
		res = res*10 + (t % 10)
		t /= 10
	}
	return res == x
}
