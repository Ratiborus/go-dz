package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func NewBin(id, name string, private bool) *Bin {
	return &Bin{id, private, time.Now(), name}
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{bins}
}
