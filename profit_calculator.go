package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	expenses, err2 := getUserInput("Expenses: ")

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	taxRate, err3 := getUserInput("Tax Rate: ")

	if err3 != nil {
		fmt.Println(err3)
		return
	}

	// calculate financials
	ebt, profit, ratio := calulateFinancials(revenue, expenses, taxRate)

	writeFinancialToFile(ebt, profit, ratio)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}

// Calculate financials
func calulateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func writeFinancialToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)

}
