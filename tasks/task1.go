package tasks

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// Создадим родительскую структуру Human и укажем что она имеет 2 свойства Name и Age

type Human struct {
	Name string
	Age  uint
}

type Action struct {
	Human // Встраиваем структуру Human в Action
	Hobby string
}

// Создадим метод структуры Human

func (h *Human) sayHi() string {
	return fmt.Sprintf("Hi my name is %s and I'm %d", h.Name, h.Age)
}

/*
Создадим экземпляр структуры Action, зададим ему атрибуты.
Выведем строку с использованием встроенного метода sayHi
*/

func Task1() {
	humanAction := Action{Human{"Ilya", 24}, "poopcoding :)"}
	fmt.Printf("%s, I'm %s", humanAction.sayHi(), humanAction.Hobby)
}
