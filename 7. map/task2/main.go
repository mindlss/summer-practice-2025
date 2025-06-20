package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	wordFreq := make(map[string]int)

	// Приводим всё к нижнему регистру и разбиваем по пробелам
	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

func main() {
	text := "Go is fun and go is fast and fun"

	wordFreq := countWords(text)

	fmt.Println("Частота слов:")
	for word, count := range wordFreq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
