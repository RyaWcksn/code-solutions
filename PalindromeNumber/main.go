package main

import "fmt"

func isPalindrome(x int) bool {
	comaparison := x
	var res int = 0
	var remainder int
	for x > 0 {
		remainder = x % 10
		res = (res * 10) + remainder
		x = x / 10
	}
	// 121
	// 121 % 10 = 1
	// hasil = (0 * 10) + 1 = 1
	// x = 121 / 10 = 12.1

	// 12
	// 12 % 10 = 2
	// hasil = (1 * 10 ) + 2 = 12
	// x = 12 / 10 = 1

	// 1
	// 1 % 10 = 1
	// hasil = (12 * 10) + 1 = 121
	// x = 1 / 10 = 0

	if res != comaparison {
		return false
	}
	return true
}

func isWordPalindrome(x string) bool {
	for i := 0; i < len(x); i++ {
		j := len(x) - 1 - i
		if x[i] != x[j] {
			return false
		}
	}
	return true
}

func main() {
	a := isPalindrome(121)
	b := isWordPalindrome("aba")
	fmt.Println(a)
	fmt.Println(b)
}
