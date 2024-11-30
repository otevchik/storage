package storage

import "github.com/otevchik/storage/internal/file"

type Storage struct {
	
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	// newFile, err := file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil, err
	// }
	return file.NewFile(filename, blob)
}