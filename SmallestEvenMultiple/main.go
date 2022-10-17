package main

func main() {
}

func smallestEvenMultiple(n int) int {
	var z int
	if 2 > n {
		z = 2
	}
	z = n

	for {
		if (z%2 == 0) && (z%n == 0) {
			break
		}
		z += 1
	}
	return z
}
