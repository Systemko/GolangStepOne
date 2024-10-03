package main

import "fmt"

func main() {
	// Указатели - переменная, хранящая в качестве значения
	// адрес в памяти другой переменной.

	// Определение указателя на переменную
	var variable int = 30
	var pointer *int = &variable // & - операция взятия адреса в памяти
	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Value of pointer %v\n", pointer)

	fmt.Println("-----")

	// Какое значение по умолчанию для указателя
	var zeroPointer *int // zeroPointer имеет значение nil
	fmt.Printf("Value of zeroPointer %v\n", zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
	}
	fmt.Printf("Value of zeroPointer after init %v\n", zeroPointer)

	fmt.Println("-----")

	// Разыменовывание указателя - получение значения по адресу
	// *pointer - возвращает значение, хранимое по адресу
	var numericValue int = 32
	pointerToNumeric := &numericValue
	fmt.Printf("Value of pointerToNumeric: %v and it's address %v\n", *pointerToNumeric, pointerToNumeric)

	fmt.Println("-----")

	// Создание указателей на нулевые значения типов
	var zeroVar int
	var zeroPoint *int = &zeroVar
	fmt.Println("zeroPoint:", zeroPoint)
	zeroNewPoint := new(int) // Создает zeroNewPoint и возвращает адрес, где 0 хранится. Лучше делать так.
	fmt.Println("zeroNewPoint:", zeroNewPoint)

	fmt.Println("-----")

	// Изменение значениея хранимого по адресу через указатель
	zeroPtoInt := new(int)
	fmt.Printf("Value of zeroPtoInt: %v and it's address %v\n", *zeroPtoInt, zeroPtoInt)
	*zeroPtoInt += 40
	fmt.Printf("Value of zeroPtoInt: %v and it's address %v\n", *zeroPtoInt, zeroPtoInt)
	b := 345
	a := &b
	c := &b
	*a++
	*c += 100
	fmt.Println("b:", b)

	fmt.Println("-----")

	// Указательная арифметика полностью отсутствует

	fmt.Println("-----")

	// Передача указателей в функции
	// Большой прирост производительности за счет того, что передается не значение
	// которое должно копироваться, а лишь адрес в памяти

	sample := 1
	fmt.Println("origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("changed value of sample:", sample)

	fmt.Println("-----")

	// Возврат указателя из функции
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Type of ptr1: %T: value of ptr1: %v and address of ptr1 %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Type of ptr2: %T: value of ptr2: %v and address of ptr2 %v\n", ptr2, ptr2, *ptr2)

	fmt.Println("-----")
}

// Функция которая принимает указатель
func changeParam(value *int) {
	*value += 100
}

// Функция которая возвращает указатель
func returnPointer() *int {
	var numeric int = 321
	return &numeric // В момент получения указателя Go перемещает данную переменную в кучу
}
