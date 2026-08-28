package processor

import (
	"sort"

	"HomeWork_05/Task_02/models"
)

// 'ProcessSales' processes sales records and generates data for reporting.
func ProcessSales(sales []models.SaleRecord) models.Report {
	// Aggregating sales by product
	revenueMap := make(map[string]float64)

	for _, sale := range sales {
		revenue := sale.Price * float64(sale.Units)
		revenueMap[sale.Product] += revenue
	}

	// Convert to slice and sort
	var products []models.ProductSales
	var totalRevenue float64
	var topProduct string
	var topRevenue float64

	for product, revenue := range revenueMap {
		products = append(products, models.ProductSales{
			Name:    product,
			Revenue: revenue,
		})
		totalRevenue += revenue

		// Finding the product with the highest revenue
		if revenue > topRevenue {
			topProduct = product
			topRevenue = revenue
		}
	}

	// Sort products by name
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})

	return models.Report{
		TotalRevenue:      totalRevenue,
		TopProduct:        topProduct,
		TopProductRevenue: topRevenue,
		Products:          products,
	}
}
