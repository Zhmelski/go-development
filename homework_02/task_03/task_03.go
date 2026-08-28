/*
Создать функцию `sequence`, получающую на вход произвольное количество целых чисел.
Логика работы функции зависит от количества её аргументов:
- одно число `a`: вернуть срез с целыми числами от `0` до `a` включительно (или от `a` до `0`, если `a` — отрицательное);
- два числа `a` и `b`: вернуть срез с целыми числами от `a` до `b` (или от `b` до `a`, если `b` меньше `a`);
- чисел больше двух — вернуть срез с этими числами;
- нет чисел — вернуть пустой срез.
*/

package main

import "fmt"

func generateSliceResult(a, b int) []int {

	if b < a {
		a, b = b, a
	}

	result := make([]int, b-a+1)
	for i := range len(result) {
		result[i] = a + i
	}

	return result
}

func sequence(numbers ...int) []int {

	switch len(numbers) {
	case 0:
		return []int{}
	case 1:
		return generateSliceResult(0, numbers[0])
	case 2:
		return generateSliceResult(numbers[0], numbers[1])
	default:
		return numbers
	}
}

func main() {

	fmt.Println("Тест 1: Нет аргументов")
	result1 := sequence()
	fmt.Println("Результат: ", result1)

	fmt.Println("Тест 2: Один аргумент")
	result2 := sequence(-5)
	fmt.Println("Результат: ", result2)

	fmt.Println("Тест 3: Два аргумента")
	result3 := sequence(-3, 4)
	fmt.Println("Результат: ", result3)

	fmt.Println("Тест 4: Больше двух аргументов")
	result4 := sequence(1, 4, 7, 10)
	fmt.Println("Результат: ", result4)

}
