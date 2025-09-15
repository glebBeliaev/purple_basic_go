package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	rates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"RUB": 85.25,
	}

	for {
		from := promptCurrency("Какую валюту хотите конвертировать", "", &rates)
		amount := promptAmount()

		to := promptCurrency("В какую валюту конвертировать", from, &rates)
		result, ok := convert(from, to, amount, &rates)
		if !ok {
			fmt.Println("Неизвестная валюта, проверьте ввод")
			continue
		}

		fmt.Printf("Результат конвертации: %.2f %s = %.2f %s\n", amount, from, result, to)

		if !promptYesNo("Повторить? (Y/N): ") {
			break
		}
	}
}

func promptCurrency(question, exclude string, rates *map[string]float64) string {
	opts := availableCurrencies(exclude, rates)
	optStr := strings.Join(opts, ", ")

	r := *rates

	for {
		fmt.Printf("%s? (%s): ", question, optStr)
		var cur string
		fmt.Scan(&cur)
		cur = strings.ToUpper(strings.TrimSpace(cur))

		if _, ok := r[cur]; !ok {
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

func convert(from, to string, amount float64, rates *map[string]float64) (float64, bool) {
	if rates == nil {
		return 0, false
	}
	r := *rates
	f, ok1 := r[from]
	t, ok2 := r[to]
	if !ok1 || !ok2 {
		return 0, false
	}
	return amount * t / f, true
}

func availableCurrencies(exclude string, rates *map[string]float64) []string {
	if rates == nil {
		return nil
	}
	r := *rates
	keys := make([]string, 0, len(r))
	for k := range r {
		if k != exclude {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	return keys
}
