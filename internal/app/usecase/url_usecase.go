package usecase

import (
	"context"
	"time"
	"tiny-url/internal/models"
)

type UrlUsecase struct {
	urlRepo        models.UrlRepository
	contextTimeout time.Duration
}

func NewUrlUsecase(ur models.UrlRepository, t time.Duration) models.UrlUsecase {
	return &UrlUsecase{
		urlRepo:        ur,
		contextTimeout: t}
}

func (u UrlUsecase) Create(ctx context.Context, baseUrl string) (string, error) {
	panic("implement me")
}

func (u UrlUsecase) Get(ctx context.Context, tinyUrl string) (string, error) {
	panic("implement me")
}
