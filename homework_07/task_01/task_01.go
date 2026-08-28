/*
Вы принимаете участие в разработке подсистемы для управления пользователями. Каждый Пользователь имеет следующие атрибуты:
- Идентификатор. Тип данных – int. Уникальный, обязательный.
- Email. Адрес электронной почты пользователя. Используется для отправки почты и как логин пользователя в системе. Уникальный, обязательный.
- Хеш пароля (password hash). Произвольная строка. Используется для сверки введённого при входе в систему пароля. Обязательный.
- Имя пользователя. Произвольная строка. Обязательный.
- Флаг «активен». Тип данных – bool (в базе данных - bit). Обязательный.
Вам необходимо создать консольную утилиту, выполняющую следующие действия:
1. При старте утилита создаёт файл базы данных SQLite с именем users.db, а затем создаёт в этой базе таблицу Users с необходимой структурой (см. атрибуты Пользователя).
2. После создания базы и таблицы утилита вносит в таблицу Users данные пользователей из текстового файла persons.txt. В файле построчно перечислены пользователи – в каждой строке указаны Имя и Фамилия пользователя через пробел. Например:
Alexey Volosevich
John Doe
Mary Robinson
3. Правила заполнения таблицы Users:
- Идентификатор генерируется автоматически.
- Email формируется по шаблону FirstnameLastname@coolcompany.com.
- Хеш пароля вычисляется по Email с помощью функции GenerateFromPassword() из пакета bcrypt.
- Флаг «активен» устанавливается в true (1).
4. Если при разборе строки из файла persons.txt обнаруживаются некорректные данные, такая строка записывается в файл errors.txt вместе со своим порядковым номером.
5. После заполнения таблицы Users утилита выполняет контрольное чтение всех данных из этой таблицы и выводит их на консоль.
*/

package main

import (
	"fmt"
	"log"

	"task_01/database"
	"task_01/services"
	"task_01/utils"
)

func main() {
	// 1. Database initialization
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// 2. Parsing the persons.txt file
	persons, parseErrors, err := utils.ParsePersonsFile("persons.txt")
	if err != nil {
		log.Fatalf("Error parsing file persons.txt: %v", err)
	}

	// 3. Recording parsing errors in errors.txt
	if len(parseErrors) > 0 {
		err = utils.WriteErrors(parseErrors)
		if err != nil {
			log.Printf("Error writing file errors.txt: %v", err)
		} else {
			fmt.Printf("%d invalid lines found, written to errors.txt", len(parseErrors))
		}
	}

	// 4. Adding users to the database
	fmt.Printf("\nAdding users to the database...\n")
	successCount := 0
	for _, person := range persons {
		user, err := services.CreateUserFromPerson(person)
		if err != nil {
			log.Printf("Error creating user for '%s %s': %v", person.FirstName, person.LastName, err)
			continue
		}

		err = database.InsertUser(db, user)
		if err != nil {
			log.Printf("Error adding user '%s': %v", user.Name, err)
			continue
		}
		successCount++
	}
	fmt.Printf("Users added successfully: %d\n", successCount)

	// 5. Control reading and data output
	fmt.Println("\n=== Test reading of data from the Users table ===")
	users, err := database.GetAllUsers(db)
	if err != nil {
		log.Fatalf("Error reading users: %v", err)
	}

	if len(users) == 0 {
		fmt.Println("The Users table is empty")
	} else {
		fmt.Printf("\nTotal users: %d\n\n", len(users))
		for _, user := range users {
			fmt.Printf("ID: %d\n", user.ID)
			fmt.Printf("  Email: %s\n", user.Email)
			fmt.Printf("  Name: %s\n", user.Name)
			fmt.Printf("  Password Hash: %s\n", user.PasswordHash[:50]+"...")
			fmt.Printf("  Active: %t\n", user.IsActive)
			fmt.Println()
		}
	}
}
