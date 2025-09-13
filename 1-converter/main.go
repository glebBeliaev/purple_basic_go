package main

import (
	"fmt"
	"strings"
)

const UsdToEuro float64 = 0.85
const UsdToRub float64 = 85.25

func main() {
	for {
		firstCurrency := getFirstCurrency()
		amount := getAmount()
		secondaryCurrency := getSecondCurrency(firstCurrency)
		result := convert(firstCurrency, secondaryCurrency, amount)
		fmt.Printf("Результат конвертации: %.2f %s = %.2f %s\n", amount, firstCurrency, result, secondaryCurrency)

		if checkRepeat() == "N" {
			break
		} else {
			continue
		}
	}
}

func getFirstCurrency() string {
	var firstCurrency string
	for {
		fmt.Print("Какую валюту хотите конвертировать? (USD, EUR, RUB): ")
		fmt.Scan(&firstCurrency)
		firstCurrency = strings.ToUpper(firstCurrency)
		if firstCurrency != "USD" && firstCurrency != "EUR" && firstCurrency != "RUB" {
			fmt.Println("Недопустимая валюта, повторите ввод")
			continue
		}
		break
	}
	fmt.Println("Валюта", firstCurrency, "выбрана")
	return firstCurrency
}

func getSecondCurrency(firstCurrency string) string {
	const WhatToConvert = "В какую валюту конвертировать?"
	var secondCurrency string
	switch firstCurrency {
	case "USD":
		for {
			fmt.Printf("%s (EUR, RUB): ", WhatToConvert)
			fmt.Scan(&secondCurrency)
			secondCurrency = strings.ToUpper(secondCurrency)
			if secondCurrency != "EUR" && secondCurrency != "RUB" {
				continue
			}
			break
		}
		return secondCurrency
	case "EUR":
		for {
			fmt.Printf("%s (USD, RUB): ", WhatToConvert)
			fmt.Scan(&secondCurrency)
			secondCurrency = strings.ToUpper(secondCurrency)
			if secondCurrency != "USD" && secondCurrency != "RUB" {
				continue
			}
			break
		}
		return secondCurrency
	default:
		for {
			fmt.Printf("%s (USD, EUR): ", WhatToConvert)
			fmt.Scan(&secondCurrency)
			secondCurrency = strings.ToUpper(secondCurrency)
			if secondCurrency != "USD" && secondCurrency != "EUR" {
				continue
			}
			break
		}
		return secondCurrency
	}
}

func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Недопустимая сумма, повторите ввод")
			continue
		}
		break
	}
	return amount
}

func convert(firstCurrency, secondCurrency string, amount float64) float64 {
	switch firstCurrency {
	case "USD":
		switch secondCurrency {
		case "EUR":
			return amount * UsdToEuro
		case "RUB":
			return amount * UsdToRub
		}
	case "EUR":
		switch secondCurrency {
		case "USD":
			return amount / UsdToEuro
		case "RUB":
			return amount * (UsdToEuro / UsdToRub)
		}
	case "RUB":
		switch secondCurrency {
		case "USD":
			return amount / UsdToRub
		case "EUR":
			return amount / (UsdToRub / UsdToEuro)
		}
	}
	return 0
}

func checkRepeat() string {
	var answer string
	for {
		fmt.Println("Повторить? (Y/N): ")
		fmt.Scan(&answer)
		answer = strings.ToUpper(answer)
		if answer != "Y" && answer != "N" {
			fmt.Println("Недопустимый ответ, повторите ввод")
			continue
		}
		break
	}
	return answer
}
