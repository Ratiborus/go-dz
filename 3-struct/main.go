package main

import (
	"app/bin/args"
	"app/bin/worker"
	"fmt"
)

func main() {
	fmt.Println("____Приложение для работы с  json файлами____")
	argsData, err := args.ReadArgs()
	if err != nil {
		panic(fmt.Errorf("Can not read the args %s", err.Error()))
	}
	err = worker.Make(argsData)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("____Программа завершила свою работу___")
}
