package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// LocalFileManager — простая реализация FileManager поверх локальной ФС.
type LocalFileManager struct{}

// Read читает весь файл в []byte.
// Для учебного примера слегка проверяем расширение, чтобы не путаться с другими файлами.
func (LocalFileManager) Read(name string) ([]byte, error) {
	if filepath.Ext(name) != ".json" {
		return nil, fmt.Errorf("expected .json file, got: %s", name)
	}
	return os.ReadFile(name)
}

// Write пишет []byte в файл с правами 0644.
func (LocalFileManager) Write(content []byte, name string) error {
	return os.WriteFile(name, content, 0o644)
}
