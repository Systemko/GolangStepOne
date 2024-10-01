package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Числовые типы
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a and b:", a+b)

	// Тип данных можно получить через форматирование %T
	// int - платформо-зависимый
	fmt.Printf("type of a: %T and type of b: %T\n", a, b)

	// Узнаем сколько байт в памяти занимает тип int
	fmt.Printf("Type %T sizes %d bytes in memory\n", a, unsafe.Sizeof(a))
	fmt.Printf("Type %T sizes %d bytes in memory\n", b, unsafe.Sizeof(b))

	// Если выполняются арифметические операции над int и intX
	// то обязательно использовать операции приведения типов,
	// т.к. int != int64
	var c64 int64 = 16_213_897
	var usualInt int = 123_567
	fmt.Println(c64 + int64(usualInt))

	// Аналогичным образом устроены uint8, uint16, uint32, uint64

	// Numercis. Float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("Type %T sizes %d bytes in memory\n", floatFirst, unsafe.Sizeof(floatFirst))
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Println("Sum:", sum, "Sub:", sub)
	fmt.Println(math.Pow(floatFirst, floatSecond))

	// Numerics. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// Strings. Строка - набор байт.
	name := "Федя"
	lastname := "Pupkin"
	fi := name + " " + lastname
	fmt.Println("Full name:", fi)
	fmt.Println("Lenght of fi in bytes:", len(fi))
	fmt.Println("Lenght of name in bytes:", len(name))
	fmt.Println("Lenght of lastname in bytes:", len(lastname))

	// Rune - руна. Это один utf символ
	fmt.Println("Lenght of name in symbols:", utf8.RuneCountInString(name))

	// Поиск подстроки в строке
	totalString, subString := "ABCDEF", "cde"
	fmt.Println(strings.Contains(totalString, subString))

	// rune -> это alias int32
	var sampleRune rune
	var anotherRune = 'Q' // Для инициализации руны используйте ''
	var thirdRune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char: %c\n", sampleRune)
	fmt.Printf("Rune as char: %c\n", anotherRune)
	fmt.Printf("Rune as char: %c\n", thirdRune)

	// Сравнение
	// -1 если первый меньше второго, 0 равны, 1 если первый больше второго
	fmt.Println(strings.Compare("abcd", string('a')))

	// Байты
	var aByte byte // alias int8
	fmt.Println("Byte:", aByte)
}
