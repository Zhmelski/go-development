/*
Реализуйте обобщённую версию функции 'where' (см. ДЗ №2, задача 2) с параллельной обработкой.
Запустите восемь горутин, каждая из которых фильтрует свою часть исходного среза. Затем объедините результаты в единый итоговый срез.

Обратите внимание: результаты работы обычной и распараллеленной версии функции 'where' могут отличаться по порядку элементов в результирующем срезе, хотя сами элементы будут идентичны.

Выполните замеры производительности. Дайте ответ: наблюдается ли прирост скорости при использовании горутин?
При каком размере входного среза этот прирост становится ощутимым (на вашем компьютере)?

Примечание: при решении этой задачи разрешено использовать пакеты math/rand и time.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 'where' is the regular sequential version of filtering
func where[T any](slice []T, predicate func(T) bool) []T {
	if predicate == nil {
		return []T{}
	}

	result := []T{}
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// 'whereParallel' — a parallel version of filtering using 8 goroutines
func whereParallel[T any](slice []T, predicate func(T) bool) []T {
	if predicate == nil {
		return []T{}
	}

	const grCount = 8
	length := len(slice)

	// If the slice is too small, use the sequential version
	if length < grCount {
		return where(slice, predicate)
	}

	// Create a channel to receive results from goroutines
	type result struct {
		index int
		data  []T
	}
	resultChan := make(chan result, grCount)

	chunkSize := (length + grCount - 1) / grCount

	for i := range grCount {
		go func(goroutineIndex int) {
			// Calculate the range of slice for current goroutine
			start := goroutineIndex * chunkSize
			end := min(start+chunkSize, length)

			// The last goroutine receives the rest of the elements
			if goroutineIndex == grCount-1 {
				end = length
			}

			// Filter elements in the current goroutine
			localResult := []T{}
			for j := start; j < end; j++ {
				if predicate(slice[j]) {
					localResult = append(localResult, slice[j])
				}
			}

			// Send the result to the channel
			resultChan <- result{index: goroutineIndex, data: localResult}
		}(i)
	}

	// Collect results from all goroutines
	resultSlice := make([][]T, grCount)
	for range grCount {
		res := <-resultChan
		resultSlice[res.index] = res.data
	}
	close(resultChan)

	// Merge results from all goroutines in one slice
	finalResult := []T{}
	for _, res := range resultSlice {
		finalResult = append(finalResult, res...)
	}

	return finalResult
}

// 'timeDecorator' is a decorator for 'where' and 'whereParallel' functions to measure the execution time
func timeDecorator[T any](name string, f func([]T, func(T) bool) []T) func([]T, func(T) bool) []T {
	return func(slice []T, predicate func(T) bool) []T {
		start := time.Now()
		result := f(slice, predicate)
		elapsed := time.Since(start)
		fmt.Printf("%s: %v (results: %d)\n", name, elapsed, len(result))
		return result
	}
}

// 'generateRandomSlice' generates a random slice of integers with the specified size
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func runBenchmark(size int) {
	fmt.Printf("\n=== Slice size: %d ===\n", size)

	data := generateRandomSlice(size)

	// Predicate: elements greater than 500
	predicate := func(x int) bool {
		return x > 500
	}

	// Measure the sequential version
	whereSequential := timeDecorator("Sequential", where[int])
	result1 := whereSequential(data, predicate)

	// Measure the parallel version
	whereParallelTimed := timeDecorator("Parallel", whereParallel[int])
	result2 := whereParallelTimed(data, predicate)

	// Check if the number of elements is the same
	if len(result1) != len(result2) {
		fmt.Printf("Results are different! (%d vs %d)\n", len(result1), len(result2))
	}
}

func main() {

	fmt.Println("\n=== Test 1: Correctness check ===")
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	predicate := func(x int) bool {
		return x%2 == 0
	}

	result1 := where(testData, predicate)
	result2 := whereParallel(testData, predicate)

	fmt.Printf("Original slice: %v\n", testData)
	fmt.Printf("Sequential: %v\n", result1)
	fmt.Printf("Parallel: %v\n", result2)

	fmt.Println("\n=== Test 2: Performance check ===")

	sizes := []int{1_000, 10_000, 50_000, 100_000, 1_000_000, 10_000_000}

	for _, size := range sizes {
		runBenchmark(size)
	}

	fmt.Println("\n=== Test 3: Working with different types ===")

	strings := []string{"apple", "banana", "apricot", "blueberry", "avocado", "berry"}
	startsWithA := whereParallel(strings, func(s string) bool {
		return len(s) > 0 && s[0] == 'a'
	})
	fmt.Printf("Strings starting with 'a': %v\n", startsWithA)

	floats := []float64{1.5, 2.7, 3.2, 4.9, 5.1, 6.3, 7.8}
	greaterThanFour := whereParallel(floats, func(f float64) bool {
		return f > 4.0
	})
	fmt.Printf("Numbers > 4.0: %v\n", greaterThanFour)

	// Conclusions
	fmt.Println("\n\nConclusions:")
	fmt.Println("For small sizes (<100k) parallelism may be slower")
	fmt.Println("The increase becomes noticeable starting from ~1 million elements")
}
