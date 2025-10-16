package bins

import (
	"context"
	"errors"

	"purple_basic_go/3-bin/api"
	"purple_basic_go/3-bin/interfaces"
	"purple_basic_go/3-bin/model"
)

// Service реализует бизнес-логику и соответствует interfaces.BinService.
type Service struct {
	store     interfaces.BinStorage
	filename  string
	data      *model.BinList
	apiClient *api.Client
}

// NewService создаёт сервис, пытаясь загрузить текущее состояние из хранилища.
// Если файла нет или он пустой — работаем с пустым списком.
func NewService(store interfaces.BinStorage, filename string, apiClient *api.Client) *Service {
	list, _ := store.LoadBins(filename) // ошибку на старте проглатываем — позволяем начать "с нуля"
	if list == nil {
		list = &model.BinList{Bins: []model.Bin{}}
	}
	return &Service{
		store:     store,
		filename:  filename,
		data:      list,
		apiClient: apiClient,
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

// RemoveByID — удаляет локально бин по id и сохраняет файл.
func (s *Service) RemoveByID(id string) error {
	if s == nil || s.data == nil {
		return errors.New("service not initialized")
	}
	out := s.data.Bins[:0]
	found := false
	for _, b := range s.data.Bins {
		if b.ID == id {
			found = true
			continue
		}
		out = append(out, b)
	}
	if !found {
		return errors.New("bin not found")
	}
	s.data.Bins = out
	return s.store.SaveBins(s.data, s.filename)
}

// ---------- Удалённые операции (jsonbin) ----------

func (s *Service) CreateRemote(ctx context.Context, name string, private bool, data map[string]any) (string, error) {
	if s.apiClient == nil {
		return "", errors.New("api client is nil")
	}
	id, err := s.apiClient.CreateBin(ctx, data, name)
	if err != nil {
		return "", err
	}
	// сохраним локально привязку (id, name, private)
	_ = s.AddBin(model.Bin{ID: id, Name: name, Private: private})
	return id, nil
}

func (s *Service) GetRemote(ctx context.Context, id string) (*model.Bin, error) {
	if s.apiClient == nil {
		return nil, errors.New("api client is nil")
	}
	return s.apiClient.GetBin(ctx, id)
}

func (s *Service) UpdateRemote(ctx context.Context, id string, data map[string]any) error {
	if s.apiClient == nil {
		return errors.New("api client is nil")
	}
	return s.apiClient.UpdateBin(ctx, id, data)
}

func (s *Service) DeleteRemote(ctx context.Context, id string) error {
	if s.apiClient == nil {
		return errors.New("api client is nil")
	}
	if err := s.apiClient.DeleteBin(ctx, id); err != nil {
		return err
	}
	// удалим локально
	return s.RemoveByID(id)
}
