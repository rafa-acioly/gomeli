package meli

import "net/http"

type BaseService interface {
	GetMeli() Meli
	GetCli() *http.Client
}

type baseService struct {
	m   Meli
	cli *http.Client
}
