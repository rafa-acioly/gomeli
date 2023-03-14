package meli

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClientWrapper interface {
	Get(resource string) (*resty.Response, error)
	Post(resource string, data interface{}) (*resty.Response, error)
	Put(resource string, data interface{}) (*resty.Response, error)
	Delete(resource string, data interface{}) (*resty.Response, error)
	SetResult(result interface{}) HttpClientWrapper
	SetError(err interface{}) HttpClientWrapper
	SetHeader(key, value string) HttpClientWrapper
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

// SetError implements HttpClientWrapper
func (h *httpClientWrapper) SetError(err interface{}) HttpClientWrapper {
	h.cli.SetError(err)

	return h
}

// SetHeader implements HttpClientWrapper
func (h *httpClientWrapper) SetHeader(key string, value string) HttpClientWrapper {
	h.cli.SetHeader(key, value)

	return h
}

// SetResult implements HttpClientWrapper
func (h *httpClientWrapper) SetResult(result interface{}) HttpClientWrapper {
	h.cli.SetResult(result)

	return h
}

func NewHttpClient(m Meli) HttpClientWrapper {
	return &httpClientWrapper{cli: resty.New().
		SetTimeout(time.Second * 15).
		SetDebug(true).
		SetBaseURL(m.GetEnvironment().GetWsHost()).
		SetRetryCount(3).
		SetRetryMaxWaitTime(time.Second * 2).
		SetHeaders(map[string]string{
			"User-Agent":   "MELI-GOLANG-SDK",
			"Content-Type": "application/json; charset=utf-8",
			"verify":       "true",
		}).R(),
	}
}
