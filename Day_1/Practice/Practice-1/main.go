// Задача №1
// Программа запрашивает имя пользователя и возраст
// Нужно вывести имя и возраст за вычетом одного года

package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)
	fmt.Printf("Вас зовут %s но Вам %d.\n", name, age-1)
}
