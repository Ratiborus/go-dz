package args

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Arg int

const (
	Create Arg = iota
	Update
	Delete
	Get
	List
	File
	Name
	Id
	NotDefined
)

type ArgsData struct {
	Data map[Arg]any
}

type ArgsPureData struct {
	Operation Arg
	FileName  string
	Name      string
	Id        string
}

func ReadArgs() (*ArgsPureData, error) {
	argsData := read()
	err := argsData.validate()
	if err != nil {
		return nil, err
	}
	return argsData.newArgsPureData(), nil
}

func (a *ArgsData) newArgsPureData() *ArgsPureData {
	return &ArgsPureData{
		Operation: a.getOperation(),
		FileName:  toString(a.Data[File]),
		Name:      toString(a.Data[Name]),
		Id:        toString(a.Data[Id]),
	}
}

func (a *ArgsData) getOperation() Arg {
	if toBool(a.Data[Create]) {
		return Create
	} else if toBool(a.Data[Update]) {
		return Update
	} else if toBool(a.Data[Delete]) {
		return Delete
	} else if toBool(a.Data[Get]) {
		return Get
	} else if toBool(a.Data[List]) {
		return List
	}
	return NotDefined
}

func read() ArgsData {
	argsData := ArgsData{
		Data: make(map[Arg]any),
	}
	create := flag.Bool("create", false, "Creation flag")
	update := flag.Bool("update", false, "Update flag")
	delete := flag.Bool("delete", false, "Delete flag")
	get := flag.Bool("get", false, "Get flag")
	list := flag.Bool("list", false, "List flag")
	file := flag.String("file", "", "File name to read json")
	name := flag.String("name", "", "Json bin name to store in jsonbin API")
	id := flag.String("id", "", "Id of json bin in remote API")
	flag.Parse()
	argsData.Data[Create] = *create
	argsData.Data[Update] = *update
	argsData.Data[Delete] = *delete
	argsData.Data[Get] = *get
	argsData.Data[List] = *list
	argsData.Data[File] = *file
	argsData.Data[Name] = *name
	argsData.Data[Id] = *id
	return argsData
}

func (a *ArgsData) validate() error {
	if a.isOperationNotDefined() {
		return errors.New("Operation is not defined, please define the operation(create, update, delete, get, list)")
	}
	if a.isFileNotDefined() && (toBool(a.Data[Create]) || toBool(a.Data[Update])) {
		return errors.New("File name is not defined, please define file to create or update jsonbin")
	}
	if a.isIdNotDefined() && (toBool(a.Data[Update]) || toBool(a.Data[Delete]) || toBool(a.Data[Get])) {
		return errors.New("Id of json bin is not defined, please define id to update, delete or get file from remote API")
	}
	return nil
}

func (a *ArgsData) isOperationNotDefined() bool {
	return !(toBool(a.Data[Create]) || toBool(a.Data[Update]) || toBool(a.Data[Delete]) || toBool(a.Data[Get]) || toBool(a.Data[List]))
}

func (a *ArgsData) isFileNotDefined() bool {
	return strings.TrimSpace(toString(a.Data[File])) == ""
}

func (a *ArgsData) isIdNotDefined() bool {
	return strings.TrimSpace(toString(a.Data[Id])) == ""
}

func toString(val any) string {
	switch v := val.(type) {
	case string:
		return v
	default:
		panic(fmt.Sprintf("Vareable is not string type %s", val))
	}
}

func toBool(val any) bool {
	switch v := val.(type) {
	case bool:
		return v
	default:
		panic(fmt.Sprintf("Vareable is not bool %s", val))
	}
}
