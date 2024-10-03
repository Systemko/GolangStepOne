package main

import "fmt"

// Структура - заименованный набор полей (состояний),
// определяющий новый тип данных

// Определение структуры (явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}

// Если имеется ряд состояний одинакового типа,
// то можно сделать так:
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

// Структура  с анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func main() {
	// Создание экземпляра структуры
	student := Student{
		firstName: "Fedor",
		lastName:  "Sumkin",
		age:       21,
	}
	fmt.Println(student)

	fmt.Println("-----")

	// Передача экземпляра структуры в функцию
	printStudent(student)

	fmt.Println("-----")

	// Порядок указания свойств - такой же как и в структуре
	student2 := Student{"Sam", "Gonjy", 19}
	printStudent(student2)

	fmt.Println("-----")

	// Что если не все поля структуры определены?
	student3 := Student{
		firstName: "Merin",
	}
	printStudent(student3)

	fmt.Println("-----")

	// Анонимные структуры
	anonStudent := struct {
		age           int
		groupID       int
		professorName string
	}{
		age:           2300,
		groupID:       2,
		professorName: "Gandalf",
	}
	fmt.Println("anonStudent:", anonStudent)

	fmt.Println("-----")

	// Доступ к полям (состояниям) и их модификация
	studMark := Student{"Mark", "Ivanov", 19}
	printStudent(studMark)
	studMark.age += 2
	printStudent(studMark)

	fmt.Println("-----")

	// Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	printStudent(emptyStudent1)
	printStudent(emptyStudent2)

	fmt.Println("-----")

	// Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       23,
	}
	fmt.Println("Value of studPointer", studPointer)

	seconfPointer := studPointer
	(*seconfPointer).age += 20
	fmt.Println("Value of studPointer after changing", studPointer)

	studPointerNew := new(Student)
	fmt.Println("Value of studPointerNew", studPointerNew)

	fmt.Println("-----")

	// Работа с доступом к полям через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age) // явное разименование
	fmt.Println("Age via .age:", studPointer.age)          // неявное разименование

	fmt.Println("-----")

	// Создание экземпляра с анонимными полями
	human := &Human{
		firstName: "Bob",
		lastName:  "Northon",
		string:    "Additional info",
		int:       1,
		bool:      true,
	}
	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)

	fmt.Println("-----")

	// Указатели на массивы
	arr := [3]int{1, 2, 3}
	fmt.Println("arr before mutation", arr)
	mutation(&arr)
	fmt.Println("arr after mutation", arr)

	// Но лучще использовать Slice
	newArr := [3]int{1, 2, 3}
	fmt.Println("newArr before mutation", newArr)
	mutationSlice(newArr[:])
	fmt.Println("newArr after mutation", newArr)

	fmt.Println("-----")
}

func printStudent(student Student) {
	fmt.Println("FirstName:", student.firstName)
	fmt.Println("LastName:", student.lastName)
	fmt.Println("Age:", student.age)
}

func mutation(arr *[3]int) {
	// Можно так
	// (*arr)[1] = 900
	// (*arr)[2] = 1000
	// Но можно и так, т.к. Go сам разыменует указатель на массив
	arr[1] = 900
	arr[2] = 1000
}

func mutationSlice(sls []int) {
	sls[1] = 900
	sls[2] = 1000
}
