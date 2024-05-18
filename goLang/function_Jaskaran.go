package main

import (
	"fmt"
	"math"
)

var inflationRate = 2.5

func function() {
	var investmentAmount, expectedReturnRate, years float64
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	fmt.Printf("Future Value: %f \n Future Real Value: %f", futureValue, futureRealValue)
}
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	//if you want to directly give return in brackets not below then dont use := use = directly
	// fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// frv := fv / math.Pow(1+inflationRate/100, years)
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return
}
