package tasks

import (
	"fmt"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func Task19() {
	someString := "главрыба"
	//printReverseByDefer(someString)
	//printReverseString(someString)
	printReversed(someString)
}

// printReverseByDefer принтит строку по-символьно в обратном порядке из-за особенности вызова из defer стэка
func printReverseByDefer(str string) {
	for _, char := range str {
		defer fmt.Printf("%s", string(char))
	}
}

// printReverseString превращает строку в руны - символы в unicode, в цикле символы строки меняются местами пока i < j
func printReverseString(str string) {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(str, string(runes))
}

//
func printReversed(str string) {
	// Создаем срез рун по длине строки
	l := len(str)
	runes := make([]rune, l)

	// Далее синтаксический сахар Го, range возвращает позицию символа в строке и руну
	for _, char := range str {
		l--
		runes[l] = char
	}
	fmt.Println(str, string(runes))
}
