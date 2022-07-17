package tasks

import "fmt"

/*
Разработать программу, которая в рантайме способна определить
тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

type TypeOfVar interface{}

func getType(v TypeOfVar) string {
	return fmt.Sprintf("%T", v)
}

func Task14() {
	var varType TypeOfVar

	varType = 10
	fmt.Printf("Переменная типа: %v\n", getType(varType))

	varType = "go go gofers"
	fmt.Printf("Переменная типа: %v\n", getType(varType))

	varType = true
	fmt.Printf("Переменная типа: %v\n", getType(varType))

	varType = make(chan int)
	fmt.Printf("Переменная типа: %v\n", getType(varType))

}
