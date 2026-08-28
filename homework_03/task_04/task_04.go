/*
Опишите интерфейс 'Cloner' с методом клонирования 'Clone'.
Создайте две-три простые структуры, которые неявно реализуют интерфейс 'Cloner'.
Создайте функцию 'sliceClone':
Её входной параметр – срез произвольных значений. Функция создаёт и возвращает новый срез, состоящий из клонов значений из входного среза.
Считаем, что клонировать можно все базовые числовые типы, булевский тип, строку, а также структуры, реализующие интерфейс 'Cloner'.
То, что нельзя клонировать, в срез-результат не попадает.
*/

package main

import "fmt"

type Cloner interface {
	Clone() Cloner
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Clone() Cloner {
	return Person{
		Name: p.Name,
		Age:  p.Age,
	}
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Clone() Cloner {
	return Rectangle{
		Width:  r.Width,
		Height: r.Height,
	}
}

type Book struct {
	Title, Author string
	Pages         int
}

func (b Book) Clone() Cloner {
	return Book{
		Title:  b.Title,
		Author: b.Author,
		Pages:  b.Pages,
	}
}

// 'sliceClone' returns a new slice of clones of the given slice.
func sliceClone(slice []any) []any {
	result := make([]any, 0)

	for _, item := range slice {
		switch v := item.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128,
			bool, string:
			result = append(result, v)
		case Cloner:
			result = append(result, v.Clone())
		default:
			// Element can't be cloned, skip it.
		}
	}

	return result
}

func printSlice(slice []any) {
	for i, item := range slice {
		fmt.Printf("[%d]: %v (type: %T)\n", i, item, item)
	}
}

func main() {

	person := Person{Name: "John", Age: 30}
	book := Book{Title: "Programming", Author: "John Doe", Pages: 350}
	rect := Rectangle{Width: 10.5, Height: 20.3}

	// Create a slice of various types
	original := []any{
		42,                     // int
		3.14,                   // float64
		true,                   // bool
		"Hello",                // string
		person,                 // Person
		book,                   // Book
		rect,                   // Rectangle
		[]int{1, 2},            // slice (can't be cloned)
		map[string]int{"a": 1}, // map (can't be cloned)
	}

	fmt.Println("Original slice:")
	printSlice(original)

	// Clone the slice
	cloned := sliceClone(original)

	fmt.Println("\nCloned slice:")
	printSlice(cloned)

	// Check the independence of clones
	fmt.Println("\n--- Check the independence of clones ---")

	fmt.Printf("\nLength of original slice: %d\n", len(original))
	fmt.Printf("Length of cloned slice: %d (elements that cannot be cloned were not included)\n", len(cloned))
}
