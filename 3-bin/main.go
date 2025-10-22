package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	cl := api.NewClient()
	bs := bins.NewService(st, dataFile, cl)
	return &App{fileManager: fm, storage: st, binService: bs}
}

func main() {
	_ = godotenv.Load()

	dataFile := os.Getenv("BINS_DATA_FILE")
	if dataFile == "" {
		dataFile = "password/data.json"
	}
	app := NewApp(dataFile)

	if err := RunCLI(context.Background(), app.binService, os.Args[1:], os.Stdout); err != nil {
		log.Fatal(err)
	}
}

// RunCLI — вынос логики CLI для тестов.
func RunCLI(ctx context.Context, svc interfaces.BinService, args []string, out io.Writer) error {
	var (
		flCreate = flag.NewFlagSet("create", flag.ContinueOnError)
		flGet    = flag.NewFlagSet("get", flag.ContinueOnError)
		flUpdate = flag.NewFlagSet("update", flag.ContinueOnError)
		flDelete = flag.NewFlagSet("delete", flag.ContinueOnError)
		flList   = flag.NewFlagSet("list", flag.ContinueOnError)
	)

	// Общие флаги
	createFile := flCreate.String("file", "", "path to JSON file with record data")
	createName := flCreate.String("name", "", "bin name")
	createPriv := flCreate.Bool("private", false, "mark bin as private (local only)")

	getID := flGet.String("id", "", "remote bin id")
	updateID := flUpdate.String("id", "", "remote bin id")
	updateFile := flUpdate.String("file", "", "path to JSON file with record data")
	deleteID := flDelete.String("id", "", "remote bin id")

	// Разбор верхнего уровня: первая позиция — команда
	if len(args) == 0 {
		usage(out)
		return nil
	}
	cmd := args[0]
	switch cmd {
	case "--create", "create":
		if err := flCreate.Parse(args[1:]); err != nil {
			return err
		}
		if *createFile == "" || *createName == "" {
			return fmt.Errorf("--create requires --file and --name")
		}
		data, err := readJSON(*createFile)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()
		id, err := svc.CreateRemote(ctx, *createName, *createPriv, data)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(out, "created: id=%s name=%s private=%v\n", id, *createName, *createPriv)
		return nil

	case "--get", "get":
		if err := flGet.Parse(args[1:]); err != nil {
			return err
		}
		if *getID == "" {
			return fmt.Errorf("--get requires --id")
		}
		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()
		bin, err := svc.GetRemote(ctx, *getID)
		if err != nil {
			return err
		}
		b, _ := json.MarshalIndent(bin, "", "  ")
		_, _ = out.Write(b)
		_, _ = out.Write([]byte("\n"))
		return nil

	case "--update", "update":
		if err := flUpdate.Parse(args[1:]); err != nil {
			return err
		}
		if *updateID == "" || *updateFile == "" {
			return fmt.Errorf("--update requires --id and --file")
		}
		data, err := readJSON(*updateFile)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()
		if err := svc.UpdateRemote(ctx, *updateID, data); err != nil {
			return err
		}
		_, _ = fmt.Fprintf(out, "updated: %s\n", *updateID)
		return nil

	case "--delete", "delete":
		if err := flDelete.Parse(args[1:]); err != nil {
			return err
		}
		if *deleteID == "" {
			return fmt.Errorf("--delete requires --id")
		}
		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()
		if err := svc.DeleteRemote(ctx, *deleteID); err != nil {
			return err
		}
		_, _ = fmt.Fprintf(out, "deleted: %s\n", *deleteID)
		return nil

	case "--list", "list":
		if err := flList.Parse(args[1:]); err != nil {
			return err
		}
		list := svc.GetBins()
		if len(list) == 0 {
			_, _ = fmt.Fprintln(out, "No local bins yet")
			return nil
		}
		for _, b := range list {
			_, _ = fmt.Fprintf(out, "%s  %s  private=%v\n", b.ID, b.Name, b.Private)
		}
		return nil

	default:
		usage(out)
		return nil
	}
}

func usage(out io.Writer) {
	fmt.Fprint(out, `Usage:
  go run ./3-bin create  --file=record.json --name=mybin [--private]
  go run ./3-bin get     --id=<id>
  go run ./3-bin update  --id=<id> --file=record.json
  go run ./3-bin delete  --id=<id>
  go run ./3-bin list
`)
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
