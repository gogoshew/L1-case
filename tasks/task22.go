package tasks

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

// Для решения данной таски будем использовать встроенный пакет math/big

/*
Пакет big предоставляет три типа данных:

1. big.Int для крупных целых чисел, когда 18 квинтиллионов недостаточно;
2. big.Float для вещественных чисел с плавающей запятой производной точности;
3. big.Rat для дробей вроде 1/3.
*/

func Task22() {

	// Создадим два числа больше чем 2^20 в формате строки
	a := "3000000000"
	b := "2000000000"

	// Создаем числа с помощью пакета big
	// Используем float потому что при делении возникают проблемы с интовыми числами
	x, _ := new(big.Float).SetString(a)
	if x == nil {
		fmt.Println("Ошибка при создании числа из a!")
		return
	}

	y, _ := new(big.Float).SetString(b)
	if y == nil {
		fmt.Println("Ошибка при создании числа из b!")
		return
	}

	multiple := new(big.Float)
	multiple.Mul(x, y)

	division := new(big.Float)
	division.Quo(x, y)

	sum := new(big.Float)
	sum.Add(x, y)

	sub := new(big.Float)
	sub.Sub(x, y)

	fmt.Println("Умножение чисел = ", multiple)
	fmt.Println("Деление чисел = ", division)
	fmt.Println("Сумма чисел = ", sum)
	fmt.Println("Разность чисел = ", sub)
}
