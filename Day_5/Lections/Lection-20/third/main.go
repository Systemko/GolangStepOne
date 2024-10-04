package main

import "fmt"

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (st Student) Describe() {
	fmt.Printf("%s and %d years old\n", st.name, st.age)
}

func (st Student) GetName() {
	fmt.Printf(st.name)
}

func TypeFinder(i interface{}) {
	// Присваиваем переменной v значение лежащее под предполагаемым интерфейсом
	switch v := i.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("This is int")
	case Describer:
		fmt.Println("I'm Describer")
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func FindStudent(i interface{}) {
	// Присваиваем переменной v значение лежащее под предполагаемым интерфейсом
	switch v := i.(type) {
	case Student:
		fmt.Println("I'm student")
		v.GetName()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	std := Student{"Alex", 23}
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(std)
	TypeFinder(nil)

	var desc Describer
	desc = std
	FindStudent(desc)
	FindStudent("Hello")
}
