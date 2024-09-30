package main

import (
	"fmt"
	"strings"
)

func main() {
	// Значение по умолчанию для булевого типа
	var firstBoolean bool
	fmt.Println(firstBoolean)

	// Булевские операторы
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	// Классический условный оператор
	if 1 == 1 {
		// body
	}

	// Классический условный оператор с блоком else
	if 1 == 1 {
		// body 1
	} else {
		// body 2
	}

	fmt.Print("Print number: ")
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Number", value, "is even")
	} else {
		fmt.Println("Number", value, "is odd")
	}

	// Классический условный оператор с блоком if else
	if 1 == 1 {
		// body 1
	} else if 2 == 2 {
		// body 2
	} else if 3 == 3 {
		// body 3
	} else {
		// body 4
	}

	fmt.Print("Enter color name: ")
	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	// Инициализация в блоке условного оператора
	// Блок присваивания - только :=
	// Инициализируемая переменная видна ТОЛЬКО внутри области
	// видимости условного оператора
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}
	// fmt.Println(num) - переменная недоступна

	// Приминение else не идиоматично в golang
	// Не идеоматично
	if width := 100; width > 100 {
		fmt.Println("width > 100")
	} else {
		fmt.Println("width <= 100")
	}

	// Идеоматично
	if width := 100; width > 100 {
		fmt.Println("width > 100")
		return
	}
	fmt.Println("width <= 100")
}
