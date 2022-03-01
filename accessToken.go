package meli

import (
	"meli/storage"
)

type AccessTokenService interface {
	Get() string
	Save(token string)
	GetRefreshToken() string
	SaveRefreshToken(token string)
}

type token struct {
	key string
	stg storage.Storage
}

func (t token) Get() string {
	result, _ := t.stg.Get("token_" + t.key)
	return result
}

func (t token) Save(tk string) {
	_ = t.stg.Set("token_"+t.key, tk)
}

func (t token) GetRefreshToken() string {
	result, _ := t.stg.Get("refresh_token_" + t.key)
	return result
}

func (t token) SaveRefreshToken(tk string) {
	_ = t.stg.Set("refresh_token_"+t.key, tk)
}

func NewAccessToken(key string, stg storage.Storage) AccessTokenService {
	return &token{key: key, stg: stg}
}
