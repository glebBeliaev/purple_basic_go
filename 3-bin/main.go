package main

import (
	"fmt"

	"purple_basic_go/3-bin/bins"
	"purple_basic_go/3-bin/file"
	"purple_basic_go/3-bin/interfaces"
	"purple_basic_go/3-bin/storage"
)

type App struct {
	fileManager interfaces.FileManager
	storage     interfaces.BinStorage
	binService  interfaces.BinService
}

// NewApp собирает зависимости.
// dataFile — путь к JSON-файлу, где хранится список (например, "password/data.json").
func NewApp(dataFile string) *App {
	fm := file.LocalFileManager{}
	st := storage.JSONStorage{FileManager: fm}
	bs := bins.NewService(st, dataFile)

	return &App{
		fileManager: fm,
		storage:     st,
		binService:  bs,
	}
}

func main() {
	const dataFile = "password/data.json"

	app := NewApp(dataFile)

	// Пример использования
	b := app.binService.CreateBin("demo", false)
	if err := app.binService.AddBin(b); err != nil {
		panic(err)
	}
	fmt.Println("bins:", app.binService.GetBins())
}
