package psql

import (
	"context"
	"gorm.io/gorm"
	"tiny-url/internal/models"
)

type UrlRepository struct {
	models.UrlRepository
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UrlRepository {
	return &UrlRepository{
		DB: db,
	}
}

func (u UrlRepository) Create(ctx context.Context, url models.Url) (string, error) {
	panic("implement me")
}

func (u UrlRepository) Get(ctx context.Context, tinyUrl string) (string, error) {
	panic("implement me")
}
