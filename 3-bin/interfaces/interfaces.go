package interfaces

import (
	"context"

	"purple_basic_go/3-bin/model"
)

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
	// Локальные операции
	CreateBin(name string, private bool) model.Bin
	AddBin(bin model.Bin) error
	GetBins() []model.Bin
	RemoveByID(id string) error

	// Удалённые операции (jsonbin)
	CreateRemote(ctx context.Context, name string, private bool, data map[string]any) (string, error)
	GetRemote(ctx context.Context, id string) (*model.Bin, error)
	UpdateRemote(ctx context.Context, id string, data map[string]any) error
	DeleteRemote(ctx context.Context, id string) error
}
