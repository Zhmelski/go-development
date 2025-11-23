/*
CSV-файл 'data.csv' содержит сведения об отпусках сотрудников компании. Каждая строка представляет собой одну запись и включает
- имя и фамилию сотрудника (в формате: Имя Пробел Фамилия),
- дату начала отпуска,
- дату окончания отпуска.
Формат дат — американский (M/D/YYYY). Пример строк из файла:
Dlsw Fmbnaswsg,1/25/2016,2/3/2016
Iqbidtaw Ndxqjdpsg,6/8/2015,6/19/2015
Lrnywj Fdgyzw,6/8/2015,6/16/2015
Имена сотрудников специально зашифрованы — их не нужно расшифровывать. Полный файл 'data.csv' приложен к заданию.
Ваша программа должна считать файл 'data.csv', вычислить для каждого сотрудника суммарную продолжительность всех его отпусков в днях и сохранить эту информацию в файл 'out.txt', расположив сотрудников в порядке убывания продолжительности отпусков. Вот так должен выглядеть файл 'out.txt' (фрагмент):
Lrnywj Fdgyzw: 10
Dlsw Fmbnaswsg: 5
Iqbidtaw Ndxqjdpsg: 1
Примечания:
1. Если у двух сотрудников одинаковое количество дней отпуска, расположите их в алфавитном порядке по фамилии.
2. Если строка в data.csv содержит ошибку, её нужно записать в файл errors.txt вместе с порядковым номером строки, например:
105 Lrnywj XXX Fdgyzw,6/8/2015,6/16/2015
*/

package main

import (
	"fmt"

	"HomeWork_05/Task_01/calculator"
	"HomeWork_05/Task_01/parser"
	"HomeWork_05/Task_01/writer"
)

func main() {
	const (
		inputFile  = "Task_01/DataSets/data.csv"
		outputFile = "Task_01/out.txt"
		errorFile  = "Task_01/errors.txt"
	)

	// Parsing a CSV file
	records, errors, err := parser.ParseCSV(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	// Calculating the total number of vacation days
	vacations := calculator.CalculateVacationDays(records)

	// Recording the results
	if err := writer.WriteOutput(outputFile, vacations); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
	}

	// Write down errors (if any)
	if err := writer.WriteErrors(errorFile, errors); err != nil {
		fmt.Printf("Error writing errors file: %v\n", err)
	}

	fmt.Printf("Records processed: %d\n", len(records))
	fmt.Printf("Errors: %d\n", len(errors))
	fmt.Printf("Employees: %d\n", len(vacations))
	fmt.Printf("The results are saved in %s\n", outputFile)
	if len(errors) > 0 {
		fmt.Printf("Errors are saved in %s\n", errorFile)
	}

}
