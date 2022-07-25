package tasks

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из слайса, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func Task9() {

	// Создаем слайс
	array := []int{2, 4, 6, 8, 10}

	// Создаем конвейер в который передаем слайс
	// Данные из слайса записываются в канал и функция WriteToChan возвращает канал с данными
	pipeline := writeToChan(array)

	// Передаем канал с данными в функицию sqDig, которая считывает их, возводит в квадрат
	// И записывает их в новый канал
	result := sqDig(pipeline)

	// Извлекаем данные из канала и принтим их
	for v := range result {
		fmt.Println(v)
	}

}
func writeToChan(myArray []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range myArray {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func sqDig(inputChan <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for v := range inputChan {
			ch <- v * v
		}
		close(ch)
	}()
	return ch
}
