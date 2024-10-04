package main

import (
	"GolangStepOne/Day_5/Lections/Lection-22/rectangle"
	"fmt"
)

// Разделяющая область видимости одного пакета
// В этом примере 2 модуля main и calculator помещены в одну директорию

// Все что находится внутри одного пакета, доступно из любого модуля без импортирования

func main() {
	// Данные функции видны компилятору, потому что они в рамках пакета main
	resAdd := add(10, 20)
	fmt.Println("resAdd", resAdd)

	resSub := substract(50, 12)
	fmt.Println("resSub", resSub)

	resMul := multiply(3, 9)
	fmt.Println("resMul", resMul)

	resDiv := divide(29, 4)
	fmt.Println("resSub", resDiv)

	// Если имя сущности написано с маленькой буквы, то ее нельзя передать за пределы пакета
	// Если имя сущности написано с большой буквы, то ее можно экспортировать за пределы пакеты
	// Как запустить:
	// go run main.go calculator.go
	// go build main.go calculator.go
	// go install main.go calculator.go
	// go run .

	//
	rec := rectangle.New(10, 20, "Red")
	fmt.Printf("%T\n", rec)
	fmt.Println("Perimeter", rec.Perimeter())

	// Если go.mod создан, то облегчается сборка
	// go build видит отправную точку
	// Как правило, при инициализации проекта, используют go.mod
}
