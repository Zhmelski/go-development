package writer

import (
	"fmt"
	"os"

	"homework_05/task_01/models"
)

// 'WriteOutput' writes the results to the out.txt file
func WriteOutput(fileName string, vacations []models.EmployeeVacation) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, vacation := range vacations {
		_, err = fmt.Fprintf(file, "%s: %d\n", vacation.Name, vacation.TotalDays)
		if err != nil {
			return err
		}
	}

	return nil
}

// 'WriteErrors' writes errors to the errors.txt file.
func WriteErrors(fileName string, errors []models.ErrorRecord) error {
	if len(errors) == 0 {
		return nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, e := range errors {
		_, err2 := fmt.Fprintf(file, "%d %s\n", e.LineNumber, e.Content)
		if err2 != nil {
			return err2
		}
	}

	return nil
}
