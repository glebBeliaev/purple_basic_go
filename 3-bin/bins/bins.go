package bins

import (
	"math/rand"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
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
