package storage

import (
	"encoding/json"
	"log"
	"os"
)

type Storage[T any] struct {
	Filename     string
	stateChanged bool
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		Filename: fileName,
	}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		log.Println(err)
		return err
	}

	if s.stateChanged {
		os.WriteFile(s.Filename, fileData, 0644)
	}

	return nil
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		log.Println(err)
		return err
	}

	return json.Unmarshal(fileData, data)
}

func (s *Storage[T]) StateChanged(changed bool) {
	s.stateChanged = changed
}
