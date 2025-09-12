package main

import "fmt"

func main() {
	const UsdToEuro float64 = 0.85
	const UsdToRub float64 = 85.25

	EroToRub := UsdToRub / UsdToEuro
	fmt.Println(EroToRub)
}
