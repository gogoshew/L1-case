package tasks

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func Task11() {
	// Создадим 2 слайса строк
	firstArray := []string{"pen", "pineapple", "apple", "pen", "lol", "gofer"}
	secondArray := []string{"apple", "orange", "pen", "go", "gofer", "apple", "go"}

	set1 := createSet(firstArray)
	set2 := createSet(secondArray)
	result := createIntersection(set1, set2)

	for key := range result {
		fmt.Printf("%s ", key)
	}

}

// Создадим функцию, которая будет заполнять мэп в виде ключ - строка из слайса, значение - bool
// Тем самым в мэп добавятся только уникальные элементы из слайса
func createSet(array []string) map[string]bool {
	set := make(map[string]bool)
	for _, val := range array {
		set[val] = true
	}
	return set
}

// Создадим функцию, которая определяет пересечение множеств
func createIntersection(set1, set2 map[string]bool) map[string]bool {
	set := make(map[string]bool)
	// В цикле определяем есть ли ключ первого множества во втором, если есть - добавляем в пересечение
	for key := range set1 {
		if set2[key] {
			set[key] = true
		}
	}
	return set
}
