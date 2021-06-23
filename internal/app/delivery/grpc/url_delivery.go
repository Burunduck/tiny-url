package grpc

import (
	"context"
	"tiny-url/internal/models"
	desc "tiny-url/pkg/tiny-url-api"
)

type Server struct {
	desc.UrlServiceServer
	UrlUsecase models.UrlUsecase
}

func NewServer(urlUsecase models.UrlUsecase) *Server {
	return &Server{UrlUsecase: urlUsecase}
}

func (s *Server) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	u := models.Url{
		BaseUrl: request.BaseUrl,
	}
	var err error
	u.TinyUrl, err = s.UrlUsecase.Create(ctx, u.BaseUrl)
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{
		TinyUrl: u.TinyUrl,
	}, err
}

func (s *Server) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	u := models.Url{
		TinyUrl: request.TinyUrl,
	}
	var err error
	u.BaseUrl, err = s.UrlUsecase.Get(ctx, u.TinyUrl)
	if err != nil {
		return nil, err
	}
	return &desc.GetResponse{
		BaseUrl: u.BaseUrl,
	}, err
}


