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

func New() Storage {
	return &LocalStorage{fileName: storage}
}

func (s *LocalStorage) Read() (*bins.BinList, error) {
	b, err := file.ReadFile(s.fileName)
	if os.IsNotExist(err) {
		file.WriteFile(s.fileName, nil)
		return bins.NewBinList(nil), nil
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
		return s.WriteList(&created)
	} else if err != nil {
		return fmt.Errorf("Can not save Bin to existed storage %w", err)
	} else {
		list.Bins = append(list.Bins, *bin)
		return s.WriteList(list)
	}
}

func (s *LocalStorage) WriteList(list *bins.BinList) error {
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
