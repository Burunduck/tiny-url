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

func (ur UrlRepository) Create(ctx context.Context, url models.Url) (string, error) {
	err:= ur.DB.WithContext(ctx).Create(url).Error
	if err != nil{
		return "", err
	}
	return url.TinyUrl, nil
}

func (ur UrlRepository) Get(ctx context.Context, tinyUrl string) (string, error) {
	url := models.Url{}
	err := ur.DB.WithContext(ctx).First(&url, "tiny_url = ?", tinyUrl).Error
	if err != nil {
		return "", err
	}
	return url.BaseUrl, nil
}
