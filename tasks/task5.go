package tasks

import (
	"context"
	"fmt"
	"time"
)

const n = 10

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func Task5() {

	// Создаем небуфферезированный канал, поскольку запись и чтение будут синхронны
	workerChan := make(chan int)

	// Создаем контекст с тайм-аутом
	ctx, _ := context.WithTimeout(context.Background(), time.Second*n)

	// Создаем аннонимную горутину для постоянной записи текущих секунд в канал
	go func() {
		for {
			workerChan <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}()

	// Создаем бесконечный цикл, в который помещаем блок select
	// Когда case читает данные из workerChan печатаем результат
	// Когда case получает данные из ctx.Done() завершаем работу горутины.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Время вышло...")
			return
		case <-workerChan:
			fmt.Printf("Результат: %d\n", <-workerChan)
		}
	}
}
