/*
3.1 Пользователь вводит числа a и b (b больше a).
    Вывести все целые числа, расположенные между ними.
*/

package main

import "fmt"

func main() {
    var (
        number1 int16
        number2 int16
    )
    fmt.Println("Введите два целых числа, где первое меньше второго: ")
    fmt.Scan(&number1, &number2)
    
    if number1 < number2 {
        fmt.Println("Целые числа между указанными: ")
        for i := number1 + 1; i < number2; i++ {
            fmt.Printf("%d ", i)
        }
    } else {
        fmt.Println("Неверные входные данные")
    }
}