package main

import (
	"fmt"
	"unicode/utf8"
)

// Описание интерфейса
type BigWorld interface {
	IsBig() bool
}

// Претендент, который должен иметь метод IsBig
type MySuperString string

func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}

	return false
}

func main() {
	sample := MySuperString("hello")
	var interfaceSample BigWorld
	interfaceSample = sample
	fmt.Println("IsBig?", interfaceSample.IsBig())

	// Попытка присвоить значение с типом, не удовлетворяющему интерфейсу
	// приводит к ошибке.
	//var interfaceBadSample BigWorld
	//interfaceBadSample = "long string" - ошибка, т.к. string не имеет реализации IsBig
}
