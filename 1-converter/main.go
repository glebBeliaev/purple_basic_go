package main

import (
	"fmt"
	"slices"
	"strings"
)

var perUSD = map[string]float64{
	"USD": 1.0,
	"EUR": 0.85,
	"RUB": 85.25,
}

func main() {
	for {
		from := promptCurrency("Какую валюту хотите конвертировать", "")
		amount := promptAmount()

		to := promptCurrency("В какую валюту конвертировать", from)
		result := convert(from, to, amount)

		fmt.Printf("Результат конвертации: %.2f %s = %.2f %s\n", amount, from, result, to)

		if !promptYesNo("Повторить? (Y/N): ") {
			break
		}
	}
}

func promptCurrency(question, exclude string) string {
	opts := availableCurrencies(exclude)
	optStr := strings.Join(opts, ", ")

	for {
		fmt.Printf("%s? (%s): ", question, optStr)
		var cur string
		fmt.Scan(&cur)
		cur = strings.ToUpper(strings.TrimSpace(cur))

		if _, ok := perUSD[cur]; !ok {
			fmt.Println("Недопустимая валюта, повторите ввод")
			continue
		}
		if exclude != "" && cur == exclude {
			fmt.Println("Нужно выбрать другую валюту, повторите ввод")
			continue
		}
		return cur
	}
}

func promptAmount() float64 {
	for {
		fmt.Print("Введите сумму: ")
		var a float64
		fmt.Scan(&a)
		if a <= 0 {
			fmt.Println("Недопустимая сумма, повторите ввод")
			continue
		}
		return a
	}
}

func promptYesNo(question string) bool {
	for {
		fmt.Print(question)
		var ans string
		fmt.Scan(&ans)
		ans = strings.ToUpper(strings.TrimSpace(ans))
		if ans == "Y" {
			return true
		}
		if ans == "N" {
			return false
		}
		fmt.Println("Недопустимый ответ, введите Y или N")
	}
}

func convert(from, to string, amount float64) float64 {
	return amount * perUSD[to] / perUSD[from]
}

func availableCurrencies(exclude string) []string {
	keys := make([]string, 0, len(perUSD))
	for k := range perUSD {
		if k != exclude {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	return keys
}
