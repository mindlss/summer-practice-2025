package main

import "fmt"

// Функция с передачей параметра по значению
func incrementByValue(a int) {
	a = a + 1
	fmt.Println("Внутри incrementByValue:", a)
}

// Функция с передачей параметра по указателю
func incrementByPointer(a *int) {
	*a = *a + 1
	fmt.Println("Внутри incrementByPointer:", *a)
}

func main() {
	value := 10
	fmt.Println("Исходное значение:", value)

	incrementByValue(value)
	fmt.Println("После incrementByValue:", value)

	incrementByPointer(&value)
	fmt.Println("После incrementByPointer:", value)
}
