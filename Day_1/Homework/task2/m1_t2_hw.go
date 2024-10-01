/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример: 
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {
    var number int
    fmt.Print("Введите трехзначное число: ")
    fmt.Scan(&number)
    
    if number > 99 && number < 1000 {
        var hundreds = number / 100
        var dozens = number / 10 % 10
        var digits = number % 10
        
        var reversedNumber = digits * 100 + dozens * 10 + hundreds
        
        fmt.Printf("Реверсная запись: %03d\n", reversedNumber)
    } else {
        fmt.Println("Неверные входные данные")
    }
}