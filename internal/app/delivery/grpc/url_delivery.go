package grpc

import (
	"tiny-url/internal/models"
	desc "tiny-url/pkg/tiny-url-api"
)

type Server struct {
	desc.UserServiceServer
	UserUsecase models.UserUsecase
}
