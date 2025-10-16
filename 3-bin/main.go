package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"purple_basic_go/3-bin/api"
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

func NewApp(dataFile string) *App {
	fm := file.LocalFileManager{}
	st := storage.JSONStorage{FileManager: fm}
	cl := api.NewClient() // клиент jsonbin
	bs := bins.NewService(st, dataFile, cl)

	return &App{
		fileManager: fm,
		storage:     st,
		binService:  bs,
	}
}

func main() {
	_ = godotenv.Load() // не критично, если .env нет — возьмём env из системы

	const dataFile = "password/data.json"
	app := NewApp(dataFile)

	// CLI-флаги
	var (
		flCreate = flag.Bool("create", false, "create remote bin from --file and --name (stores id locally)")
		flGet    = flag.Bool("get", false, "get remote bin by --id")
		flUpdate = flag.Bool("update", false, "update remote bin by --id from --file")
		flDelete = flag.Bool("delete", false, "delete remote bin by --id (also removes locally)")
		flList   = flag.Bool("list", false, "list local bins (id, name, private)")

		flFile    = flag.String("file", "", "path to JSON file with record data")
		flName    = flag.String("name", "", "bin name (for create)")
		flID      = flag.String("id", "", "remote bin id")
		flPrivate = flag.Bool("private", false, "mark bin as private (local flag only)")
	)
	flag.Parse()

	// Ровно одна команда
	cmdCount := 0
	for _, v := range []bool{*flCreate, *flGet, *flUpdate, *flDelete, *flList} {
		if v {
			cmdCount++
		}
	}
	if cmdCount != 1 {
		usage()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	switch {
	case *flList:
		list := app.binService.GetBins()
		if len(list) == 0 {
			fmt.Println("No local bins yet")
			return
		}
		for _, b := range list {
			fmt.Printf("%s  %s  private=%v\n", b.ID, b.Name, b.Private)
		}

	case *flCreate:
		require(*flFile != "" && *flName != "", "--create requires --file and --name")
		data, err := readJSON(*flFile)
		if err != nil {
			log.Fatal(err)
		}
		id, err := app.binService.CreateRemote(ctx, *flName, *flPrivate, data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created: id=%s name=%s private=%v\n", id, *flName, *flPrivate)

	case *flGet:
		require(*flID != "", "--get requires --id")
		bin, err := app.binService.GetRemote(ctx, *flID)
		if err != nil {
			log.Fatal(err)
		}
		out, _ := json.MarshalIndent(bin, "", "  ")
		fmt.Println(string(out))

	case *flUpdate:
		require(*flID != "" && *flFile != "", "--update requires --id and --file")
		data, err := readJSON(*flFile)
		if err != nil {
			log.Fatal(err)
		}
		if err := app.binService.UpdateRemote(ctx, *flID, data); err != nil {
			log.Fatal(err)
		}
		fmt.Println("updated:", *flID)

	case *flDelete:
		require(*flID != "", "--delete requires --id")
		if err := app.binService.DeleteRemote(ctx, *flID); err != nil {
			log.Fatal(err)
		}
		fmt.Println("deleted:", *flID)
	}
}

func usage() {
	fmt.Print(`Usage:
  go run ./3-bin --create --file=record.json --name=mybin [--private]
  go run ./3-bin --get    --id=<id>
  go run ./3-bin --update --id=<id> --file=record.json
  go run ./3-bin --delete --id=<id>
  go run ./3-bin --list
`)
}

func require(ok bool, msg string) {
	if !ok {
		log.Fatal(msg)
	}
}

func readJSON(path string) (map[string]any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}
