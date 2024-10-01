package main

import "fmt"

func main() {
	// Массивы. Основа
	// Определение массива
	// При инициализации массива важно передать информацию
	// сколько элементов в нем должно находиться
	var arr [5]int
	fmt.Println("This is my array:", arr)

	// Обращение к элементу массива осуществляется по индексу
	arr[0] = 10
	arr[1] = 20
	arr[3] = -100
	arr[4] = 10000
	fmt.Println("Array after change:", arr)

	// Определение массива с указанием элементов на месте
	// Если при инициализации кол-во элементов меньше номинальной длины
	// то недостающие элементы инициализируются нулями
	newArr := [5]int{10, 20, 30}
	fmt.Println("This is my newArr:", newArr)

	// Создание массива через инициализацию переменных
	arrWithValues := [...]int{10, 20, 30, 40}
	fmt.Println("This is my arrWithValues:", arrWithValues, len(arrWithValues))
	arrWithValues[0] = 10_000
	fmt.Println("This is modified arrWithValues:", arrWithValues, len(arrWithValues))

	// Массив - это набор значений. При всех операциях массив копируется.
	// (жестко, на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 1000
	fmt.Println("First", first)
	fmt.Println("Second", second)

	// Массив и его размер - это две составляющие одного типа
	//var third [2]int
	//third = first - Ошибка. Разные типы данных

	// Итерирование по массиву
	floatArray := [...]float64{12.5, 10.0, 12.1, 15.2, 13.5}
	for i := 0; i < len(floatArray); i++ {
		fmt.Printf("%d element if array is %.2f\n", i, floatArray[i])
	}
	fmt.Println("-----")

	// Итерирование по массиву через оператор range
	var sum float64
	for idx, value := range floatArray {
		fmt.Printf("%d element if array is %.2f\n", idx+1, floatArray[idx])
		sum += value
	}
	fmt.Println("Total sum:", sum)

	// Игнорирование idx при обращении к массиву через оператор range
	var summa float64
	for _, value := range floatArray {
		summa += value
	}
	fmt.Println("Repat total sum:", summa)

	// Многомерные массивы
	words := [2][2]string{
		{"Mark", "Alice"},
		{"Viktor", "Inna"},
	}
	fmt.Println("Multidimensional array:", words)

	// Итерирование по многомерному массиву
	for _, innArr := range words {
		for _, word := range innArr {
			fmt.Printf("%s ", word)
		}
		fmt.Println()
	}
}
