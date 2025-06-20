package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	AvgGrade  float64
}

// Метод для вычисления текущего возраста студента
func (s Student) Age() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

// Метод для получения статуса по среднему баллу
func (s Student) Status() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "Отличник"
	case s.AvgGrade >= 3.5:
		return "Хорошист"
	default:
		return "Троечник" // ну то есть - я
	}
}

func main() {
	student := Student{Name: "Антон", BirthYear: 2005, AvgGrade: 3.2}

	fmt.Printf("Студент: %s\n", student.Name)
	fmt.Printf("Возраст: %d\n", student.Age())
	fmt.Printf("Статус: %s\n", student.Status())
}
