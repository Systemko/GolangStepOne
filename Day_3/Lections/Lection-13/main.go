package main

import (
	"fmt"
)

func main() {
	// Ассоциативные коллекции
	// Map - неупорядоченная коллекция ключ/значение

	// Инициализация пустой коллекции
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	fmt.Println("-----")

	// Добавление элементов
	mapper["Petr"] = 23
	mapper["Elena"] = 12
	fmt.Println("Map after adding two values:", mapper)

	fmt.Println("-----")

	// Инициализация с указанием значений
	newMapper := map[string]int{
		"Alice": 192,
		"Bob":   134,
	}
	newMapper["Nikolay"] = 87
	fmt.Println("newMapper:", newMapper)

	fmt.Println("-----")

	// Что может быть ключом в ассоциативной коллекции?
	// Ключом в мапе может быть только сравнимый тип, т.е. определены операции ==, !=

	fmt.Println("-----")

	// нулевое значение для map является nil
	var mapZeroValue map[string]int // mapZeroValue = nil
	//mapZeroValue["Alice"] = 12 - Будет ошибка, т.к. коллекция не создана
	fmt.Println(mapZeroValue)

	fmt.Println("-----")

	// Получение значений из ассоциативной коллекции
	testPerson := "Alice"
	fmt.Printf("Salary of %s: %d\n", testPerson, newMapper[testPerson])
	testPerson = "Vasya"
	fmt.Printf("Salary of %s: %d\n", testPerson, newMapper[testPerson]) // Ошибки не будет

	fmt.Println("-----")

	// Проверка на наличие ключа
	employee := map[string]int{
		"Denis": 0,
		"Alice": 0,
		"Vasya": 0,
	}

	// При обращении по ключу - возвращаются два значения
	if value, ok := employee["Denis"]; ok {
		fmt.Println("Denis exists in collection and its value:", value)
	} else {
		fmt.Println("Denis does not exist in collection")
	}

	fmt.Println("-----")

	// Перебор значений ассоциативной коллекции
	for key, value := range employee {
		fmt.Printf("key: %s and value: %v\n", key, value)
	}

	fmt.Println("-----")

	// Как удалять пары:
	// Удаление существующей пары
	fmt.Println("Before deleting", employee)
	delete(employee, "Vasya")
	fmt.Println("After deleting", employee)

	// Удаление несуществующей пары
	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna")
	}

	// Количество пар = длина ассоциативной коллекции
	fmt.Println("Pairs amount in map:", len(employee))

	fmt.Println("-----")

	// Ассоциативная коллекция как и слайс - ссылочный тип
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}
	newWords := words
	newWords["Three"] = "Три"
	delete(newWords, "One")
	fmt.Println("words:", words)
	fmt.Println("newWords:", newWords)

	fmt.Println("-----")

	// Сравнение коллекций
	// Сравнение массивов (массив можно использовать как ключ в map)
	if [3]int{1, 2, 3} == [3]int{1, 2, 3} {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
	// Сравнение слайсов. Из-за того, что слайсы ссылочный тип,
	// его можно сравнивать на равенство только с nil
	// if []int{1, 2, 3} == []int{1, 2, 3} { // Ошибка
	// 	fmt.Println("Equal")
	// } else {
	// 	fmt.Println("Not equal")
	// }

	// Сравнение ассоциативных коллекций. Из-за того, что map ссылочный тип,
	// его можно сравнивать на равенство только с nil.
	aMap := map[string]int{
		"a": 1,
	}
	var bMap map[string]int

	if aMap == nil {
		fmt.Println("aMap is zero value map")
	}

	if bMap == nil {
		fmt.Println("bMap is Zero value map")
	}

	// Следствие
	// Если slice или map являются составляющими какой либо структуры
	// То мы можем сказать, что структура автоматически является несравниваемой

}
