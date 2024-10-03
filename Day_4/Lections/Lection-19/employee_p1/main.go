package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

// Методы - функции привязанные к определенным структурам.
// Шаблон
// func (s Struct) MethodName (parameters type) outType {
// body
//}
// Receiver - получатель метода

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}

func main() {
	// Создание экземпляра
	emp := Employee{
		name:     "Mark",
		position: "Senior golang programmer",
		salary:   200_000,
		currency: "RUB",
	}

	// Вызов метода
	emp.DisplayInfo()

	fmt.Println("-----")
}
