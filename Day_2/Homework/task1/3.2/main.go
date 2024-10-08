/*
3.2 Доработать предыдущую задачу так, чтобы выводились только числа,
    делящиеся на 5 без остатка.
*/

package main

import "fmt"

const divider int16 = 5

func main() {
    var (
        number1 int16
        number2 int16
    )
    fmt.Println("Введите два целых числа, где первое меньше второго: ")
    fmt.Scan(&number1, &number2)
    
    if number1 < number2 {
        fmt.Printf("Целые числа между указанными, которые делятся на %d:\n", divider)
        for i := number1 + 1; i < number2; i++ {
            if i % divider == 0 {
                fmt.Printf("%d ", i)
            }
        }
    } else {
        fmt.Println("Неверные входные данные")
    }
}