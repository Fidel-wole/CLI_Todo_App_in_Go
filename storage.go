package main

import (
	"encoding/json"
	"os"
)

type Storage[t any] struct {
	FileName string
}

func NewStorage[t any](fileName string) *Storage[t] {
	return &Storage[t]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "   ")
	if err != nil{
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
fileData, err := os.ReadFile(s.FileName)
if err != nil {
	return err
}
return json.Unmarshal(fileData, data)
}