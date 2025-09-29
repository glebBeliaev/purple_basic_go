package interfaces

type FileManager interface {
	Read(name string) ([]byte, error)
	Write(content []byte, name string)
}

type BinStorage interface {
	SaveBins(binList interface{}, filename string) error
	LoadBins(filename string) (interface{}, error)
}

type BinService interface {
	CreateBin(name string, private bool) interface{}
	AddBin(bin interface{})
	GetBins() []interface{}
}
