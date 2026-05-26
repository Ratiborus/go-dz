package main

import (
	"app/bin/api"
	"app/bin/bins"
	"app/bin/storage"
	"fmt"
)

func main() {
	var storage storage.Storage = storage.NewLocalStorage()
	storage.Write(bins.NewBin("1", "first-bin", false))
	bin, err := storage.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Bean example %v: %v\n", bin.Bins[0].Id, bin.Bins[0].Name)
	api.Test()
}
