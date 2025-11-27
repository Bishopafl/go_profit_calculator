package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")

<<<<<<< HEAD
	var revenueLabel string = "Revenus: "
	var expensesLabel string = "Expenses: "
	var taxRateLabel string = "Tax Rate: "

	revenue = askAndSetValue(revenueLabel)
	expenses = askAndSetValue(expensesLabel)
	taxRate = askAndSetValue(taxRateLabel)

	ebt := calculateEbt(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := determineRatio(ebt, profit)

	outputValue(ebt)
	outputValue(profit)
	outputValue(ratio)

}

func askAndSetValue(v string) (re float64) {
	// send pointer with & before var name
	fmt.Print(v)
	fmt.Scan(&re)
	return re
}

func calculateEbt(revenue, expenses float64) (ebt float64) {
	ebt = revenue - expenses
	return ebt
}

func calculateProfit(ebt, taxRate float64) (profit float64) {
	return ebt * (1 - taxRate/100)
}

func determineRatio(ebt, profit float64) (r float64) {
	return ebt / profit
}

func outputValue(iv float64) {
	fmt.Println(iv)
=======
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

>>>>>>> f3a5229 (Doing some error handling and writing to a file of the revenue, expenses and tax rates. this is much more robust than the previous code that was implemented and it creates a results.txt file that shows the information)
}
