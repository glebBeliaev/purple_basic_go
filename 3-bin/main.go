package main

import (
	"fmt"
	"purple_basic_go/3-bin/bins"
	"purple_basic_go/3-bin/interfaces"
)

type App struct {
	fileManager interfaces.FileManager
	storage     interfaces.BinService
	binService  interfaces.BinService
}

func NewApp() *App {
	fileManager := &file.FileManager{}
	storage := &storage.BinStorage{FileManager: fileManager}
	binService := &bins.BinService{Storage: storage}

	return &App{
		fileManager: fileManager,
		storage:     storage,
		binService:  binService,
	}
}

func main() {

	binList := bins.BinList{}
	bin := bins.Bin{}
	fmt.Println("Введите данные")
	name := promtData("Введите название: ")
	privateReq := promtData("Приватный бин? (Y/N)")
	private := false
	if privateReq == "y" || privateReq == "Y" {
		private = true
	}
	bin.NewBin(name, private)
	binList.AddBin(bin)
	fmt.Println(binList.Bins)
}

func promtData(promt string) string {
	var data string
	fmt.Print(promt)
	fmt.Scanln(&data)
	return data
}
