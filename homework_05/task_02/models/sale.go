package models

// 'SaleRecord' represents a single sale record
type SaleRecord struct {
	Product string
	Price   float64
	Units   int
}

// 'ProductSales' contains aggregated data on the product
type ProductSales struct {
	Name    string
	Revenue float64
}

// 'Report' contains data for the report
type Report struct {
	Date              string
	TotalRevenue      float64
	TopProduct        string
	TopProductRevenue float64
	Products          []ProductSales
}

// 'ErrorRecord' represents an erroneous row from the CSV file
type ErrorRecord struct {
	LineNumber int
	Content    string
}
