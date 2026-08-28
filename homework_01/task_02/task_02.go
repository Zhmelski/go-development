/*
Программа запрашивает положительное целое число. Определить, является ли оно палиндромом в троичной системе счисления.
Вывести "YES", если число — палиндром, иначе "NO".
Пример: ввели число 16. В троичной системе счисления это число записывается как 121. Значит, вывод программы – "YES".
*/

package main

import "fmt"

func main() {

	var inputValue int
	fmt.Println("Введите положительное целое число: ")
	fmt.Scan(&inputValue)

	if inputValue <= 0 {
		fmt.Println("Input error")
		return
	}

	// Преобразуем число в троичную систему и одновременно строим перевёрнутое,
	// после переводя это перевернутое в десятичное
	original := inputValue
	reversedValue := 0
	temp := inputValue

	for temp > 0 {
		reversedValue = reversedValue*3 + temp%3
		temp /= 3
	}

	// Преобразуем исходное число в троичную систему для сравнения
	ternary := 0
	temp = original
	multiplier := 1

	for temp > 0 {
		ternary = ternary + (temp%3)*multiplier
		multiplier *= 10
		temp /= 3
	}

	// Перевернутое число reversedValue (десятичная сис-ма) преобразуем в троичную
	reversedValueRepresentation := 0
	temp = reversedValue
	multiplier = 1

	for temp > 0 {
		reversedValueRepresentation = reversedValueRepresentation + (temp%3)*multiplier
		multiplier *= 10
		temp /= 3
	}

	// Сравниваем прямое и обратное представления в троичной системе
	fmt.Println("Число ", inputValue, " в троичной системе: ", ternary)
	fmt.Print("Является ли палиндромом: ")

	if ternary == reversedValueRepresentation {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
