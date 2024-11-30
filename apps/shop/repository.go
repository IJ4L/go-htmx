package shop

import (
	"context"

	"github.com/ij4l/go-htmx/apps"
	db "github.com/ij4l/go-htmx/external/database/sqlc"
	typ "github.com/ij4l/go-htmx/external/database/type"
	"github.com/ij4l/go-htmx/model"
)

type ShopRepository struct {
	apps apps.AppRepository
}

// getTotalProducts implements ShopRepositoryContract.
func (h ShopRepository) getTotalProducts(ctx context.Context, ctg string) (int, error) {
	total, err := h.apps.GetTotalProductsByCategory(ctx, typ.Text(ctg))
	if err != nil {
		return 0, err
	}

	return int(total), nil
}

func (h ShopRepository) getProductsByCategory(ctx context.Context, ctg string, limit, offset int64) ([]model.Product, error) {
	arg := db.GetProductsByCategoryParams{
		Category: typ.Text(ctg),
		Limit:    limit,
		Offset:   offset,
	}

	products, err := h.apps.GetProductsByCategory(ctx, arg)
	if err != nil {
		return nil, err
	}

	return model.NewList(products), nil
}

func (h ShopRepository) getProducts(ctx context.Context) ([]model.Product, error) {
	arg := db.GetAllProductsWithPaginationParams{
		Limit:  10,
		Offset: 0,
	}

	products, err := h.apps.GetAllProductsWithPagination(ctx, arg)
	if err != nil {
		return nil, err
	}

	return model.NewList(products), nil
}

func NewShopRepository(apps *apps.AppRepository) ShopRepository {
	return ShopRepository{apps: *apps}
}
