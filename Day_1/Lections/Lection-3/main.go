package main

import "fmt"

func main() {
	// Деление
	var result float64
	var d, e float64 = 10, 4
	var check float64 = 20 / 3

	a := 42
	b := 20
	c := a / b
	result = d / e

	fmt.Printf("type of c %T and value of c %v\n", c, c)
	fmt.Printf("type of result %T and value of result %v\n", result, result)
	fmt.Printf("type of check %T and value of check %v\n", check, check)

	// Инкремент, декремент
	fmt.Print("Инкремент: ")
	check++
	fmt.Println("check =", check)
	fmt.Print("Декремент: ")
	c--
	fmt.Println("c =", c)

	// Операция взятия остатка
	newCheck := int(check) % 3
	fmt.Println("Остаток от деления check на 3 =", newCheck)
}
