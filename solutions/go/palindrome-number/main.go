package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	b := len(str)

	for a := 0; a < len(str); a++ {
		b--
		if str[a] != str[b] {
			return false
		}
	}
	return true

}

//https://leetcode.com/problems/palindrome-number/
