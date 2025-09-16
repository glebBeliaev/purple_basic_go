package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func (b *Bin) NewBin(name string, private bool) {
	b.id = b.generateId()
	b.private = private
	b.createdAt = time.Now()
	b.name = name
}

func (b *Bin) generateId() string {
	var letterRuns = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, 8)
	for i := range res {
		res[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	return string(res)
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
	name := promtData("Введите название: ")
	privateReq := promtData("Приватный бин? (Y/N)")
	private := false
	if privateReq == "y" || privateReq == "Y" {
		private = true
	}
	bin.NewBin(name, private)
	binList.AddBin(bin)
	fmt.Println(binList.bins)
}

func promtData(promt string) string {
	var data string
	fmt.Print(promt)
	fmt.Scanln(&data)
	return data
}
