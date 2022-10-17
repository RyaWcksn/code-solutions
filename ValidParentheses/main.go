package main

import "fmt"

func main() {
	data := "}}{}[]"
	isValid(data)
}

// {}
// [}
// []{}
// {{[]

func isValid(s string) bool {
	stack := []string{}

	dict := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	for _, v := range s {
		if parentheses, ok := dict[string(v)]; ok {
			stack = append(stack, parentheses)
			continue
		}

		l := len(stack) - 1
		fmt.Println(l)
		fmt.Println(stack)
		if l < 0 || string(v) != stack[l] {
			return false
		}
		fmt.Println(stack)
		stack = stack[:l]
		fmt.Println(stack)
	}

	return len(stack) == 0
}
