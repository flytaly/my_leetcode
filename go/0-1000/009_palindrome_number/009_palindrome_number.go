package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0

	for i := 1; i <= x; i *= 10 {
		digit := x % (i * 10) / i
		reverse = reverse*10 + digit
	}

	return reverse == x
}
