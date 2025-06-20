package main

import (
	"fmt"
	"math"
)

// Интерфейс "Форма"
type Shape interface {
	Area() float64
}

// Структура прямоугольника
type Rectangle struct {
	Width, Height float64
}

// Метод вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура круга
type Circle struct {
	Radius float64
}

// Метод вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 4}

	printArea(rect)
	printArea(circ)
}
