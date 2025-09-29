package model

// Bin — сущность "контейнер"
type Bin struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

// BinList — агрегат из бинов
type BinList struct {
	Bins []Bin `json:"bins"`
}
