package main

import (
	"fmt"
	"strings"
)

func main() {
	data := TruncateString("who are you fellas iam arya", 4)
	fmt.Println(data)
}

func TruncateString(s string, k int) string {
	var value string
	var valueArr []string
	words := strings.Fields(s)
	for i := 0; i < k; i++ {
		valueArr = append(valueArr, words[i])
		value = strings.Join(valueArr, " ")
	}

	return value
}
