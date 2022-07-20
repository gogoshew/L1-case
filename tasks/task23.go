package tasks

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func Task23() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(deleteFromSliceInOrder(4, slice))
	//fmt.Println(deleteFromSliceWithoutOrder(6, slice))
}

func deleteFromSliceInOrder(i int, slice []int) []int {
	// Присваиваем переменной slice значение входящего слайса
	// Используем append добавляя к слайсу до i-го элемента не включительно все значение после i-го элемента
	// имитируя тем самым удаление i-го элемента
	slice = append(slice[:i], slice[i+1:]...) // "..." позволяет добавить сразу несколько элементов к слайсу
	return slice
}

// deleteFromSliceWithoutOrder также удаляет i-ый элемент из слайса, только с помощью перемещения последнего элемента
// на место i-го
func deleteFromSliceWithoutOrder(i int, slice []int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
