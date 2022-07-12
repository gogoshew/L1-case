package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func Task6_1() {
	workerChan := make(chan int)
	// Способ остановки горутины с помощью контекста WithTimeout, также работает с WithDeadline
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	go func() {
		for {
			workerChan <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-workerChan:
			fmt.Println(<-workerChan)
		}
	}
}

func Task6() {
	//ch := closeChanWay()
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//close(ch)
	//fmt.Println("Горутина остановлена...")
	//closeByWaitGroup()
	closeByCtxCancel()
}

func closeChanWay() chan int {
	workerChan := make(chan int)

	// Удобный способ закрывать горутины с помощью закрытия канала
	go func() {
		n := 1
		for {
			select {
			case workerChan <- n:
				n++
				fmt.Println(n)
			case <-workerChan:
				return
			}
		}
	}()
	return workerChan
}

func closeByWaitGroup() {
	// Способ управления с помощью WaitGroup

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Принт из горутины...")
		time.Sleep(time.Second)
	}()
	wg.Wait()
	fmt.Println("Горутина остановлена...")
}

func closeByCtxCancel() {
	// Остановка работы горутины с помощью context.WithCancel
	keyName := "key"
	ctx, cancel := context.WithCancel(context.Background())
	ctxWithValue := context.WithValue(ctx, keyName, "some value")

	go writeSmth(ctxWithValue, "GoGoGo")
	go writeSmth(ctxWithValue, "Coding...")

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 3)
}

func writeSmth(ctx context.Context, someVal string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Горутина %s остановлена...\n", someVal)
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("Горутина со значением %s работает даже ночью...\n", someVal)
		}
	}
}
