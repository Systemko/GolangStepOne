package main

import "fmt"

func main() {
	// Анонимная функция
	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println("anon(10, 20):", anon(10, 20))

	fmt.Println("-----")

	// Анонимная функция внутри явной
	fmt.Println("bigFunction(30, 40):", bigFunction(30, 40))

	fmt.Println("-----")

	// Явные функции в Go (в принципе луюбая функция) - это экземпляр 1-го уровня
	// поэтому ее можно присваивать переменной, передавать в качестве параметра,
	// и возвращать в качестве результата
	var command = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("Result with command", command, "is", res(10, 20))

	fmt.Println("-----")

	// Тип функции
	fmt.Printf("Тип функции res: %T\n", res)
	fmt.Printf("Тип функции calcAndReturnValidFunc: %T\n", calcAndReturnValidFunc)

	fmt.Println("-----")

	// Функция как аргумент функции
	fmt.Println("receiveFuncAndReturnValue(add):", receiveFuncAndReturnValue(add))
	fmt.Println("receiveFuncAndReturnValue(res):", receiveFuncAndReturnValue(res))
	fmt.Println(
		"receiveFuncAndReturnValue(random):", 
		receiveFuncAndReturnValue(func(a, b int) int { return a*a + b*b + 2*a*b }))

	fmt.Println("-----")

}

// Анонимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int {
		return a + b + 100
	}(aArg, bArg)
}

// Функция возвращающая функцию
func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int {
			return a + b
		}
	} else if command == "substraction" {
		return func(a, b int) int {
			return a - b
		}
	} else if command == "multiplication" {
		return func(a, b int) int {
			return a * b
		}
	} else if command == "division" {
		return func(a, b int) int {
			return a / b
		}
	}
	return nil
}

// Функция как аргумент функции
func receiveFuncAndReturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200

	return f(intVarA, intVarB)
}

func add(a, b int) int{
	return a + b
}