package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// send pointer with & before var name
	fmt.Print("Revenus: ")
	fmt.Scan(&revenue)

	// send pointer with & before var name
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	// send pointer with & before var name
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
