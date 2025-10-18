/*
Программа запрашивает у пользователя натуральное число, меньшее 1_000_000_000.
Если число выходит за диапазон, вывести "Input error" и завершить работу.
Иначе вывести наибольшую цифру этого числа.
Пример: у числа 12321 наибольшая цифра – это 3.
*/

package main

import "fmt"

func main() {

	var inputValue int
	fmt.Println("Введите натуральное число, меньшее 1 000 000: ")
	fmt.Scan(&inputValue)

	if inputValue <= 0 || inputValue >= 1_000_000 {
		fmt.Println("Input error")
		return
	}

	maxDigit := 0
	for inputValue > 0 {
		digit := inputValue % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		inputValue /= 10
	}
	fmt.Println("Максимальная цифра введенного чилса: ", maxDigit)
}
