package main

import (
	"fmt"
)

func longestString(strings ...string) string {
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result := longestString("cat", "elephant", "dog", "hippopotamus", "mouse")
	fmt.Println("Самая длинная строка:", result)
}
