package storage

import (
	"fmt"

	"github.com/d-jackalope/storage/internal/file"
	"github.com/google/uuid"
)

// Храним данные файлов в мапе
// Причем мы спрятали реализацию, нигде кроме
// этого пакета поля структуры Storage - не видно
// Пример инкапсуляции, изолировали абстракцию от реализации
type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		// При обработке ошибок лучше стараться возвращать
		// остальные аргументы функции пустыми (*file.File = nil)
		return nil, err
	}
	// Сохраняем в хранилище Storage
	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}

	return foundFile, nil
}
