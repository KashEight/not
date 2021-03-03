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
	Webhook
	Content   string          `json:"content,omitempty"`
	Username  string          `json:"username,omitempty"`
	AvatarUrl string          `json:"avatar_url,omitempty"`
	TTS       bool            `json:"tts,omitempty"`
	Embeds    []*DiscordEmbed `json:"embeds,omitempty"`
}

type DiscordEmbed struct {
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   time.Time       `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

func NewDiscordProvider(url string) *discordProvider {
	w := Webhook{url: url}
	d := &discordProvider{Webhook: w}

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
