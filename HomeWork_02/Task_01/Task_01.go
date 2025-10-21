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

func CheckNumberDivisors(number int) (string, bool) {

	// Проверяем, что число > 1.
	if number <= 1 {
		return "", false
	}

	// Считаем количество простых делителей.
	count := 0
	temp := number

	for temp%2 == 0 {
		count++
		temp /= 2
	}

	// Проверяем нечетные делители
	for i := 3; i <= temp; i += 2 {
		for temp%i == 0 {
			count++
			temp /= i
		}
	}

	if count%2 == 0 {
		return "even ", true
	}
	return "odd ", true
}

func main() {

	var inputValue int
	fmt.Println("Введите целое число: ")
	fmt.Scan(&inputValue)

	fmt.Print(CheckNumberDivisors(inputValue))
}
