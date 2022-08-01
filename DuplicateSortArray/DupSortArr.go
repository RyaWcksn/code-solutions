package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	solution := RemoveDuplicates(arr)
	fmt.Println(solution)
}

func RemoveDuplicates(nums []int) int {
	a := 0
	if len(nums) == 0 {
		return a
	}
	for b := 1; b < len(nums); b++ {
		if nums[a] != nums[b] {
			a++
			nums[a] = nums[b]
			fmt.Println(nums[a])
		}
	}
	return a + 1
}
