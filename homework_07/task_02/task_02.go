/*
В вашем распоряжении находится файл базы данных SQLite с именем users.db (см. Задачу 1). Разработайте простейшее веб-приложение, позволяющее выполнить следующие REST-запросы (ответ с данными пересылается клиенту в виде JSON):
- GET /users – получить информацию (кроме хеша пароля) о всех пользователях.
- GET /users/id – получить информацию (кроме хеша пароля) о пользователе с Id = id.
- DELETE /users/id – удалить пользователя с Id = id.
*/

/*
Запуск:
Get all users: curl http://localhost:8080/users
Get user by ID: curl http://localhost:8080/users/1
Delete user: curl.exe -X DELETE http://localhost:8080/users/1
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"task_02/database"
	"task_02/handlers"
)

const (
	dbPath = "../task_01/users.db" // Path to the database
	port   = ":8080"
)

func main() {
	// Open an existing database
	db, err := database.OpenDatabase(dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connection to the database has been established successfully")

	// Create a handler for /users
	userHandler := handlers.NewUserHandler(db)

	// Registering routes
	http.Handle("/users", userHandler)
	http.Handle("/users/", userHandler)

	// Starting the server
	fmt.Printf("The server is running on http://localhost%s\n", port)
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  GET    /users      - Get all users")
	fmt.Println("  GET    /users/{id} - Get user by ID")
	fmt.Println("  DELETE /users/{id} - Delete user by ID")
	fmt.Println()

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server startup error: %v", err)
	}
}
