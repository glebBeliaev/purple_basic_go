package model

// Bin — сущность "контейнер"
type Bin struct {
	ID      string `json:"id,omitempty"` // id jsonbin'а
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

// BinList — агрегат из бинов
type BinList struct {
	Bins []Bin `json:"bins"`
}
