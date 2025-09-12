package main

import "fmt"

func main() {
	const UsdToEuro float64 = 0.85
	const UsdToRub float64 = 85.25

	getUserInput()
	converter("USD", "EUR")
}

func getUserInput() float64 {
	var currency string
	var amount float64

	fmt.Print("Введите валюту: ")
	fmt.Scan(&currency)
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	return amount
}

func converter(a, b string) float64 {
	return 0
}
