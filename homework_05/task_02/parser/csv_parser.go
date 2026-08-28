package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"HomeWork_05/Task_02/models"
)

// 'ParseCSV' reads and parses a CSV file with sales data.
func ParseCSV(fileName string) ([]models.SaleRecord, []models.ErrorRecord, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all entries
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	var sales []models.SaleRecord
	var errorRecords []models.ErrorRecord

	for i, record := range records {
		lineNum := i + 1 // numbering from 1

		if len(record) != 3 {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNum,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		// Parse the price
		var price float64
		n, err := fmt.Sscanf(record[1], "%f", &price)
		if err != nil || n != 1 || price <= 0 {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNum,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		// Parse the units
		var units int
		n, err = fmt.Sscanf(record[2], "%d", &units)
		if err != nil || n != 1 || units < 0 {
			errorRecords = append(errorRecords, models.ErrorRecord{
				LineNumber: lineNum,
				Content:    strings.Join(record, ","),
			})
			continue
		}

		sales = append(sales, models.SaleRecord{
			Product: record[0],
			Price:   price,
			Units:   units,
		})
	}

	return sales, errorRecords, nil
}
