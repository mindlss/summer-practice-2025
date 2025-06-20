package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"

	// Подсчёт символов
	length := len(text)
	fmt.Println("Длина строки:", length)

	// Поиск подстроки
	found := strings.Contains(text, "World")
	fmt.Println("Содержит 'World':", found)

	// Изменение регистра
	fmt.Println("В верхнем регистре:", strings.ToUpper(text))
	fmt.Println("В нижнем регистре:", strings.ToLower(text))
}
