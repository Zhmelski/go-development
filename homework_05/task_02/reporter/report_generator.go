package reporter

import (
	"os"
	"text/template"
	"time"

	"HomeWork_05/Task_02/models"
)

const reportTemplate = `# Report for {{.Date}}

## Total Revenue: {{printf "%.2f" .TotalRevenue}}

## Highest Revenue Product: {{.TopProduct}} ({{printf "%.2f" .TopProductRevenue}})

## Products
{{range $index, $product := .Products}}{{add $index 1}}. {{$product.Name}} = {{printf "%.2f" $product.Revenue}}
{{end}}`

// 'GenerateReport' creates a report in Markdown format.
func GenerateReport(fileName string, report models.Report) error {
	// Add the current date
	report.Date = time.Now().Format("02 January, 2006")

	// Creat function for the template
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	// Parse the template
	tmpl, err := template.New("report").Funcs(funcMap).Parse(reportTemplate)
	if err != nil {
		return err
	}

	// Create the output file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template
	return tmpl.Execute(file, report)
}
