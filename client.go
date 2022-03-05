package meli

import (
	"github.com/idoubi/goz"
)

func buildHttpClient(m Meli) *goz.Request {
	return goz.NewClient(goz.Options{
		Debug:   true,
		BaseURI: m.GetEnvironment().GetWsHost(),
		Timeout: 15,
		Headers: map[string]interface{}{
			"User-Agent":   "MELI-GOLANG-SDK",
			"Content-Type": "application/json; charset=utf-8",
			"verify":       true,
		},
		Query: map[string]interface{}{},
	})
}

type HttpClientWrapper interface {
	Get(resource string) (*goz.Response, error)
	Post(resource string, data interface{}) (*goz.Response, error)
	Put(resource string, data interface{}) (*goz.Response, error)
	Delete(resource string, data interface{}) (*goz.Response, error)
}

type httpClientWrapper struct {
	cli *goz.Request
}

func (h httpClientWrapper) Post(resource string, data interface{}) (*goz.Response, error) {
	response, err := h.cli.Post(resource, goz.Options{JSON: data})
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h httpClientWrapper) Put(resource string, data interface{}) (*goz.Response, error) {
	response, err := h.cli.Put(resource, goz.Options{JSON: data})
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h httpClientWrapper) Delete(resource string, data interface{}) (*goz.Response, error) {
	response, err := h.cli.Delete(resource, goz.Options{JSON: data})
	if err != nil {
		return response, err
	}

	return response, nil
}

func (h httpClientWrapper) Get(resource string) (*goz.Response, error) {
	response, err := h.cli.Get(resource)
	if err != nil {
		return response, err
	}

	return response, nil
}

func NewHttpClient(m Meli) HttpClientWrapper {
	return httpClientWrapper{
		cli: goz.NewClient(goz.Options{
			Debug:   true,
			BaseURI: m.GetEnvironment().GetWsHost(),
			Timeout: 15,
			Headers: map[string]interface{}{
				"User-Agent":   "MELI-GOLANG-SDK",
				"Content-Type": "application/json; charset=utf-8",
				"verify":       true,
			},
			Query: map[string]interface{}{},
		}),
	}
}
