package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := selectOperation()
	raw := getNumbers()
	numbers := createSlice(raw)
	if len(numbers) == 0 {
		fmt.Println("Нет чисел для расчёта.")
		return
	}

	// Карта операций: каждая функция замыкает numbers
	ops := map[string]func(){
		"AVG": func() { fmt.Println("Среднее:", calculateAverage(numbers)) },
		"SUM": func() { fmt.Println("Сумма:", calculateSum(numbers)) },
		"MED": func() { fmt.Println("Медиана:", calculateMedian(numbers)) },
	}

	if run, ok := ops[strings.ToUpper(operation)]; ok {
		run()
	} else {
		fmt.Println("Неизвестная операция:", operation)
	}
}

func selectOperation() string {
	var operation string
	fmt.Print("Выберите операцию (AVG - среднее, SUM - сумму, MED - медиану): ")
	fmt.Scan(&operation)
	return operation
}

func getNumbers() string {
	var numbers string
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&numbers)
	return numbers
}

func createSlice(num string) []int {
	parts := strings.Split(num, ",")
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			fmt.Println("Ошибка конвертации, пропущено:", p)
			continue
		}
		out = append(out, n)
	}
	return out
}

func calculateAverage(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return float64(sum) / float64(len(numbers))
}

func calculateSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func calculateMedian(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	nums := append([]int(nil), numbers...) // копия
	sort.Ints(nums)
	n := len(nums)
	if n%2 == 0 {
		mid := n / 2
		return float64(nums[mid-1]+nums[mid]) / 2
	}
	return float64(nums[n/2])
}
