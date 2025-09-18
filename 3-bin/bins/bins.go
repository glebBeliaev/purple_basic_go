package bins

import (
	"encoding/json"
	"math/rand"
	"purple_basic_go/password/files"
	"time"

	"github.com/fatih/color"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

func (b *Bin) NewBin(name string, private bool) {
	b.Id = b.generateId()
	b.Private = private
	b.CreatedAt = time.Now()
	b.Name = name
}

func (b *Bin) generateId() string {
	var letterRuns = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, 8)
	for i := range res {
		res[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	return string(res)
}

type BinList struct {
	Bins []Bin
}

func (b *BinList) AddBin(bin Bin) {
	b.Bins = append(b.Bins, bin)
}

func NewBins() *BinList {
	file, err := files.ReadFile("password/data.json")
	if err != nil {
		return &BinList{
			Bins: []Bin{},
		}
	}
	var binList BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		color.Red("Не удалось разобрать файл JSON")
		return &BinList{
			Bins: []Bin{},
		}
	}
	return &binList
}

func (bin *Bin) ToBytes() ([]byte, error) {
	file, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}
	return file, nil
}
