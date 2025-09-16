package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id       string
	private  bool
	createAt time.Time
	name     string
}

func (b *Bin) NewBin() {
	fmt.Scan(&b.id)
	fmt.Scan(&b.private)
	b.createAt = time.Now()
	fmt.Scan(&b.name)
}

type BinList struct {
	bins []Bin
}

func (b *BinList) AddBin(bin Bin) {
	b.bins = append(b.bins, bin)
}
func main() {

	binList := BinList{}
	bin := Bin{}
	fmt.Println("Введите данные")
	bin.NewBin()
	binList.AddBin(bin)
	fmt.Println(binList.bins)
}
