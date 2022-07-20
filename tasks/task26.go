package tasks

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func Task26() {
	// Создаем слайс строк
	coupleStrings := []string{"abcd", "abCdefAaf", "aabcd", "aaaaaaaaaa", "qwertyzxcvbnt", "qwerty"}
	for _, str := range coupleStrings {
		fmt.Println(str, checkUnique(str))
	}
}

func checkUnique(str string) bool {
	// Преобразовываем каждый символ строки в нижний регистр
	str = strings.ToLower(str)
	// Для проверки уникальности символа в строке воспользуемся мэпом
	// Каждому ключу присвоим true, если ключ уже существует - функция вернет false
	// Если символы уникальны - функция возвращает true
	m := make(map[rune]bool)
	for _, char := range str {
		_, ok := m[char]
		if ok {
			return false
		}
		m[char] = true
	}
	return true
}
