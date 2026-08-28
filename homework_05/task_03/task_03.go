/*
Метод k-ближайших соседей – один из простейших алгоритмов машинного обучения.
Суть алгоритма: пусть имеется набор объектов с известными названиями и числовыми характеристиками.
Эти характеристики дают возможность описать объект как точку в n-мерном пространстве.
Пусть есть объект, у которого заданы характеристики, но нет названия.
Чтобы найти название этого объекта, найдём k ближайших к нему известных объектов.
То название, которое будет у ближайших соседей повторяться чаще, сделаем названием нашего неизвестного объекта (алгоритм классификации).
Пример: апельсины и грейпфруты, k=3. Неизвестный объект оказался апельсином.

1. Необходимо реализовать алгоритм метода k-ближайших соседей в двумерном пространстве характеристик.
Вход алгоритма: набор известных объектов, две характеристики неизвестного объекта и значение k.
Каждый объект состоит из строкового названия и двух (для простоты) числовых характеристик.
Выход алгоритма: строка с названием неизвестного объекта.

2. Создайте консольное приложение для тестирования алгоритма.
При старте приложение читает входные данные из текстового файла 'data.txt'.
Затем приложение в бесконечном цикле запрашивает у пользователя две характеристики и выводит строку с названием объекта.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"homework_05/task_03/knn_method"
	"homework_05/task_03/parser"
)

func main() {
	const (
		inputFile = "dataset/data.txt"
		k         = 3
	)

	// Loading data from a file
	fmt.Println("Loading data from", inputFile)
	data, err := parser.ParseTXT(inputFile)
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d objects\n", len(data))

	// Creating a classifier with normalization
	classifier := knn_method.NewClassifier(data, k)
	fmt.Printf("KNN classifier created with k=%d and minimax normalization\n\n", k)

	// Interactive cycle
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter two characteristics (X Y) or 'exit' to exit:")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			break
		}

		// Parsing input
		var x, y float64
		n, err := fmt.Sscanf(input, "%f %f", &x, &y)
		if err != nil || n != 2 {
			fmt.Println("Error: Please enter exactly two numbers separated by a space")
			continue
		}

		// Classify
		result := classifier.Classify(x, y)
		fmt.Printf("Result: %s\n\n", result)
	}

	fmt.Println("Program completed!")
}
