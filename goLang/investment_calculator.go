package main

import (
	"fmt"
	"math"
)

func main() {
	//region---------------------------------------------------------------------------------------------------------------------------------------
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	var investmentAmount1 float64 = 1000
	var expectedReturnRate1 = 5.5
	var years1 float64 = 10

	var futureValue1 = investmentAmount1 * math.Pow(1+expectedReturnRate1/100, years1)
	fmt.Println(futureValue1)
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//we can remove var and use := instead if we are not writing var something float64=

	// futureValue2:=44.44 like this

	//futureValue2 float64:= is not allowed
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	investmentAmount3, years3 := 1000.0, 10.0
	expectedReturnRate3 := 5.5
	futureValue3 := investmentAmount3 * math.Pow(1+expectedReturnRate3/100, years3)
	fmt.Println(futureValue3)
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	var investmentAmount4 float64
	var expectedReturnRate4 float64
	var years4 float64
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount4)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate4)

	fmt.Print("Years: ")
	fmt.Scan(&years4)

	futureValue4 := investmentAmount4 * math.Pow(1+expectedReturnRate4/100, years4)
	fmt.Print("Future Value: ", futureValue4)
	//end region-------------------------------------------------------------------------------------------------------------------------
	profitCalculator()
	printString("hello", "hello2")
	function()
	profitCalcVersion2()
}
func printString(text1, text2 string) {
	fmt.Printf("text1: %s\ntext2: %s\n", text1, text2)
}
