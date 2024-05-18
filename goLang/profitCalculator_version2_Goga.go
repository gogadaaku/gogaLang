package main

import "fmt"

func profitCalcVersion2() {
	// var revenue, expenses, taxRate, earningsBeforeTax, earningsAfterTax float64
	revenue := revenueFunction()
	expenses := expensesFunction()
	taxRate := taxRateFunction()
	earningsBeforeTax := earningsBeforeTaxFunction(revenue, expenses)
	earningsAfterTax := earningsAfterTaxFunction(earningsBeforeTax, taxRate)
	ratio := ratioCalculator(earningsAfterTax, earningsBeforeTax)
	displayFunction(earningsBeforeTax, earningsAfterTax, ratio)
}
func revenueFunction() (revenue float64) {

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	return
}
func expensesFunction() (expenses float64) {
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	return
}
func taxRateFunction() (taxRate float64) {
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	return
}
func earningsBeforeTaxFunction(revenue, expenses float64) float64 {
	return (revenue - expenses)
}
func earningsAfterTaxFunction(earningsBeforeTax, taxRate float64) float64 {
	return (earningsBeforeTax - (earningsBeforeTax * (taxRate / 100)))
}
func ratioCalculator(earningsAfterTax, earningsBeforeTax float64) float64 {
	return (earningsAfterTax / earningsBeforeTax)
}
func displayFunction(earningsBeforeTax, earningsAfterTax, ratio float64) {
	fmt.Printf("Earning Before Tax: %f \n Earning After Tax:%f \n Ratio: %f", earningsBeforeTax, earningsAfterTax, ratio)
}
