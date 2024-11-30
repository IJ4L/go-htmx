package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/external/public/admin"
)

type AdminHandler struct {
	hs AdminService
}

func NewAdminHandler(hs AdminService) AdminHandler {
	return AdminHandler{hs: hs}
}

func (hh AdminHandler) GetIndex(ctx *gin.Context) {
	products, err := hh.hs.getProducts()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "", admin.Admin(products))
}

func (hh AdminHandler) PostAddProduct(ctx *gin.Context) {
	var inp AddProduct
	if err := ctx.ShouldBind(&inp); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := hh.hs.addProduct(inp, ctx); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	products, err := hh.hs.getProducts()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "", admin.Admin(products))
}

func (hh AdminHandler) PostDeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := hh.hs.deleteProduct(productID); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
