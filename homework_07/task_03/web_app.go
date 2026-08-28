/*
Запуск:
go run WebApp.go
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"task_02/storage"
	"task_02/todoapp"
	"task_02/web"
)

func main() {
	// Creating an application instance
	app, err := todoapp.NewTodoApp(storage.DefaultFileName)
	if err != nil {
		log.Fatalf("❌ Initialization error: %v\n", err)
	}

	// Creating a web handler
	handler, err := web.NewHandler(app)
	if err != nil {
		log.Fatalf("❌ Error creating handler: %v\n", err)
	}

	// Setting up routes
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/add", handler.AddTaskHandler)
	http.HandleFunc("/complete", handler.CompleteTaskHandler)
	http.HandleFunc("/delete", handler.DeleteTaskHandler)

	// Starting the server
	port := ":8080"
	fmt.Printf("🚀 TODO App launched on http://localhost%s\n", port)
	fmt.Println("📋 Open your browser and go to the address above.")
	fmt.Println("⏹  Press Ctrl+C to stop the server.")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("❌ Server startup error: %v\n", err)
	}
}
