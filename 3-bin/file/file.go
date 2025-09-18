package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	if filepath.Ext(name) != ".json" {
		return nil, fmt.Errorf("файл %s не имеет расширение .json", name)
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("File created successfully")
}
