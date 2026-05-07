package main

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func main() {

}

func newBin(id, name string, private bool) *Bin {
	return &Bin{id, private, time.Now(), name}
}

func newBinList(bins []Bin) *BinList {
	return &BinList{bins}
}
