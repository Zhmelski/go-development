/*
Разрежённые числовые матрицы — это такие матрицы, в которых лишь небольшая часть элементов отлична от нуля.
Например, если в матрице 100x200 ненулевыми являются только четыре угловых элемента, её можно считать разрежённой.
Для экономии памяти нулевые элементы такой матрицы следует не хранить.

Реализуйте обобщённый тип данных для представления разрежённой числовой матрицы.
Элементами матрицы могут быть любые стандартные числовые типы, кроме комплексных, а также типы, определённые на их основе.
Реализуйте методы для чтения и записи элементов по паре индексов.
При недопустимых значениях (например, выход за пределы допустимых индексов) следует вызывать панику.
*/

package main

import "fmt"

// 'Number' define a constraint for numeric types (excluding complex types)
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// 'position' represents a position (row and column) in a matrix
type position struct {
	row, col int
}

// 'SparceMatrix' represents a sparse matrix
type SparceMatrix[T Number] struct {
	rows, cols int
	data       map[position]T
}

// 'NewSparceMatrix' creates a new sparse matrix
func NewSparseMatrix[T Number](rows, cols int) *SparceMatrix[T] {
	if rows <= 0 || cols <= 0 {
		panic("rows and cols must be positive")
	}

	return &SparceMatrix[T]{
		rows: rows,
		cols: cols,
		data: make(map[position]T),
	}
}

// 'Set' sets a value into a matrix cell at index [row, column]
func (sm *SparceMatrix[T]) Set(row, col int, value T) {
	if sm == nil {
		panic("cannot call Set on nil receiver")
	}

	if row < 0 || row >= sm.rows {
		panic("row index is out of range")
	}
	if col < 0 || col >= sm.cols {
		panic("col index is out of range")
	}

	pos := position{row, col}

	// If value is 0, delete the element from the map (memory optimization)
	if value == 0 {
		delete(sm.data, pos)
	} else {
		sm.data[pos] = value
	}
}

// 'Get' returns a value from a matrix cell at index [row, column]
func (sm *SparceMatrix[T]) Get(row, col int) T {
	if sm == nil {
		panic("cannot call Get on nil receiver")
	}

	if row < 0 || row >= sm.rows {
		panic("row index is out of range")
	}
	if col < 0 || col >= sm.cols {
		panic("col index is out of range")
	}

	pos := position{row, col}
	// If the element is not found in the map, a null value of type T is returned
	return sm.data[pos]
}

// 'NonZeroCount' returns the number of non-zero elements in the matrix
func (sm *SparceMatrix[T]) NonZeroCount() int {
	if sm == nil {
		panic("cannot call NonZeroCount on nil receiver")
	}

	return len(sm.data)
}

// 'Rows' return the number of rows in the matrix
func (sm *SparceMatrix[T]) Rows() int {
	if sm == nil {
		panic("cannot call Rows on nil receiver")
	}

	return sm.rows
}

// 'Cols' return the number of columns in the matrix
func (sm *SparceMatrix[T]) Cols() int {
	if sm == nil {
		panic("cannot call Cols on nil receiver")
	}

	return sm.cols
}

// 'PrintMatrix' prints the matrix to the console
func (sm *SparceMatrix[T]) PrintMatrix() {
	if sm == nil {
		panic("cannot call PrintMatrix on nil receiver")
	}

	fmt.Printf("Sparse matrix %dx%d (non-zero elements: %d):\n", sm.rows, sm.cols, sm.NonZeroCount())

	for i := 0; i < sm.rows; i++ {
		for j := 0; j < sm.cols; j++ {
			fmt.Printf("%v\t", sm.Get(i, j))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	fmt.Println("=== Example 1: 5x5 int matrix ===")
	matrix1 := NewSparseMatrix[int](5, 5)

	// Sets only corner elements
	matrix1.Set(0, 0, 10)
	matrix1.Set(0, 4, 20)
	matrix1.Set(4, 0, 30)
	matrix1.Set(4, 4, 40)

	matrix1.PrintMatrix()

	fmt.Printf("Element [0, 0] = %d\n", matrix1.Get(0, 0))
	fmt.Printf("Element [2, 2] = %d (not exist)\n", matrix1.Get(2, 2))
	fmt.Println()

	fmt.Println("=== Example 2: 100x200 float64 matrix ===")
	matrix2 := NewSparseMatrix[float64](100, 200)

	matrix2.Set(0, 0, 1.5)
	matrix2.Set(0, 199, 2.5)
	matrix2.Set(99, 0, 3.5)
	matrix2.Set(99, 199, 4.5)
	matrix2.Set(50, 100, 9.9)

	fmt.Printf("Size: %dx%d\n", matrix2.Rows(), matrix2.Cols())
	fmt.Printf("Non-zero elements: %d\n", matrix2.NonZeroCount())
	fmt.Printf("Element [50, 100] = %.1f\n", matrix2.Get(50, 100))
	fmt.Printf("Element [50, 101] = %.1f (not exist)\n", matrix2.Get(50, 101))
	fmt.Println()

	type MyInt int
	fmt.Println("=== Example 3: User type MyInt ===")
	matrix3 := NewSparseMatrix[MyInt](3, 3)

	matrix3.Set(1, 1, 100)
	matrix3.Set(2, 2, 200)

	matrix3.PrintMatrix()

	fmt.Println("=== Example 4: Zeroing an element ===")
	fmt.Printf("Before zeroing: non-zero elements = %d\n", matrix3.NonZeroCount())
	matrix3.Set(1, 1, 0)
	fmt.Printf("After zeroing [1, 1]: non-zero elements = %d\n", matrix3.NonZeroCount())
	matrix3.PrintMatrix()

	fmt.Println("=== Example 5: Attempt to go beyond the limits(panic) ===")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic: %v\n", err)
		}
	}()

	matrix1.Get(10, 10)
}
