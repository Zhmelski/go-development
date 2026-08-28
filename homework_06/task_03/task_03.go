/*
Разработайте структуру для представления кеша объектов в памяти. Каждый элемент кеша ассоциирован со строковым ключом. Кеш должен иметь ограниченную ёмкость и механизм автоматического удаления неиспользуемых объектов. Реализация должна быть оформлена в отдельном пакете. Предусмотрите консольное приложение для тестирования.
1) Функция создания кеша должна принимать в качестве аргумента положительное целое число. Это ёмкость кеша, то есть максимальное количество элементов в кеше.
2) Снабдите кеш методом Set(key string, obj any) для добавления объекта в кеш. Если в кеше уже есть объект с ключом key, то этот старый объект заменяется на obj (с обновлением времени последнего доступа). Если количество элементов в кеше равно его ёмкости, то перед добавлением нового объекта из кеша удаляется элемент, к которому дольше всего не обращались.
3) Снабдите кеш методом Get(key string) any для получения объекта по ключу. Если в кеше нет объекта с ключом key, метод возвращает значение nil. Если в кеше есть объект с ключом key, метод возвращает этот объект, при этом время последнего доступа к элементу кеша обновляется.
4) Создайте метод Remove(key string) bool для удаления из кеша объекта по ключу. Если объект нашли по ключу key и удалили, метод возвращает значение true. Если в кеше нет объекта с ключом key, метод возвращает значение false.
5) После создания, кеш должен автоматически сканировать свои элементы (раз в секунду) и удалять те элементы, к которым не обращались более 10 секунд.
*/

package main

import (
	"fmt"
	"time"

	"task_03/cache"
)

func main() {
	fmt.Println("=== Cache Testing Application ===\n")

	// Create a cache with a capacity of 3
	c := cache.NewCache(3)
	defer c.Stop()

	// Test 1: Basic Set and Get operations
	fmt.Println("Test 1: Basic Set/Get operations")
	c.Set("key1", "value1")
	c.Set("key2", 42)
	c.Set("key3", []string{"a", "b", "c"})

	fmt.Printf("key1: %v\n", c.Get("key1"))
	fmt.Printf("key2: %v\n", c.Get("key2"))
	fmt.Printf("key3: %v\n", c.Get("key3"))
	fmt.Printf("Cache size: %d\n\n", c.Size())

	// Test 2: Cache Overflow (LRU elimination)
	fmt.Println("Test 2: Cache overflow (LRU eviction)")
	fmt.Println("Adding key4 (should remove key1 as least recently used)")
	c.Set("key4", "value4")

	fmt.Printf("key1 (should be nil): %v\n", c.Get("key1"))
	fmt.Printf("key2: %v\n", c.Get("key2"))
	fmt.Printf("key4: %v\n", c.Get("key4"))
	fmt.Printf("Cache size: %d\n\n", c.Size())

	// Test 3: Updating an existing key
	fmt.Println("Test 3: Update existing key")
	c.Set("key2", "updated_value")
	fmt.Printf("key2 (updated): %v\n", c.Get("key2"))
	fmt.Printf("Cache size: %d\n\n", c.Size())

	// Test 4: Remove method
	fmt.Println("Test 4: Remove method")
	removed := c.Remove("key3")
	fmt.Printf("Removed key3: %v\n", removed)
	fmt.Printf("key3 (should be nil): %v\n", c.Get("key3"))
	fmt.Printf("Cache size: %d\n\n", c.Size())

	notRemoved := c.Remove("nonexistent")
	fmt.Printf("Tried to remove nonexistent key: %v\n", notRemoved)
	fmt.Printf("Cache size: %d\n\n", c.Size())

	// Test 5: Automatic deletion on timeout
	fmt.Println("Test 5: Automatic expiration (waiting 12 seconds)")
	c.Set("temp1", "will expire")
	c.Set("temp2", "will expire too")
	fmt.Printf("Cache size before expiration: %d\n", c.Size())

	fmt.Println("Waiting for items to expire...")
	time.Sleep(12 * time.Second)

	fmt.Printf("Cache size after expiration: %d\n", c.Size())
	fmt.Printf("temp1 (should be nil): %v\n", c.Get("temp1"))
	fmt.Printf("temp2 (should be nil): %v\n\n", c.Get("temp2"))

	// Test 6: Updating access time
	fmt.Println("Test 6: Access time update prevents expiration")
	c.Set("persistent", "I will survive")
	fmt.Println("Added 'persistent' key")

	// Access the key every 5 seconds so that it does not get deleted
	for i := 1; i <= 3; i++ {
		time.Sleep(5 * time.Second)
		value := c.Get("persistent")
		fmt.Printf("After %d seconds - persistent: %v\n", i*5, value)
	}

	fmt.Printf("Cache size: %d\n", c.Size())

	fmt.Println("\n=== All tests completed ===")
}
