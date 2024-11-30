package home

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/model"
)

type HomeRepositoryContract interface {
	getProducts(ctx context.Context) ([]model.Product, error)
}

type HomeService struct {
	ctx  *gin.Context
	repo HomeRepositoryContract
}

func NewHomeService(ctx *gin.Context, repo HomeRepositoryContract) HomeService {
	return HomeService{ctx: ctx, repo: repo}
}

func (s HomeService) GetProducts() ([]model.Product, error) {
	return s.repo.getProducts(s.ctx)
}
