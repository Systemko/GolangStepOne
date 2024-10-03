package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

// 1. Метод в котором получатель копируется,
// и в его теле происходит работа с локальной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

// Метод в которм получатель передается по ссылке,
// и в теле метода работаем с ссылкой на экземпляр
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

// Используйте методы с PointerReceiver-ом в ситуациях когда:
// - Изменения в работе метода над экземпляром должны быть видны на
//   вызывающей стороне
// - Когда экземпляр достаточно большой, и копировать его 'дорого'
//   с точки зрения производительности
// - С ними может работать любой вид экземпляра

func main() {
	e := Employee{"Alex", 2000}
	fmt.Println("Before setting new name:", e)
	e.SetName("Yan")
	fmt.Println("After setting new name:", e)

	fmt.Println("-----")

	fmt.Println("Before setting new salary:", e)
	e.SetSalary(4000)
	fmt.Println("After setting new salary:", e)

	fmt.Println("-----")
}
