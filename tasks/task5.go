package tasks

import (
	"fmt"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func Task5() {
	const n = 5

	// Создаем небуфферезированный канал, поскольку запись и чтение будут синхронны
	workerChan := make(chan int)

	// Создаем таймер для установки времени продолжительности работы
	timer := time.NewTimer(time.Second * n)

	// Используем WaitGroup для ожидания завершения работы воркера
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			num, ok := <-workerChan
			if ok {
				// Канал открыт, ok = true
				fmt.Printf("Значение, которое передал воркер, равняется: %d\n", num)
			} else {
				// Канал закрыт, ok = false
				fmt.Printf("Канал закрыт, работа выполнена")

				// Уменьшаем счетчик группы на 1
				wg.Done()
				return
			}

		}
	}()

	// Пишем основной цикл на 1000 итераций
	for i := 1; i <= 1000; i++ {
		//
		select {

		case <-timer.C:
			fmt.Println("Время вышло...")
			close(workerChan)
			return

		default:
			workerChan <- i
			time.Sleep(time.Second)
		}
	}
	wg.Wait()
}
