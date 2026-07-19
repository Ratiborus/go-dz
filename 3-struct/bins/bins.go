package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	Record   map[string]any `json:"record"`
	Metadata Metadata       `json:"metadata"`
}

type Metadata struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type BinList struct {
	Bins []Bin
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{bins}
}

func (list *BinList) FindById(id string) (*Bin, error) {
	for _, b := range list.Bins {
		if b.Metadata.Id == id {
			return &b, nil
		}
	}
	return nil, fmt.Errorf("Bin is not found by id: %s", id)
}

func (list *BinList) DeleteById(id string) (*Bin, error) {
	for i, b := range list.Bins {
		if b.Metadata.Id == id {
			list.Bins = append(list.Bins[:i], list.Bins[i+1:]...)
			return &b, nil
		}
	}
	return nil, fmt.Errorf("Bin does not exist by id id: %s", id)
}
