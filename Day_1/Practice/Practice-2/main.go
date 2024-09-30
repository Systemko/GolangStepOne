// Задача №2
// Вход: трехзначное число.
// Найти первую, вторую и последнюю цифры заданного трехзначного числа.

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&number)

	var hundreds = number / 100
	var dozens = (number - hundreds*100) / 10
	var digits = number % 10

	fmt.Println("Первая цифра числа:", hundreds)
	fmt.Println("Вторая цифра числа:", dozens)
	fmt.Println("Третья цифра числа:", digits)
}
