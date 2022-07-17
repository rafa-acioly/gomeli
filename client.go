package meli

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type HttpClientWrapper interface {
	Get(resource string) (*resty.Response, error)
	Post(resource string, data interface{}) (*resty.Response, error)
	Put(resource string, data interface{}) (*resty.Response, error)
	Delete(resource string, data interface{}) (*resty.Response, error)
}

type httpClientWrapper struct {
	cli *resty.Request
}

func (h httpClientWrapper) Post(resource string, data interface{}) (*resty.Response, error) {
	response, err := h.cli.SetBody(data).Post(resource)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h httpClientWrapper) Put(resource string, data interface{}) (*resty.Response, error) {
	response, err := h.cli.SetBody(data).Put(resource)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h httpClientWrapper) Delete(resource string, data interface{}) (*resty.Response, error) {
	response, err := h.cli.SetBody(data).Delete(resource)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h httpClientWrapper) Get(resource string) (*resty.Response, error) {
	response, err := h.cli.Get(resource)
	if err != nil {
		return response, err
	}

	return response, nil
}

func NewHttpClient(m Meli) HttpClientWrapper {
	return httpClientWrapper{cli: resty.New().
		SetTimeout(time.Minute * 1).
		SetDebug(true).
		SetBaseURL(m.GetEnvironment().GetWsHost()).
		SetHeaders(map[string]string{
			"User-Agent":   "MELI-GOLANG-SDK",
			"Content-Type": "application/json; charset=utf-8",
			"verify":       "true",
		}).R(),
	}
}
