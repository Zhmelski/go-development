package calculator

import (
	"sort"
	"strings"

	"HomeWork_05/Task_01/models"
)

// 'CalculateVacationDays' calculates the total vacation days for each employee.
func CalculateVacationDays(records []models.VacationRecord) []models.EmployeeVacation {

	// Map: employee name -> total days
	vacationMap := make(map[string]int)

	for _, record := range records {
		// Calculate the number of days inclusive
		days := int(record.EndDate.Sub(record.StartDate).Hours()/24) + 1
		vacationMap[record.EmployeeName] += days
	}

	// Convert to slice
	var result []models.EmployeeVacation
	for name, days := range vacationMap {
		// Extracting the last name for sorting
		parts := strings.Fields(name)
		lastName := parts[1] // the last name is the second word

		result = append(result, models.EmployeeVacation{
			Name:      name,
			LastName:  lastName,
			TotalDays: days,
		})
	}

	// Sorting by total days, then by last name
	sort.Slice(result, func(i, j int) bool {
		if result[i].TotalDays == result[j].TotalDays {
			return result[i].LastName < result[j].LastName
		}
		return result[i].TotalDays > result[j].TotalDays
	})

	return result
}
