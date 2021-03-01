package utils

import (
	"bytes"
	"context"
	"net/http"
)

type request struct {
	method string
	url    string
	header map[string]string
	body   []byte
}

func NewRequest(method string) *request {
	return &request{method: method}
}

func (r *request) SetURL(url string) *request {
	r.url = url
	return r
}

func (r *request) SetHeader(key string, value string) *request {
	if r.header == nil {
		r.header = make(map[string]string)
	}
	r.header[key] = value
	return r
}

func (r *request) SetBody(body []byte) *request {
	r.body = body
	return r
}

func (r *request) Do(ctx *context.Context) (*http.Response, error) {
	req, err := http.NewRequestWithContext(*ctx, r.method, r.url, bytes.NewBuffer(r.body))

	if err != nil {
		return nil, err
	}

	for k, v := range r.header {
		req.Header.Set(k, v)
	}

	client := new(http.Client)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// TODO: add some funcs which validate `*http.Response`
