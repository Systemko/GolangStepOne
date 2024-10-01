package main

import (
	"fmt"
	"strings"
)

func main() {
	// Шаблон цикла for как for
	// for init; condition; post {
	// body
	// }
	// где,
	// init - блок инициализации переменных
	// condition - блок условия (если условие верно - тело цикла выполняется)
	// post - блок изменений
	// body - блок тела цикла

	for i := 0; i <= 5; i++ {
		fmt.Println("Current value of i:", i)
	}

	// Важный момент В качестве init можно использовать
	// выражение только через := при создании переменной

	// break - команда, которая прерывает текущее выполнение цикла и передает
	// управление следующим за циклом инструкциям.
	for i := 11; i <= 22; i++ {
		if i > 20 {
			break
		}
		fmt.Println("Current value of i:", i)
	}
	fmt.Println("After break")

	// continue - команда, которая передает управление следующей итерации цикла
	for i := 100; i <= 120; i++ {
		if i > 110 && i <= 115 {
			continue
		}
		fmt.Println("Current value of i:", i)
	}

	// Вложенные циклы и лейблы
	fmt.Println("Треугольник:")
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Бывают ситуации, когда нужно прервать сразу оба цикла
	// Здесь помогут лейблы. Лейблы - это синтаксический сахар
	// т.к. break outer, то цикл не будет выполняться.
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j:%d\n", i, j, i+j)
			if i == j {
				break outer // Остановка циклов
			}
		}
	}
	fmt.Println("After outer stop")

	// Шаблон цикла for как while (модификация for)
	// Классический цикл while do
	var loopVar int = 0

	for loopVar < 10 {
		fmt.Println("Current value of loopVar:", loopVar)
		loopVar++
	}

	// Классический бесконечный цикл
	var password string
outer2:
	for {
		fmt.Println("Insert password:")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again.")
		} else {
			fmt.Println("Password accepted.")
			break outer2
		}
	}

	// Цикл с множественными переменными
	for x, y := 0, 1; x < 5 && y < 6; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}
}
