package storage

import (
	"app/bin/bins"
	"app/bin/file"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const storage string = "storage.json"

type LocalStorage struct {
	fileName string
}

func NewLocalStorage() Storage {
	return &LocalStorage{fileName: storage}
}

func (s *LocalStorage) Read() (*bins.BinList, error) {
	b, err := os.ReadFile(s.fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("Storage file is not found %w", err)
	} else {
		var list bins.BinList
		err := json.Unmarshal(b, &list)
		if err != nil {
			return nil, fmt.Errorf("Can not deserialize exists Bin list data from storage %w", err)
		}
		return &list, nil
	}
}

func (s *LocalStorage) Write(bin *bins.Bin) error {
	list, err := s.Read()
	if errors.Is(err, os.ErrNotExist) {
		created := bins.BinList{
			Bins: []bins.Bin{*bin},
		}
		return s.saveBinList(&created)
	} else if err != nil {
		return fmt.Errorf("Can not save Bin to existed storage %w", err)
	} else {
		list.Bins = append(list.Bins, *bin)
		return s.saveBinList(list)
	}
}

func (s *LocalStorage) saveBinList(list *bins.BinList) error {
	marshal, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("Can not marshal Bin list %w", err)
	}
	err = file.WriteFile(s.fileName, marshal)
	if err != nil {
		return fmt.Errorf("Can not write Bin list to storage %w", err)
	}
	return nil
}
