package main

import "fmt"

// Интерфейс "Транспорт"
type Transport interface {
	Move()
	Stop()
}

// Машина
type Car struct {
	Brand string
}

func (c Car) Move() {
	fmt.Println(c.Brand + " едет")
}

func (c Car) Stop() {
	fmt.Println(c.Brand + " остановилась")
}

// Велосипед
type Bike struct {
	Brand string
}

func (b Bike) Move() {
	fmt.Println(b.Brand + " едет")
}

func (b Bike) Stop() {
	fmt.Println(b.Brand + " остановился")
}

func operate(t Transport) {
	t.Move()
	t.Stop()
}

func main() {
	car := Car{Brand: "Toyota"}
	bike := Bike{Brand: "Giant"}

	operate(car)
	operate(bike)
}
