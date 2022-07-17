package tasks

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func Task16() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	quickSort(ar)
	fmt.Println(ar)
}

func quickSort(ar []int) {
	// Если длина сортируемого массива меньше либо равна 1 - выходим из функции, в ином случае продолжаем сортировку
	if len(ar) <= 1 {
		return
	}

	// Создаем переменную для разделения массива на 2 части
	split := partition(ar)

	// Сортируем сначала первую часть, затем вторую
	quickSort(ar[:split])
	quickSort(ar[split:])
}

func partition(ar []int) int {

	// Создаем опорную точку pivot
	pivot := ar[len(ar)/2]

	// Устанавливаем границы, левая часть = 0, правая = длина массива - 1
	left := 0
	right := len(ar) - 1

	// Проходим массив, элементы меньше опорной точки перемещаются влево, а больше - вправо
	for {
		for ; ar[left] < pivot; left++ {
		}
		for ; ar[right] > pivot; right-- {
		}
		// Проверка на то, чтобы левый индекс не "обогнал" правый
		if left >= right {
			return right
		}
		swaps(ar, left, right)
	}

}

func swaps(ar []int, i, j int) {
	ar[i], ar[j] = ar[j], ar[i]
}
