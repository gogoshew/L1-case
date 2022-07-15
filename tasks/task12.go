package tasks

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/

func Task12() {
	strArray := []string{"cat", "cat", "dog", "cat", "tree"}

	var res []string
	for key := range getSet(strArray) {
		res = append(res, key)
	}
	fmt.Println(res)
}

// Создадим функцию, которая будет заполнять мэп в виде ключ - строка из массива, значение - bool
// Тем самым в мэп добавятся только уникальные элементы из массива
func getSet(array []string) map[string]bool {
	set := make(map[string]bool)
	for _, val := range array {
		set[val] = true
	}
	return set
}
