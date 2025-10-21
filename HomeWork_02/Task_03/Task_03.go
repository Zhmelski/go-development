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

func sequence(numbers ...int) (result []int) {
	length := len(numbers)

	if length == 0 {
		return []int{}
	}

	if length == 1 {
		a := numbers[0]
		result := []int{}

		if a >= 0 {
			for i := 0; i <= a; i++ {
				result = append(result, i)
			}
		} else {
			for i := a; i <= 0; i++ {
				result = append(result, i)
			}
		}

		return result
	}

	if length == 2 {
		a := numbers[0]
		b := numbers[1]
		result := []int{}

		if a >= b {
			a, b = b, a
		}

		for i := a; i <= b; i++ {
			result = append(result, i)
		}

		/*if a <= b {
			for i := a; i <= b; i++ {
				result = append(result, i)
			}
		} else {
			for i := a; i >= b; i-- {
				result = append(result, i)
			}
		}*/

		return result
	}

	result = make([]int, length)
	copy(result, numbers)
	return result
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
