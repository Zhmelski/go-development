package parser

import (
	"encoding/csv"
	"os"
	"strings"
	"time"

	"HomeWork_05/Task_01/models"
)

const dateLayout = "1/2/2006"

// 'ParseCSV' reads a CSV file and returns valid entries and errors.
func ParseCSV(fileName string) ([]models.VacationRecord, []models.ErrorRecord, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	var validRecords []models.VacationRecord
	var errorRecords []models.ErrorRecord

	for i, record := range records {
		lineNumber := i + 1

		// Checking the number of columns in the record
		if len(record) != 3 {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNumber,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		// Check the name format (must be exactly 2 words)
		nameParts := strings.Fields(strings.TrimSpace(record[0]))
		if len(nameParts) != 2 {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNumber,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		// Parsing the dates
		startDate, err1 := time.Parse(dateLayout, strings.TrimSpace(record[1]))
		endDate, err2 := time.Parse(dateLayout, strings.TrimSpace(record[2]))

		if err1 != nil || err2 != nil {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNumber,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		// Check that the end date is >= the start date.
		if endDate.Before(startDate) {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNumber,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		validRecords = append(validRecords, models.VacationRecord{
			EmployeeName: strings.TrimSpace(record[0]),
			StartDate:    startDate,
			EndDate:      endDate,
			LineNumber:   lineNumber,
		})
	}

	return validRecords, errorRecords, nil
}
