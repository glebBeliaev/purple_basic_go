package cloud

type CloudDb struct {
	Url string
}

func NewJsonDb(url string) *CloudDb {
	return &CloudDb{
		Url: url,
	}
}
func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *CloudDb) Write(content []byte) {

}
