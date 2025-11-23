package models

import "time"

// 'VacationRecord' represents one vacation record from a CSV
type VacationRecord struct {
	EmployeeName string
	StartDate    time.Time
	EndDate      time.Time
	LineNumber   int
}

// 'EmployeeVacation' contains information about the employee's total vacation time
type EmployeeVacation struct {
	Name      string
	TotalDays int
	LastName  string // to sort by last name
}

// 'ErrorRecord' represents an erroneous row from the CSV
type ErrorRecord struct {
	LineNumber int
	Content    string
}
