package shop

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/model"
)

type ShopRepositoryContract interface {
	getProducts(ctx context.Context) ([]model.Product, error)
	getProductsByCategory(ctx context.Context, ctg string, limit, offset int64) ([]model.Product, error)
	getTotalProducts(ctx context.Context, ctg string) (int, error)
}

type ShopService struct {
	ctx  *gin.Context
	repo ShopRepositoryContract
}

func NewShopService(ctx *gin.Context, repo ShopRepositoryContract) ShopService {
	return ShopService{ctx: ctx, repo: repo}
}

func (s ShopService) getProductsByCategory(ctg string, page, pageSize int) ([]model.Product, error) {
	offset := (page - 1) * pageSize
	return s.repo.getProductsByCategory(s.ctx, ctg, int64(page), int64(offset))
}

func (s ShopService) GetProducts() ([]model.Product, error) {
	return s.repo.getProducts(s.ctx)
}

func (s ShopService) GetTotalProductsByCategory(ctg string) (int, error) {
	return s.repo.getTotalProducts(s.ctx, ctg)
}
