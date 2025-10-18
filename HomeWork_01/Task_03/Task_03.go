/*
Система оценок в США (в зависимости от процента от максимального количества баллов):
A – 90–100%,
B – 80–89%,
C – 70–79%,
D – 65–69%,
F – 64% и ниже.
Программа запрашивает максимальное количество баллов за одно задание, количество заданий, а затем — баллы за каждое задание.
(Имеется в виду, что некий студент выполнил ряд заданий, и мы хотим оценить успеваемость этого студента.)
Программа выводит оценку в виде буквы. Считаем, что ввод пользователя всегда корректен и его можно не проверять.
*/

package main

import "fmt"

func main() {

	var maxScore, taskCount int

	fmt.Print("Введите максимальное количество баллов за одно задание: ")
	fmt.Scanln(&maxScore)

	fmt.Print("Введите количество заданий: ")
	fmt.Scanln(&taskCount)

	// Подсчитываем общую сумму баллов
	totalScore := 0
	for i := 1; i <= taskCount; i++ {
		var score int
		fmt.Printf("Введите баллы за задание %d: ", i)
		fmt.Scanln(&score)
		totalScore += score
	}

	// Вычисляем процент
	maxPossibleScore := maxScore * taskCount
	percentage := (totalScore * 100) / maxPossibleScore

	// Определяем оценку
	var grade string
	switch {
	case percentage >= 90:
		grade = "A"
	case percentage >= 80:
		grade = "B"
	case percentage >= 70:
		grade = "C"
	case percentage >= 65:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("Оценка: %s\n", grade)
}
