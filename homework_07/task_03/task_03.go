/*
В предыдущем домашнем задании вы создавали консольное приложения для управления списком задач – TODO App.
Дополните (!) этот программный проект веб-приложением, отображающим список задач в виде таблицы.
*/

/*
Запуск:
go run task_03.go -task="Купить продукты"

go run task_03.go -list

go run task_03.go -complete=1

go run task_03.go -task="Сделать дз" -list
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"task_02/storage"
	"task_02/todoapp"
)

func main() {
	// Defining command line flags
	listFlag := flag.Bool("list", false, "Show list of all tasks")
	taskDesc := flag.String("task", "", "Add a new task")
	completeID := flag.Int("complete", 0, "Mark task as completed (specify number)")

	flag.Parse()

	// Creating an application instance
	app, err := todoapp.NewTodoApp(storage.DefaultFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Initialization error: %v\n", err)
		os.Exit(1)
	}

	// Processing flags
	actionPerformed := false

	// Adding a new task
	if *taskDesc != "" {
		if err := app.AddTask(*taskDesc); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error adding task: %v\n", err)
			os.Exit(1)
		}
		actionPerformed = true
	}

	// Completing a task
	if *completeID > 0 {
		if err := app.CompleteTask(*completeID); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Task completion error: %v\n", err)
			os.Exit(1)
		}
		actionPerformed = true
	}

	// Displaying a list of tasks
	if *listFlag {
		app.ListTasks()
		actionPerformed = true
	}

	// If no action was performed, show the help
	if !actionPerformed {
		fmt.Println("TODO App - Task Manager")
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  go run main.go -task=\"Купить продукты\"")
		fmt.Println("  go run main.go -list")
		fmt.Println("  go run main.go -complete=1")
		fmt.Println("  go run main.go -task=\"Новая задача\" -list")
	}
}
