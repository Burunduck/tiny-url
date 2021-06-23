package usecase

import (
	"context"
	"time"
	"tiny-url/internal/models"
	"tiny-url/pkg/utils"
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
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	newTinyUrl := utils.Hash(baseUrl)
	baseUrl, err := u.Get(ctx, newTinyUrl)

	for baseUrl != "" {
		newTinyUrl := utils.Hash(baseUrl)
		baseUrl, _ = u.Get(ctx, newTinyUrl)
	}

	tinyUrl, err := u.urlRepo.Create(ctx, models.Url{
		TinyUrl: newTinyUrl,
		BaseUrl: baseUrl,
	})

	if err != nil {
		return "", err
	}

	return tinyUrl, nil
}

func (u UrlUsecase) Get(ctx context.Context, tinyUrl string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	baseUrl, err := u.urlRepo.Get(ctx, tinyUrl)
	if err != nil {
		return "", err
	}
	return baseUrl, nil
}
