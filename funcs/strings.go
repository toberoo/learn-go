package funcs

// isPalindrome checks if the string is a palindrome
func isPalindrome(s string) bool {
	if len(s) < 2 {
		return false
	}

	strArr := []rune(s)

	start := 0
	end := len(s) - 1

	for end > start {
		if strArr[start] != strArr[end] {
			return false
		}
	}

	return true
}

// isUnique determines if all the characters in an array are unique
func isUnqiue(s string) bool {
	if len(s) < 2 {
		return false
	}

	strArray := []rune(s)
	m := make(map[rune]bool)

	for i := 0; i < len(strArray); i++ {
		if m[strArray[0]] == true {
			return false
		}
		m[strArray[0]] = true
	}

	return true
}
