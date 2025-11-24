package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

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
}
