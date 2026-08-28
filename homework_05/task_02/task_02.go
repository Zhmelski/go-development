/*
Файл 'sales.csv' содержит информацию о продажах магазина за один день. Каждая строка включает:
- название товара,
- цену за единицу (положительное вещественное число),
- количество проданных единиц (положительное целое число).
Один и тот же товар может встречаться в файле несколько раз.
Ваша задача — обработать данные и сформировать отчёт 'report.md' в формате Markdown, содержащий:
- Сумму продаж по каждому товару;
- Общий доход магазина за день;
- Название товара с наибольшей суммой продаж.
CSV-файл sales.csv имеет следующую структуру (приведено в качестве примера):
Product,Price,Units
"Instant coffee",20.99,2
"Chewing gum",1.99,1
"Instant coffee",20.99,1
"Chewing gum",1.99,5
Структура файла report.md:
# Report for 30 January, 2025
## Total Revenue: 74.91
## Highest Revenue Product: Instant coffee (62.97)
## Products
1. Chewing gum = 11.94
2. Instant coffee = 62.97
Важно – при генерации файла report.md используйте текстовые шаблоны (пакет text/template).
Для работы с CSV-файлом вы можете (но не обязаны) использовать стандартный пакет encoding/csv.
*/

package main

import (
	"fmt"

	"homework_05/task_02/parser"
	"homework_05/task_02/processor"
	"homework_05/task_02/reporter"
)

func main() {
	const (
		inputFile  = "Task_02/DataSets/sales.csv"
		outputFile = "Task_02/report.md"
	)

	// Parsing a CSV file
	sales, errorRecords, err := parser.ParseCSV(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
	if errorRecords != nil {
		fmt.Printf("ErrorRecords: %v\n", errorRecords)
	}

	fmt.Printf("Sales records loaded: %d\n", len(sales))

	// Processing sales records
	report := processor.ProcessSales(sales)

	fmt.Printf("Unique products: %d\n", len(report.Products))
	fmt.Printf("Total Revenue: %.2f\n", report.TotalRevenue)
	fmt.Printf("Highest Revenue Product: %s (%.2f)\n", report.TopProduct, report.TopProductRevenue)

	// Generating a report
	if err := reporter.GenerateReport(outputFile, report); err != nil {
		fmt.Printf("Error generating report: %v\n", err)
	}

	fmt.Printf("\nReport saved in %s\n", outputFile)
}
