package main

import (
	"fmt"
)

func main() {
	var transactions []float32
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	sum := sum(transactions)
	fmt.Println("Общая сумма транзакций:", sum)
}

func scanTransaction() float32 {
	var amount float32
	fmt.Print("Введите сумму транзакции (0 - выход): ")
	fmt.Scan(&amount)
	return amount
}

func sum(transactions []float32) float32 {
	var sum float32
	for _, transaction := range transactions {
		sum += transaction
	}
	return sum
}
