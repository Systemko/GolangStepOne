/*
3.5 В бесконечном цикле приложение запрашивает у пользователя числа.
    Ввод завершается, как только пользователь вводит число "-1".
    После завершения ввода приложение выводит сумму чисел.
    Используем конструкцию:
    for {
      // body
    }
*/

package main

import "fmt"

const stopWord int = -1

func main() {
    var number int
    var sum int
    
    fmt.Println("Для получения суммы чисел вводите числа каждое с новой строки.")
    fmt.Printf("Для завершения ввода и получения результата введите %d\n", stopWord)
    
    for {
        fmt.Scanf("%d", &number)
        
        if (number != -1) {
            sum += number
        } else {
            break
        }
    }
    
    fmt.Printf("Сумма чисел: %d", sum)
}