package models

import "context"

type Url struct {
	TinyUrl string `gorm:"size:10"`
	BaseUrl string
}

type UrlUsecase interface {
	Create(ctx context.Context, baseUrl string) (string, error)
	Get(ctx context.Context, tinyUrl string) (string, error)
}

type UrlRepository interface {
	Create(ctx context.Context, baseUrl string) (string, error)
	Get(ctx context.Context, tinyUrl string) (string, error)
}