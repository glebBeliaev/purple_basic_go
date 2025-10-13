package api

import (
	"purple_basic_go/password/config"
)

type API struct {
	key string
}

func NewAPI() *API {
	cfg := config.NewConfig()
	return &API{key: cfg.Key}
}
