package tasks

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func Task17() {
	sortedArray := []int{2, 4, 8, 14, 21, 27, 34, 58, 77, 89, 101}

	findDigit := 58
	result, count := binarySearch(sortedArray, findDigit)
	if result < 0 {
		fmt.Printf("%d отсутствует в слайсе, введите другое число!\n", findDigit)
	} else {
		fmt.Printf("Число %d найдено в слайсе на позиции %d на поиск ушло %d итераций \n\n", findDigit, result, count)
	}
}

func binarySearch(ar []int, n int) (res int, count int) {
	mid := len(ar) / 2

	switch {
	// Если длина слайса после деления равна нулю - элемент не найден в слайсе, возвращаем несуществующий индекс
	case len(ar) == 0:
		res = -1
	// Если элемент, который мы ищем меньше нашего предположения(средний элемент), ищем в левой части слайса
	case ar[mid] > n:
		res, count = binarySearch(ar[:mid], n)
		// Аналогично, только ищем в правой части слайса, если элемент больше
	case ar[mid] < n:
		res, count = binarySearch(ar[mid+1:], n)

		// Если элемент был найден - прибавляем к позиции после вычислений и среднюю позицию среза + 1
		// в итоге получим позицию искомого элемента
		if res >= 0 {
			res += mid + 1
		}
	default: // ar[mid] == n
		// По умолчанию, если искомый элемент равен среднему элементу слайса
		res = mid
	}

	count++
	return res, count
}
