/*
Реализуйте обобщённый тип данных для стека, используя односвязный список.
Добавьте методы 'Push', 'Pop' и 'IsEmpty'. Метод 'Pop' должен возвращать пару: значение и ошибку.
Тип ошибки — стандартный интерфейс 'error' (используйте пакет 'errors').
Ошибку следует возвращать при попытке извлечь элемент из пустого стека.
*/

package main

import (
	"errors"
	"fmt"
)

// 'LinkedList' is a node of a singly linked list
type LinkedList[T any] struct {
	data T
	next *LinkedList[T]
}

// 'Stack' represents a generalized stack based on a singly linked list
type Stack[T any] struct {
	top *LinkedList[T]
}

// 'Push' adds a new node to the top of the stack
func (s *Stack[T]) Push(data T) {
	if s == nil {
		panic("cannot call Push on nil receiver")
	}

	newNode := &LinkedList[T]{
		data: data,
		next: s.top,
	}
	s.top = newNode
}

// 'IsEmpty' checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

// 'Pop' extracts the top node from the stack and returns its value
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty, can't pop")
	}

	value := s.top.data
	s.top = s.top.next
	return value, nil
}

// 'PrintStackValues' prints the stack values to the console
func (s *Stack[T]) PrintStackValues() {
	if s == nil {
		panic("cannot call PrintStackValues on nil receiver")
	}

	for !s.IsEmpty() {
		value, err := s.Pop()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(value)
	}
	fmt.Println()
}

func main() {

	stringStack := Stack[string]{}

	stringStack.Push("First")
	stringStack.Push("Second")
	stringStack.Push("Third")

	fmt.Println("Extraction elements from the stack:")
	stringStack.PrintStackValues()

	// Attempt to pop from an empty stack
	_, err := stringStack.Pop()
	if err != nil {
		fmt.Println("Error popping from empty stack:", err)
	}

}
