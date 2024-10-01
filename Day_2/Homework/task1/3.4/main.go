/*
3.4 В цикле получать от пользователя оценки по четырём экзаменам.
    Вывести сумму набранных им баллов.
    Функцию fmt.Scan() в коде используем только один раз.
*/

package main

import "fmt"

func main() {
    const examCount int = 4
    
    var score int
    var totalScore int
    
    for i := 1; i <= examCount; i++ {
        fmt.Printf("Введите кол-во баллов по %d-му экзамену: \n", i)
        fmt.Scan(&score)
        totalScore += score
    }
    
    fmt.Printf("Общий балл по всем экзаменам: %d", totalScore)
}