/*
3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.
Решить с помощью индексов, т.е. работаем с числом, как со строкой.
*/

package main

import (
    "fmt"
    //"strings"
    "strconv"
)

func main() {
    var 
    (
        number int64
        digits string
        sum int64
    )
    
    fmt.Print("Введите натуральное число: ")
    fmt.Scanf("%d", &number)
    
    stringNumber := fmt.Sprintf("%d", number)
    
    for i := 0; i < len(stringNumber); i++ {
        symbol := string(stringNumber[i])
        digit, _ := strconv.ParseInt(symbol, 10, 64)
        digits += fmt.Sprintf("%d + ", digit)
        sum += digit
    }

    digits, _ = strings.CutPrefix(digits, " + ")

    fmt.Println(digits, "=", sum)
}