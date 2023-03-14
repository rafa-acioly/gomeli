package services

import (
	"fmt"
	"meli/storage"
	"strconv"
	"time"
)

type AccessTokenService interface {
	Get() string
	Save(token string)
	GetRefreshToken() string
	SaveRefreshToken(token string)
	GetExpiration() int
	SaveExpiration(expiration string)
	IsValid() bool
	IsExpired() bool
}

type accessTokenService struct {
	key string
	stg storage.Storage
}

func (t accessTokenService) Get() string {
	result, _ := t.stg.Get("token_" + t.key)
	return result
}

func (t accessTokenService) Save(tk string) {
	_ = t.stg.Set("token_"+t.key, tk)
}

func (t accessTokenService) SaveRefreshToken(tk string) {
	_ = t.stg.Set("refresh_token_"+t.key, tk)
}

func (t accessTokenService) GetRefreshToken() string {
	result, _ := t.stg.Get("refresh_token_" + t.key)
	return result
}

func (t accessTokenService) SaveExpiration(expiration string) {
	_ = t.stg.Set("expire_in_"+t.key, expiration)
}

func (t accessTokenService) GetExpiration() int {
	result, _ := t.stg.Get("expire_in_" + t.key)
	content, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}

	return content
}

func (t accessTokenService) IsExpired() bool {
	tokenExpiration := time.UnixMilli(int64(t.GetExpiration()))

	return !tokenExpiration.After(time.Now())
}

func (t accessTokenService) IsValid() bool {
	fmt.Println(t.Get(), t.IsExpired())
	if len(t.Get()) == 0 || t.IsExpired() {
		return false
	}

	return true
}

func NewAccessTokenService(key string, stg storage.Storage) AccessTokenService {
	return &accessTokenService{key: key, stg: stg}
}
