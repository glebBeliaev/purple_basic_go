package bins

import (
	"errors"
	"purple_basic_go/3-bin/interfaces"
	"purple_basic_go/3-bin/model"
)

// Service реализует бизнес-логику и соответствует interfaces.BinService.
type Service struct {
	store    interfaces.BinStorage
	filename string
	data     *model.BinList
}

// NewService создаёт сервис, пытаясь загрузить текущее состояние из хранилища.
// Если файла нет или он пустой — работаем с пустым списком.
func NewService(store interfaces.BinStorage, filename string) *Service {
	list, _ := store.LoadBins(filename) // ошибку на старте проглатываем — позволяем начать "с нуля"
	if list == nil {
		list = &model.BinList{Bins: []model.Bin{}}
	}
	return &Service{
		store:    store,
		filename: filename,
		data:     list,
	}
}

// CreateBin — фабричный метод: создаёт Bin, но не добавляет в список.
func (s *Service) CreateBin(name string, private bool) model.Bin {
	return model.Bin{Name: name, Private: private}
}

// AddBin — добавляет Bin в список и сразу сохраняет состояние.
func (s *Service) AddBin(b model.Bin) error {
	if s == nil || s.data == nil {
		return errors.New("service not initialized")
	}
	s.data.Bins = append(s.data.Bins, b)
	return s.store.SaveBins(s.data, s.filename)
}

// GetBins — отдаёт срез (копию) текущих бинов.
func (s *Service) GetBins() []model.Bin {
	if s == nil || s.data == nil {
		return nil
	}
	out := make([]model.Bin, len(s.data.Bins))
	copy(out, s.data.Bins)
	return out
}
