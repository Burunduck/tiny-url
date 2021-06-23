package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"tiny-url/internal/models"
)

type MockUrlRepo struct {
	mock.Mock
}

func (mock *MockUrlRepo) Create(ctx context.Context, url models.Url) (string, error) {
	args := mock.Called(ctx, url)
	return args.Get(0).(string), args.Error(1)
}

func (mock *MockUrlRepo) Get(ctx context.Context, tinyUrl string) (string, error){
	args := mock.Called(ctx, tinyUrl)
	return args.Get(0).(string), args.Error(1)
}
