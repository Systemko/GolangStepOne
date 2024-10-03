package main

import "fmt"

type University struct {
	city string
	name string
}

// Данный метод связан только со структурой Universirty

func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

// В структуру Professor встроены поля структуры University
// (все методы тоже переходят)
type Professor struct {
	fullname string
	age      int
	University
}

func (p Professor) Info() {
	fmt.Println("name:", p.fullname)
	fmt.Println("age:", p.age)
}

func main() {
	prof := Professor{
		fullname: "Boris Bobroff",
		age:      42,
		University: University{
			city: "Moscow",
			name: "BMSTU",
		},
	}

	// Вызываем метод структуры University
	// через экземпляр профессора
	prof.FullInfoUniversity()

	// Вызываем свой метод
	prof.Info()
}
