/*
3.7 Вводим строку без знаков препинания(5 слов).
    Найти самое длинное слово в строке и вывести сколько в нем букв.

    Пример:
    In: Скажи как дела в учебе 
    Out: учебе, 5

    In: Закрепляем циклы в языке Golang
    Out: Закрепляем, 10
*/

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    var words [5]string
    
    fmt.Print("Введите предложение из 5 слов без знаков препинания: ")

    fmt.Scan(&words[0], &words[1], &words[2], &words[3], &words[4])
    
    var index int = 0;
    
    for i := 1; i < len(words); i++ {
        len1 := utf8.RuneCountInString(words[index])
        len2 := utf8.RuneCountInString(words[i])
        if len2 >= len1 {
            index = i
        }
    }

    fmt.Printf("%s, %d\n", words[index], utf8.RuneCountInString(words[index]))
}