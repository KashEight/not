package webhook

import (
	"context"
	"net/http"
)

type webhook struct {
	url string
}

type ProviderWebhook interface {
	SetText(text string) *ProviderWebhook
	SetAdditionalKey(key string, value interface{}) *ProviderWebhook
	Do(ctx *context.Context) (*http.Response, error)
}

// TODO: make `Provider` more customizable
