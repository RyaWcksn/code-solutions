package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 3, 5, 6, 7, 8, 4, 2, 5, 6}
	target := 15
	solution := TwoSum(arr, target)
	fmt.Println(solution)
}

func TwoSum(list []int, target int) []int {
	/* for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == target {
				return []int{list[i], list[j]}
			}
		}
	}
	return []int{} */

	m := make(map[int]int)
	for idx, num := range list {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx
	}
	return nil
}
