/*
Создайте структуру 'Matrix' для представления двумерной матрицы чисел (элементы матрицы имеют тип float64).
Поля структуры:
1. 'rows' – число строк матрицы (больше нуля);
2. 'cols' – число столбцов матрицы (больше нуля);
3. 'data' – одномерный срез с элементами матрицы.
Опишите функцию-конструктор для создания экземпляра структуры 'Matrix':
Параметры функции – число строк и число столбцов (целые числа). В случае некорректных значений параметров создавайте матрицу размером 1x1.
Опишите два метода для структуры 'Matrix':
1. Метод 'Get(i, j)' возвращает элемент матрицы с индексами i,j (индексация начинается с нуля);
2. Метод 'Set(i, j, value)' устанавливает элемент матрицы с индексами i,j.
Обратите внимание, что методы должны транслировать два входных индекса в один индекс в срезе с элементами матрицы.
Если индексы выходят за допустимый диапазон, метод 'Get' возвращает значение 0, а метод 'Set' ничего не делает.
Создайте метод 'Print' для структуры 'Matrix'. Он должен выводить матрицу на консоль.
*/

package main

import "fmt"

// 'Matrix' represents a two-dimensional matrix of numbers
type Matrix struct {
	rows int
	cols int
	data []float64
}

// 'NewMatrix' creates a new matrix with the given dimensions.
func NewMatrix(rows, cols int) *Matrix {

	// Checking the correctness of parameters
	if rows <= 0 || cols <= 0 {
		rows = 1
		cols = 1
	}

	return &Matrix{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}
}

// 'Get' returns the matrix element with indices i, j
func (m *Matrix) Get(i, j int) float64 {
	if m == nil {
		return 0
	}

	// Bounds checking
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return 0
	}

	// Translation of two-dimensional indices into one-dimensional ones
	index := i*m.cols + j
	return m.data[index]
}

// 'Set' sets the element of the matrix with indices i, j
func (m *Matrix) Set(i, j int, value float64) {
	if m == nil {
		return
	}

	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return
	}

	index := i*m.cols + j
	m.data[index] = value
}

// 'Print' outputs the matrix to the console.
func (m *Matrix) Print() {
	if m == nil {
		fmt.Println("Matrix is nil")
		return
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			fmt.Printf("%.2f ", m.Get(i, j))
		}
		fmt.Println()
	}
}

func main() {

	// Create a matrix
	matrix := NewMatrix(3, 4)

	// Fill the matrix
	matrix.Set(0, 0, 1.5)
	matrix.Set(0, 1, 2.3)
	matrix.Set(1, 2, 7.8)
	matrix.Set(2, 3, 9.1)

	// Print the matrix
	fmt.Println("Matrix 3x4:")
	matrix.Print()

	// Access elements
	fmt.Printf("\nElement [1][2] = %.2f\n", matrix.Get(1, 2))
	fmt.Printf("Element [0][0] = %.2f\n", matrix.Get(0, 0))

	// Access beyond borders
	fmt.Printf("Element [5][5] (beyond borders) = %.2f\n", matrix.Get(5, 5))

	// Create a matrix with incorrect parameters
	fmt.Println("\nMatrix with incorrect parameters (-2, 5) -> 1x1:")
	invalidMatrix := NewMatrix(-2, 5)
	invalidMatrix.Set(0, 0, 42)
	invalidMatrix.Print()
}
