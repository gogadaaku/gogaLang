package main

import "fmt"

func profitCalculator() {
	var revenue, expenses, taxRate float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Println("")
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Println("")

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	fmt.Println("")

	earningsBeforeTax := revenue - expenses
	fmt.Print("Earnings Before tax:", earningsBeforeTax)
	fmt.Println("")

	tax := earningsBeforeTax * taxRate
	fmt.Print("Tax: ", tax)
	fmt.Println("")

	earningsAfterTax := earningsBeforeTax - tax
	fmt.Print("Earning after tax: ", earningsAfterTax)
	fmt.Println("")

	ratio := earningsAfterTax / earningsBeforeTax
	fmt.Print("Ratio: ", ratio)
	fmt.Println(" ")
	fmt.Printf("Earning Before tax: %v\nTax:%v\nEarning After Tax: %v\n", earningsBeforeTax, tax, earningsAfterTax)
	fmt.Printf("Earning Before tax: %f\nTax:%f\nEarning After Tax: %f\n", earningsBeforeTax, tax, earningsAfterTax)
	fmt.Printf("Earning Before tax: %.0f\nTax:%.0f\nEarning After Tax: %.0f\n", earningsBeforeTax, tax, earningsAfterTax)
	fmt.Printf("Earning Before tax: %.0f\nTax:%.0f\nEarning After Tax: %.0f\n", earningsBeforeTax, tax, earningsAfterTax)
	fmt.Printf("Earning Before tax: %.1f\nTax:%.1f\nEarning After Tax: %.1f\n", earningsBeforeTax, tax, earningsAfterTax)
	fmt.Printf(
   `Earning Before tax: %.1f
	Tax:%.1f
	Earning After Tax: %.1f`, earningsBeforeTax, tax, earningsAfterTax)

}
