package worker

import (
	"app/bin/api"
	"app/bin/args"
	"app/bin/config"
	"app/bin/file"
	"app/bin/storage"
	"fmt"
)

var stor storage.Storage = storage.New()
var remoteApi api.Api = *api.NewApi(config.NewConfig())

func Make(data *args.ArgsPureData) error {
	switch data.Operation {
	case args.Create:
		return CreateBin(data)
	case args.Update:
		return UpdateBin(data)
	case args.Delete:
		return DeleteBin(data)
	case args.Get:
		return GetBin(data)
	case args.List:
		return ListBin(data)
	default:
		return nil
	}
}

func CreateBin(data *args.ArgsPureData) error {
	toStore, err := file.ReadFile(data.FileName)
	if err != nil {
		return err
	}
	bin, err := remoteApi.CreateBin(toStore, data.Name)
	if err != nil {
		return err
	}
	err = stor.Write(bin)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBin(data *args.ArgsPureData) error {
	source, err := file.ReadFile(data.FileName)
	if err != nil {
		return err
	}
	localList, err := stor.Read()
	if err != nil {
		return err
	}
	target, err := localList.DeleteById(data.Id)
	if err != nil {
		return err
	}
	err = stor.WriteList(localList)
	if err != nil {
		return err
	}
	remoteBin, err := remoteApi.UpdateBin(source, data.Id)
	if err != nil {
		return err
	}
	target.Record = remoteBin.Record
	stor.Write(target)
	return nil
}

func DeleteBin(data *args.ArgsPureData) error {
	err := remoteApi.DeleteBin(data.Id)
	if err != nil {
		return err
	}
	localList, err := stor.Read()
	if err != nil {
		return err
	}
	_, err = localList.DeleteById(data.Id)
	if err != nil {
		return err
	}
	err = stor.WriteList(localList)
	if err != nil {
		return err
	}
	fmt.Println("Bin was successfully deleted from remote api and local storage.")
	return nil
}

func GetBin(data *args.ArgsPureData) error {
	localList, err := stor.Read()
	if err != nil {
		return err
	}
	bin, err := localList.FindById(data.Id)
	if err != nil {
		return err
	}
	fmt.Printf("Bean with id: %s was found: %+v\n", data.Id, bin)
	return nil
}

func ListBin(data *args.ArgsPureData) error {
	localList, err := stor.Read()
	if err != nil {
		return err
	}
	for _, bin := range localList.Bins {
		fmt.Printf("Bin with Id: %s and name: %s\n", bin.Metadata.Id, bin.Metadata.Name)
	}
	return nil
}
