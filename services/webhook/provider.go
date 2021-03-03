package webhook

import (
	"context"
	"net/http"
)

type Webhook struct {
	url string
}

type Provider interface {
	SetText(text string) *Provider
	SetAdditionalKey(key string, value interface{}) *Provider
	Do(ctx *context.Context) (*http.Response, error)
}

// TODO: make `Provider` more customizable
