package files

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(name string) ([]byte, error) {
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
