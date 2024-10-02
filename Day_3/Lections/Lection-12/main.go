package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var name string

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		// Команда захвата потока ввода и сохранения в буфер
		// (захват идет до символа окончания строки)
		name = input.Text()
		// Метод Text() возвращает элементы, помещенные в буфер в виде строки (string)
	}
	fmt.Println(name)

	fmt.Println("-----")

	// Вариант с циклом
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is", result)
		}
	}

	fmt.Println("-----")

	// Преобразование строкового литерала к чему-нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr) // Atoi -> Anyting to Int
	fmt.Printf("number: %d and type is: %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	// Но это 64-разрядное числов будет ГАРАНТИРОВАННО без проблем приводится к float32
	numFloat32, _ := strconv.ParseFloat(numStr, 32)
	fmt.Println(numInt64, numFloat32)
	fmt.Printf("%.3f and %T\n", numFloat32, float32(numFloat32))

	fmt.Println("-----")

	// Пример с bufio.NewReader
	fmt.Println("bufio.NewReader")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("line as is: %#v\n", line)
	data := strings.Fields(line)
	fmt.Printf("data value: %v and as is: %#v\n", data, data)
}
