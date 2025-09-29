package interfaces

import "purple_basic_go/3-bin/model"

// FileManager — абстракция над чтением/записью байтов в файл
type FileManager interface {
	Read(name string) ([]byte, error)
	Write(content []byte, name string) error
}

// BinStorage — абстракция над сохранением/загрузкой списка бинов
type BinStorage interface {
	SaveBins(binList *model.BinList, filename string) error
	LoadBins(filename string) (*model.BinList, error)
}

// BinService — операции над доменной логикой "Bin"
type BinService interface {
	CreateBin(name string, private bool) model.Bin
	AddBin(bin model.Bin) error
	GetBins() []model.Bin
}
