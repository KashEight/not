package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/KashEight/not/utils"
	"net/http"
	"reflect"
	"time"
)

type discordProvider struct {
	webhook
	Content   string          `json:"content,omitempty"`
	Username  string          `json:"username,omitempty"`
	AvatarUrl string          `json:"avatar_url,omitempty"`
	TTS       bool            `json:"tts,omitempty"`
	Embeds    []*discordEmbed `json:"embeds,omitempty"`
}

type discordEmbed struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Color       int       `json:"color,omitempty"`
	// TODO: Add Embedfooters'...
}

func NewDiscordProvider(url string) *discordProvider {
	w := webhook{url: url}
	d := &discordProvider{webhook: w}

	return d
}

func (d *discordProvider) SetText(text string) *discordProvider {
	d.Content = text
	return d
}

func (d *discordProvider) SetAdditionalKey(key string, value interface{}) *discordProvider {
	// TODO: check invalid `struct`/`slice` (in struct) values
	isInvalid := true
	dv := reflect.ValueOf(&d)

	for i := 0; i < dv.NumField(); i++ {
		field := dv.Type().Field(i)
		jsonTag := field.Tag.Get("json")

		if key != jsonTag {
			continue
		}

		if v := reflect.ValueOf(&value); v.Type() == field.Type {
			dv.Elem().Set(v)
			isInvalid = false
		}
	}

	if isInvalid {
		errS := fmt.Sprintf("invalid `key` or `value`. key: %s, value (type): %s", key, dv.Type())
		panic(errS)
	}

	return d
}

func (d *discordProvider) Do(ctx *context.Context) (*http.Response, error) {
	body, err := json.Marshal(d)

	if err != nil {
		return nil, err
	}

	req := utils.NewRequest("POST")
	res, err := req.SetURL(d.url).SetBody(body).SetHeader("Content-Type", "application/json").Do(ctx)

	return res, err
}
