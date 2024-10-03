package main

import "fmt"

// Вложенные структуры - использование одной структуры
// как тип поля другой структуры.

type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

// Встроенные структуры - когда добавляем поля одной структуры
// к другой без имени
type Professor struct {
	firstName  string
	lastName   string
	age        int
	greatWorks string
	University // В этом месте происходит добавление всех полей структуры University
}

func main() {
	// Создание экземпляров структур с вложением

	student := Student{
		firstName: "Fedor",
		lastName:  "Sumkin",
		university: University{
			yearBased: 1991,
			infoShort: "good University",
			infoLong:  "It's very good University",
		},
	}

	// Получение доступа к вложенным полям структуры
	fmt.Println("firstName:", student.firstName)
	fmt.Println("lastName:", student.lastName)
	fmt.Println("year based of University:", student.university.yearBased)

	fmt.Println("-----")

	// Создание экземпляра с встроенной структурой
	prof := Professor{
		firstName:  "Anatoly",
		lastName:   "Smirnov",
		age:        35,
		greatWorks: "Ultimate Go programming",
		University: University{
			yearBased: 1734,
			infoShort: "Short info",
			age:       2024 - 1734,
		},
	}

	// Обращение к состояниям встроенной структуры
	fmt.Println("firstName:", prof.firstName)
	fmt.Println("lastName:", prof.lastName)
	fmt.Println("age:", prof.age)
	fmt.Println("greatWorks:", prof.greatWorks)
	fmt.Println("year based of University:", prof.yearBased)
	fmt.Println("infoLong of University:", prof.infoLong)
	fmt.Println("age of University:", prof.University.age) // Доступ к полю вышележащей структуры

	fmt.Println("-----")

	// Сравнение экземпляров
	// Если хотя бы одно из полей структуры несравнимо, то и вся структура тоже несравнима

	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)

	fmt.Println("-----")
}
