package api

import (
	"app/bin/config"
)

type Api struct {
	config *config.Config
}

func NewApi(conf *config.Config) *Api {
	return &Api{config: conf}
}
