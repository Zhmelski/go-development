/*
Создать функцию `where`, которая получает на вход срез целых чисел и функцию-предикат.
Функция-предикат получает на вход целое число и возвращает `true/false`.
Функция `where` возвращает срез, состоящий из чисел среза-аргумента, удовлетворяющих функции-предикату.
Создать функцию `foreach` для среза целых чисел и функции-действия.
Функция-действие выполняет некий код для своего целого аргумента.
Функция `foreach` запускает функцию-действие для каждого числа из своего среза-аргумента.
Протестировать работу функций `where` и `foreach`, используя в качестве предиката и действия анонимные функции.
*/

package main

import "fmt"

func where(numbers []int, predicate func(int) bool) []int {

	if predicate == nil {
		return []int{} // Возвращаем пустой срез, если predicate передается nil
	}

	result := []int{}
	for _, number := range numbers {
		if predicate(number) {
			result = append(result, number)
		}
	}
	return result
}

func forEach(numbers []int, action func(int)) {

	if action == nil {
		return
	}

	for _, number := range numbers {
		action(number)
	}
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}

	fmt.Println("Первоначальный срез: ", numbers)

	// Тест 1: Вывод каждого числа
	fmt.Println("\nТест 1: все числа")
	fmt.Print("Результат: ")
	forEach(numbers, func(number int) {
		fmt.Print(number, " ")
	})

	// Тест 2: Фильтрация четных чисел
	fmt.Println("\nТест 2: четные числа")
	evenNumbers := where(numbers, func(number int) bool {
		return number%2 == 0
	})
	fmt.Println("Результат:", evenNumbers)

	// Тест 3: Фильтрация положительных чисел
	fmt.Println("Тест 3: положительные числа")
	positiveNumbers := where(numbers, func(number int) bool {
		return number > 0
	})
	fmt.Println("Результат:", positiveNumbers)

	// Тест 4: Возведение каждого положительного числа в квадрат
	fmt.Println("Тест 4: квадраты положительных чисел")
	fmt.Print("Результат: ")
	squareNumbers := where(numbers, func(number int) bool {
		return number > 0
	})
	forEach(squareNumbers, func(number int) {
		fmt.Print(number*number, " ")
	})
	fmt.Println()

	// Тест 5: Сумма всех чисел
	fmt.Println("Тест 5: сумма всех чисел")
	sum := 0
	forEach(numbers, func(number int) {
		sum += number
	})
	fmt.Println("Результат:", sum)
}
