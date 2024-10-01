package main

import "fmt"

func main() {
	// Оператор множественного выбора
	// В switch-case запрещены дублирующиеся условия в case-ах
	var price int
	fmt.Scan(&price)
	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	default:
		fmt.Println("Default case")
	}

	// Case с множеством вариантов
	var ageGroup string = "u" // Возрастные группы "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("First group")
	case "d", "e":
		fmt.Println("Second group")
	default:
		fmt.Println("Default group")
	}

	// Case с выражениями
	var age int
	fmt.Scan(&age)
	switch {
	case age <= 18:
		fmt.Println("Too young")
	case age > 18 && age <= 40:
		fmt.Println("Middle")
	case age > 40 && age <= 100:
		fmt.Println("Too old")
	default:
		fmt.Println("Who are you?")
	}

	// Case с проваливаниями. Проваливания выполняют даже Ложные кейсы
	// В момент выполнения fallthrough у следующего кейсы не проверяется условие,
	// а сразу выполняется тело.
	var number int
	fmt.Scan(&number)
outer:
	switch {
	case number < 100:
		fmt.Printf("%d less than 100\n", number)
		if number%2 == 0 {
			break outer
		}
		fallthrough
	case number > 200:
		fmt.Printf("%d greater than 200\n", number)
		fallthrough
	case number > 1000:
		fmt.Printf("%d greater than 1000\n", number)
		fallthrough
	default:
		fmt.Print("Default")
	}

	// Нехорошая конструкция
	var i int
outloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even", i)
			break outloop
		}
	}
	fmt.Println("END")
}
