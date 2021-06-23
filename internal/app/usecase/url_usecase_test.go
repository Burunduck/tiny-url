package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"tiny-url/internal/app/mocks"
)

func TestUrlUsecase_Create(t *testing.T) {

	urlRepo := new(mocks.MockUrlRepo)

	timeoutContext := 2 * time.Second
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	baseUrl := "ozon.com"
	tinyUrl := "aabbccddee"

	urlRepo.On("Create", mock.Anything, mock.AnythingOfType("models.Url")).Return(tinyUrl, nil)
	urlRepo.On("Get", mock.Anything, mock.AnythingOfType("string")).Return("somehash", errors.New("no record"))

	urlUsecase := NewUrlUsecase(urlRepo, timeoutContext)

	tinyUrlMocked, err := urlUsecase.Create(ctx, baseUrl)

	urlRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, tinyUrl, tinyUrlMocked)
}


func TestUrlUsecase_Get(t *testing.T) {
	urlRepo := new(mocks.MockUrlRepo)

	timeoutContext := 2 * time.Second
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	tinyUrl := "aabbccddee"
	baseUrl := "ozon.com"


	urlRepo.On("Get", mock.Anything, mock.AnythingOfType("string")).Return(baseUrl, nil)

	urlUsecase := NewUrlUsecase(urlRepo, timeoutContext)

	baseUrlMock, err := urlUsecase.Get(ctx, tinyUrl)

	urlRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, baseUrl, baseUrlMock)
}
