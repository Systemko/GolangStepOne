/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
    var number int16
    fmt.Print("Введите четырехзначное число: ")
    fmt.Scan(&number)
    
    if number > 999 && number < 10000 {
        var thousands = number / 1000
        var hundreds = number / 100 % 10
        var dozens = number / 10 % 10
        var digits = number % 10
        
        if thousands == digits && hundreds == dozens {
            fmt.Println("Выход:", number, "- палиндром")
        } else {
            fmt.Println("Выход:", number, "- не палиндром")
        }
    } else {
        fmt.Println("Неверные входные данные")
    }
}