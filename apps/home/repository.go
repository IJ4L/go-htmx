package home

import (
	"context"

	"github.com/ij4l/go-htmx/apps"
	db "github.com/ij4l/go-htmx/external/database/sqlc"
	"github.com/ij4l/go-htmx/model"
)

type HomeRepository struct {
	apps apps.AppRepository
}

func (h HomeRepository) getProducts(ctx context.Context) ([]model.Product, error) {
	arg := db.GetAllProductsWithPaginationParams{
		Limit:  8,
		Offset: 0,
	}

	products, err := h.apps.GetAllProductsWithPagination(ctx, arg)
	if err != nil {
		return nil, err
	}

	return model.NewList(products), nil
}

func NewHomeRepository(apps *apps.AppRepository) HomeRepository {
	return HomeRepository{apps: *apps}
}
