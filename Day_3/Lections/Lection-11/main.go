package main

import "fmt"

func main() {
	//
	var array [3]int

	array[0] = 100
	array[0] = 200
	array[0] = 300
	fmt.Printf("Array's values: %v, array as is: %#v\n", array, array)

	// При создании массива необходимо определить размер на этапе компиляции
	// a := 6
	// var array[a] int - так нельзя
	// const b = 6
	// var array[6] int - так можно

	fmt.Println("-----")

	// Слайс - это динамическая обвязка над массивом
	startArray := [4]int{10, 20, 30, 40}
	fmt.Println("startArray:", startArray)
	var startSlice []int = startArray[0:2] // Слайс инициализируется пустыми кадратными скобками
	fmt.Println("slice[0:2] of startArray:", startSlice)

	fmt.Println("-----")

	// Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 25, 30, 70}
	fmt.Println("secondSlice is:", secondSlice)

	fmt.Println("-----")

	// Изменение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	fmt.Println("originArr:", originArr)
	thirdSlice := originArr[1:4] // Это набор ссылок на элементы нижележащего массива
	for i := range thirdSlice {
		thirdSlice[i]++
	}
	fmt.Println("changed originArr:", originArr)
	fmt.Println("thirdSlice:", thirdSlice)

	fmt.Println("-----")

	// Один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]
	fmt.Println("Before changing. Arr:", originArr, "fSlice", fSlice, "sSlice", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After changing. Arr:", originArr, "fSlice", fSlice, "sSlice", sSlice)

	fmt.Println("-----")

	// Срез как встроенный тип
	// type slice struct {
	// Length int
	// Capacity int
	// ZeroElement *byte
	// }

	// Длина и ёмкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice), "capacity:", cap(wordsSlice))
	// Добавляем новый элемент
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice), "capacity:", cap(wordsSlice))
	// Capacity или емкость - это значение показывающее сколько элементов может вместить слайс
	// не выделяя дополнительной памяти
	// Если нет места под новый элемент, то выделяется дополнительная память равная n, где n - текущий размер до изменения
	// по общей формуле: n * 2.
	numerics := []int{1, 2}
	fmt.Println("numerics:", numerics)
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("length:", len(numerics), "capacity:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	fmt.Println("-----")

	// Важно: после выделения памяти под новый массив,
	// ссылки со старого массива будут перенесены в новый
	numArr := []int{1, 2}
	numSlice := numArr[:]
	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 0
	fmt.Println("numArr", numArr)
	fmt.Println("numSlice", numSlice)

	fmt.Println("-----")

	// Как создавать слайсы наиболее эффективно
	// make() - это функция, позволяющаю более детально создавать срезы
	// Сначала инициализируется массив arr = [15]int
	// Затем по нему делается срезу arr[0:10]
	// После этого он возвращается
	sl := make([]int, 10, 15)
	fmt.Println("sl is:", sl, "lenght:", len(sl), "capacity:", cap(sl))

	fmt.Println("-----")

	// Добавление нескольких элементов в срез
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherWords := []string{"four", "five"}
	myWords = append(myWords, anotherWords...)
	myWords = append(myWords, "six", "seven")
	fmt.Println("myWords:", myWords)

	fmt.Println("-----")

	// Многомерный слайс
	mySlice := [][]int{
		{1, 2},
		{2, 3, 4, 5},
		{10, 20, 30},
		{},
	}
	fmt.Printf("mySlice: %v\nmySlice as is: %#v\n", mySlice, mySlice)

	fmt.Println("-----")

	// Слайсы рун

	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRunes := string(runeHexSlice)
	fmt.Println("Runes:", runeHexSlice, "myStrFromRunes:", myStrFromRunes)

	fmt.Println("-----")

	// Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 't'}
	myStrFromLiteralRunes := string(runeLiteralSlice)
	fmt.Println("Runes:", runeLiteralSlice, "myStrFromLiteralRunes:", myStrFromLiteralRunes)

	fmt.Println("-----")

	// Вычисление емкости строки - бессмысленно, т.к. строка - базовый тип
	fmt.Println(cap([]rune("Вася")))

	fmt.Println("-----")

	// Строки неизменяемые (immutable), а слайсы рун - изменяемые (mutable)
	firstName := "John"
	fmt.Println("firstName", firstName)
	// firstName[0] = 'B' - Нельзя изменить

	firstSliceName := []rune("John")
	firstSliceName[0] = 'B'
	fmt.Println("firstSliceName", string(firstSliceName))

	fmt.Println("-----")

	// Синтаксический сахар (можно использовать десятичное представление байтов)
	myDecimalByteSlice := []byte{100, 101, 102, 103}
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println("myDecimalStr", myDecimalStr)

	fmt.Println("-----")
}
