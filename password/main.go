package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[4]int) {
	for index, value := range *a {
		(*a)[len(a)-index-1] = value
	}
}
