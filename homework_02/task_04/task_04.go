/*
Функция `group` получает в качестве аргумента карту, в которой ключ имеет тип `byte`, а значение для ключа — тип `string`.
Результат работы функции — карта с байтовыми ключами из диапазона `0...9` и значениями для ключа в виде срезов из строк.
Функция `group` выполняет группировку значений из исходной карты по признаку последней цифры числового ключа.
Например, если в исходной карте есть пары `11: "red"` и `51: "green"`, то они должны быть сгруппированы в пару `1: ["red", "green"]`.
При этом пустых срезов в карте-результате быть не должно.
*/

package main

import "fmt"

func group(input map[byte]string) map[byte][]string {
	result := make(map[byte][]string)

	for key, value := range input {
		lastDigit := key % 10
		result[lastDigit] = append(result[lastDigit], value)
	}

	return result
}

func main() {

	// Тест 1: Числа и цвета
	fmt.Println("Тест 1: числа и цвета")
	input1 := map[byte]string{
		11: "red",
		51: "green",
		21: "blue",
		32: "yellow",
		42: "orange",
	}
	fmt.Println("Входная карта:", input1)
	result1 := group(input1)
	fmt.Println("Результат:", result1)

	// Тест 2: Все ключи с разными последними цифрами
	fmt.Println("\nТест 2: Разные последние цифры")
	input2 := map[byte]string{
		10: "zero",
		11: "one",
		12: "two",
		13: "three",
		14: "four",
		15: "five",
		16: "six",
		17: "seven",
		18: "eight",
		19: "nine",
	}
	fmt.Println("Входная карта:", input2)
	result2 := group(input2)
	fmt.Println("Результат:", result2)

	// Тест 3: Пустая карта
	fmt.Println("\nТест 3: Пустая карта")
	input5 := map[byte]string{}
	fmt.Println("Входная карта:", input5)
	result5 := group(input5)
	fmt.Println("Результат:", result5)
}
