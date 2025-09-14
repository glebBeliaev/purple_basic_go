package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	operation := selectOperation()
	numbers := getNumbers()
	numbersSlice := createSlice(numbers)
	switch operation {
	case "AVG":
		fmt.Println("Среднее:", calculateAverage(numbersSlice))
	case "SUM":
		fmt.Println("Сумма:", calculateSum(numbersSlice))
	case "MED":
		fmt.Println("Медиана:", calculateMedian(numbersSlice))
	}

}

func selectOperation() string {
	var operation string
	fmt.Print("Выберите операцию (AVG - среднее, SUM - сумму, MED - медиану):")
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
	split := strings.Split(num, ",")

	var numbers []int

	for _, s := range split {
		s, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			fmt.Println("Ошибка конвертации:", err)
			continue
		}
		numbers = append(numbers, s)
	}
	return numbers
}

func calculateAverage(numbers []int) float64 {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return float64(sum) / float64(len(numbers))
}

func calculateSum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func calculateMedian(numbers []int) float64 {
	if len(numbers)%2 == 0 {
		return float64(numbers[len(numbers)/2]+numbers[len(numbers)/2-1]) / 2
	}
	return float64(numbers[len(numbers)/2])
}
