package services

import (
	"meli"
	"net/http"
)

type BaseService interface {
	GetMeli() meli.Meli
	GetCli() *http.Client
}

type baseService struct {
	m   meli.Meli
	cli *http.Client
}
