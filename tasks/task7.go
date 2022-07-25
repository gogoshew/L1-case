package tasks

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись данных в map.
*/

func Task7() {
	/*
		Создадим пустую хэш-таблицу, слайс, Mutex и WaitGroup.
		Далее заполним слайс данными, чтобы потом заполнять ими хэш-таблицу.
		Для заполнения данными слайса, я решил сделать цикл с шагом 2.
	*/
	myMap := make(map[int]int)
	array := []int{}
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 2; i <= 20; i += 2 {
		array = append(array, i)
	}

	/*
		Создадим цикл с помощью range по слайсу
		и будем доставать из него номер итерации и значение в слайсе
	*/
	for i, value := range array {

		// Добавляем счетчику WaitGroup 1
		wg.Add(1)

		// Инициализируем анонимную горутину, на ввод запрашиваем значения из слайса
		go func(key int, val int) {

			/*
				Создаем блокировку мьютекса, чтобы ограничить другим горутинам доступ
				на изменение хэш-таблицы пока одна горутина выполняет задачу.
				Разблокируем мьютекс после того, как изменения будут внесены.
			*/
			mtx.Lock()
			myMap[key] = val
			time.Sleep(time.Millisecond * 150)
			mtx.Unlock()

			/*
				Принтим строку, чтобы визуельно понимать, значение под таким-то ключом записалось.
				Далее говорим, что горутина завершила работу и вычитаем из счетчика wg единицу.
			*/
			fmt.Printf("Под ключом %d записано значение %d\n", key, val)
			wg.Done()
		}(i, value)
	}

	// Ждем пока горутины выполнят запись и принтим что работа завершена
	wg.Wait()
	fmt.Println("Работа горутин выполнена...")
}
