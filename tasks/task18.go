package tasks

import (
	"fmt"
	"log"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Counter Создаем структуру-счетчик, поле num для инкрементирования числа, sync.Mutex для конкурентности
type Counter struct {
	num int
	sync.Mutex
}

// Функция increment() блокирует воркера пока значение счетчика не инкрементируется
func (c *Counter) increment() {
	c.Lock()
	c.num++
	c.Unlock()
}

// Функция getValue() возвращает значение счетчика
func (c *Counter) getValue() int {
	return c.num
}

// Функция start() начинает отсчет счетчика
func start(c *Counter, finish chan struct{}) {
	// Создаем WaitGroup
	wg := sync.WaitGroup{}

	// Создаем цикл на 100 итераций, на каждой итерации в wg будет добавляться 1 для успешной работы каждой горутины
	// Одна горутина = одна задача по wg
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int, count *Counter, wg *sync.WaitGroup) {
			// wg.Done() записывается в defer stack и закрывается когда горутина завершает работу
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", n)
			count.increment()
			fmt.Printf("Worker %d done\n", n)
		}(i, c, &wg)
	}
	wg.Wait()
	// Используем use-case для завершения работы программы, создав канал принимающий пустую структуру
	// После выполнения всех задач передаем пустую структуру в канал и закрываем его
	// Когда канал закроется, в основной функции в блоке select сработает case и мы выведем результат счетчика
	// Пустую структуру очень удобно использовать поскольку она весит 0
	finish <- struct{}{}
	close(finish)
}

func Task18() {
	counter := &Counter{
		num: 0,
	}
	finish := make(chan struct{})

	go start(counter, finish)

	select {
	case <-finish:
		log.Printf("Work done with count %d\n", counter.getValue())
	}
}
