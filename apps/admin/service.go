package admin

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/internal/utils"
	"github.com/ij4l/go-htmx/model"
)

type adminRepositoryContract interface {
	getProducts(ctx context.Context) ([]model.Product, error)
	getProductByID(ctx context.Context, inp int) (model.Product, error)
	addProduct(ctx context.Context, inp AddProduct, imgUri string) error
	deleteProduct(ctx context.Context, inp int) error
}

type AdminService struct {
	ctx  *gin.Context
	repo adminRepositoryContract
}

func NewAdminService(ctx *gin.Context, repo adminRepositoryContract) AdminService {
	return AdminService{ctx: ctx, repo: repo}
}

func (s AdminService) getProducts() ([]model.Product, error) {
	return s.repo.getProducts(s.ctx)
}



func (s AdminService) addProduct(inp AddProduct, ctx *gin.Context) error {
	file, err := ctx.FormFile("productImage")
	if err != nil {
		return err
	}

	randStr, err := utils.GenerateRandomString(10)
	if err != nil {
		return err
	}

	imgUri := "/home/ijulyy/go-htmx/external/public/assets/storages/" + randStr + filepath.Ext(file.Filename)
	if err := ctx.SaveUploadedFile(file, imgUri); err != nil {
		return err
	}

	_, err = os.ReadFile(imgUri)
	if err != nil {
		return err
	}

	imgName := randStr + filepath.Ext(file.Filename)
	return s.repo.addProduct(s.ctx, inp, imgName)
}

func (s AdminService) deleteProduct(inp int) error {
	product, err := s.repo.getProductByID(s.ctx, inp)
	if err != nil {
		return fmt.Errorf("failed to get product: %w", err)
	}

	imagePath := fmt.Sprintf("/home/ijulyy/go-htmx/external/public/assets/storages/%s", product.ImageUrl)
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return fmt.Errorf("image not found: %w", err)
	}

	err = os.Remove(imagePath)
	if err != nil {
		return fmt.Errorf("failed to delete image: %w", err)
	}

	return s.repo.deleteProduct(s.ctx, inp)
}
