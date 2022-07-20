package tasks

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func Task25() {
	fmt.Println("Ложимся спать...")
	Sleep(5)
	fmt.Println("Опять работать?..")
}

func Sleep(n int) {
	timeOut := time.After(time.Duration(n) * time.Second)
	<-timeOut
}
