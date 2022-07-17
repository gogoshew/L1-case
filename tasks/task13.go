package tasks

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func Task13() {
	x, y := 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
	fmt.Println(swap(x, y))
}

func swap(a, b int) (int, int) {
	return b, a
}
