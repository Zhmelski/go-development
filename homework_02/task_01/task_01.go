/*
Создайте функцию, которая работает следующим образом.
Функция получает на вход целое число (тип int) и возвращает пару значений.
Первое значение в паре — строка "even" или "odd", в зависимости от количества простых делителей входного целого числа.
То есть:
- 15 = 3 × 5 → even
- 30 = 2 × 3 × 5 → odd
- 9 = 3 × 3 → even
- 7 = 7 → odd
Второе значение в паре — булевский флаг.
Он равен true, если входное целое число больше или равно 2, и false, если входное число меньше или равно единице (в этом случае первое значение в паре — пустая строка).
В основной функции main протестировать созданную функцию на нескольких входных данных.
*/

package main

import "fmt"

func FindingDivisors(number int) (result []int) {

	temp := number

	// Считаем количество простых делителей.
	for temp%2 == 0 {
		result = append(result, 2)
		temp /= 2
	}

	// Проверяем нечетные делители
	for i := 3; i*i <= temp; {
		if temp%i == 0 {
			result = append(result, i)
			temp /= i
		} else {
			i += 2
		}
	}

	if temp > 2 {
		result = append(result, temp)
	}

	return result
}

func CheckNumberDivisors(number int) (string, bool) {

	// Проверяем, что число > 1.
	if number <= 1 {
		return "", false
	}

	factors := FindingDivisors(number)
	if len(factors)%2 == 0 {
		return "even", true
	} else {
		return "odd", true
	}
}

func main() {

	var inputValue int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&inputValue)

	fmt.Println(CheckNumberDivisors(inputValue))
}
