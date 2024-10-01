/*
Задача №3

На входе размер вклада(float64), кол-во лет(int64) и процент по вкладу(int64)
Проверить условия (от и до включительно):
вклад от 100 до 1_000_000
кол-во лет от 1 до 100
процент от 1 до 20

и посчитать размер вклада при выполнении условий.
В противном случае вывести сообщение о неправильных данных

Использовать ежегодную капитализацию.
*/

package main
 
import (
	"fmt"
	"math"
)

func main() {
	var (
		amount float64
		years int64
		percentage int64
	)
	fmt.Print("Введите размер вклада: ")
	fmt.Scan(&amount)
	fmt.Print("Введите кол-во лет: ")
	fmt.Scan(&years)
	fmt.Print("Введите процент по вкладу: ")
	fmt.Scan(&percentage)

	if amount >= 100 && amount <= 1_000_000 && years >= 1 && years <= 100 && percentage >= 1 && percentage <= 20 {
		var profit float64 
		profit = amount * math.Pow(1 + (float64(percentage) / 100.00), float64(years))

		fmt.Println("Размер вклада:", profit)
	} else {
		fmt.Println("Неверные входные данные")
	}
}
