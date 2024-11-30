package shop

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/external/public/landing"
	"github.com/ij4l/go-htmx/model"
)

type ShopHandler struct {
	hs ShopService
}

func NewShopHandler(hs ShopService) ShopHandler {
	return ShopHandler{hs: hs}
}

func (sh ShopHandler) GetShop(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	ctgStr := ctx.DefaultQuery("ctg", "Bucket Bunga")

	products, err := sh.hs.getProductsByCategory(ctgStr, 4, page-1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	total, err := sh.hs.GetTotalProductsByCategory(ctgStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(3)))
	prevPage := page > 1
	nextPage := page < totalPages

	rsp := model.ProductANTD{
		Page:         page,
		NextPage:     nextPage,
		PrevPage:     prevPage,
		TotalShow:    len(products),
		TotalProduct: total,
		Products:     products,
	}

	ctx.HTML(http.StatusOK, "", landing.Shop(rsp))
}
