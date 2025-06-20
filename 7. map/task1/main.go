package main

import (
	"fmt"
)

// Функция для добавления оценки
func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

// Функция для поиска оценки по имени
func getGrade(grades map[string]int, name string) {
	grade, exists := grades[name]
	if exists {
		fmt.Printf("Оценка для %s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

// Функция для удаления студента
func removeStudent(grades map[string]int, name string) {
	delete(grades, name)
}

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Анна", 90)
	addGrade(grades, "Борис", 85)

	getGrade(grades, "Анна")
	getGrade(grades, "Игорь") // нет такого

	removeStudent(grades, "Борис")

	getGrade(grades, "Борис") // уже удалён
}
