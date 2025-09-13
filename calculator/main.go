package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTpower = 2

func main() {
	for {
		userHeight, userWeight := getUserInput()
		IMT, err := calcalateIMT(userHeight, userWeight)
		if err != nil {
			fmt.Println("Недопустимые значения, повторите ввод")
			continue
		}
		fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
		outputResult(IMT)
		if !checkRepeat() {
			break
		}
	}

}

func calcalateIMT(height, weight float64) (float64, error) {
	if weight == 0 || height == 0 {
		return 0, errors.New("Invalid input")
	}
	return weight / math.Pow(height/100, IMTpower), nil
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64

	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func outputResult(IMT float64) {
	switch {
	case IMT < 16:
		fmt.Println(" - выраженный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println(" - недостаточная масса тела")
	case IMT < 25:
		fmt.Println(" - нормальная масса тела")
	case IMT < 30:
		fmt.Println(" - избыточная масса тела")
	default:
		fmt.Println(" - ожирение")
	}
}

func checkRepeat() bool {
	var answer string
	fmt.Print("Повторить? (Y/N): ")
	fmt.Scan(&answer)
	return answer == "y" || answer == "Y"
}
