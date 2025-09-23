package files

import (
	"fmt"
	"log"
	"os"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(fileName string) *JsonDb {
	return &JsonDb{
		fileName: fileName,
	}
}
func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.fileName)
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
