package main

import (
	"fmt"
	"math"
)

const IMTpower = 2

func main() {

	userHeight, userWeight := getUserInput()
	IMT := calcalateIMT(userHeight, userWeight)
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)

}

func calcalateIMT(height, weight float64) float64 {
	return weight / math.Pow(height/100, IMTpower)
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64

	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
