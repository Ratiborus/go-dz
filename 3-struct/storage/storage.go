package storage

import "app/bin/bins"

type Storage interface {
	Read() (*bins.BinList, error)
	Write(bin *bins.Bin) error
	WriteList(list *bins.BinList) error
}
