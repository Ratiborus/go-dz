package api

import (
	"app/bin/config"
	"fmt"
)

func Test() {
	conf := config.NewConfig()
	fmt.Println(conf)
}
