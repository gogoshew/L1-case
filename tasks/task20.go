package tasks

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func Task20() {
	someString := "snow dog sun"
	// Функция Fields из пакета strings принимает строку, делит ее на слова и создает срез,
	// используя пробелы как элемент разделения
	words := strings.Fields(someString)
	result := "" // переменная для записи результата

	// Цикл в обратном порядке принтит слова из среза слов строки
	for i := len(words) - 1; i >= 0; i-- {
		//fmt.Printf("%s %s ", someString, words[i])
		result += words[i] + " "
	}
	fmt.Printf("Исходная строка - %s, перевернутая строка - %s\n", someString, result)
}
