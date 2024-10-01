/*
3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)
*/

package main

import "fmt"

func main() {
    var number int8
    fmt.Print("Введите число от 1 до 10: ")
    fmt.Scan(&number)
    
    if number > 0 && number <= 10 {
        fmt.Printf("Таблица умножения на число %d:\n", number)
        for i := 1; i <= 10; i++ {
            fmt.Printf("%d * %d = %d\n", i, number, int8(i) * number)
        }
    } else {
        fmt.Println("Неверные входные данные")
    }
}