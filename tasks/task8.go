package tasks

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func Task8() {

	res := setPosition(8, 2, 1)
	fmt.Printf("Результат: %d (%b)", res, res)
}

// Создадим функцию, которая будет принимать десятичное число, позицию бита в числе и значение бита на этой позиции
func setPosition(num int64, position uint8, bit uint8) int64 {
	// Для работы с конкретной позицией сдвинем единицу на определенный разряд
	shift := int64(1) << position

	switch bit {
	case 0:
		// | Поразрядное сложение. Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
		// ^ Поразрядное исключающее ИЛИ. Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
		res := (num | shift) ^ shift
		fmt.Printf("[%d (%b) | %d (%b)] ^ %d (%b) = %d (%b)\n", num, num, shift, shift, shift, shift, res, res)
		return res

	case 1:
		res := num | shift
		fmt.Printf("%d (%b) | %d (%b) = %d (%b)\n", num, num, shift, shift, res, res)
		return res
	default:
		return num
	}
}
