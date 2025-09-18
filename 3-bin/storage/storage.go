package storage

import (
	"encoding/json"
	"purple_basic_go/3-bin/bins"
	"purple_basic_go/3-bin/file"
)

// SaveBins - сохранение списка bin в JSON файл
func SaveBins(binList *bins.BinList, filename string) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}
	file.WriteFile(data, filename)
	return nil
}

// LoadBins - загрузка списка bin из JSON файла
func LoadBins(filename string) (*bins.BinList, error) {
	data, err := file.ReadFile(filename)
	if err != nil {
		return &bins.BinList{Bins: []bins.Bin{}}, err
	}

	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return &bins.BinList{Bins: []bins.Bin{}}, err
	}

	return &binList, nil
}
