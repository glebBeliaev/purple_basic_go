package storage

import (
	"encoding/json"
	"purple_basic_go/3-bin/interfaces"
	"purple_basic_go/3-bin/model"
)

// JSONStorage — хранит BinList в JSON-файле, используя FileManager.
type JSONStorage struct {
	FileManager interfaces.FileManager
}

func (s JSONStorage) SaveBins(binList *model.BinList, filename string) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return err
	}
	return s.FileManager.Write(data, filename)
}

func (s JSONStorage) LoadBins(filename string) (*model.BinList, error) {
	data, err := s.FileManager.Read(filename)
	if err != nil {
		// Возвращаем пустой список и саму ошибку наружу — вызывающий решит, критично это или нет.
		return &model.BinList{Bins: []model.Bin{}}, err
	}
	var out model.BinList
	if err := json.Unmarshal(data, &out); err != nil {
		return &model.BinList{Bins: []model.Bin{}}, err
	}
	return &out, nil
}
