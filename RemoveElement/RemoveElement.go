package main

func main() {
	arr := []int{3, 2, 2, 3}
	_ = RemoveElement(arr, 3)
}

func RemoveElement(nums []int, val int) int {
	a := 0
	for k, v := range nums {
		if nums[k] != val {
			nums[a] = v
			a++
		}
	}

	return a
}
