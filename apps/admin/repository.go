package admin

import (
	"context"

	"github.com/ij4l/go-htmx/apps"
	db "github.com/ij4l/go-htmx/external/database/sqlc"
	typ "github.com/ij4l/go-htmx/external/database/type"
	"github.com/ij4l/go-htmx/model"
)

type AdminRepository struct {
	apps apps.AppRepository
}

func (h AdminRepository) getProductByID(ctx context.Context, inp int) (model.Product, error) {
	product, err := h.apps.GetProductById(ctx, int32(inp))
	if err != nil {
		return model.Product{}, err
	}

	return model.New(product), nil
}

func (h AdminRepository) addProduct(ctx context.Context, inp AddProduct, imgUri string) error {
	arg := db.CreateProductParams{
		Name:     inp.Name,
		Category: typ.Text(inp.Category),
		Price:    typ.Numeric(inp.Price),
		ImageUrl: typ.Text(imgUri),
	}

	_, err := h.apps.CreateProduct(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (h AdminRepository) deleteProduct(ctx context.Context, inp int) error {
	err := h.apps.DeleteProduct(ctx, int32(inp))
	if err != nil {
		return err
	}

	return nil
}

func (h AdminRepository) getProducts(ctx context.Context) ([]model.Product, error) {
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

func NewAdminRepository(apps *apps.AppRepository) AdminRepository {
	return AdminRepository{apps: *apps}
}
