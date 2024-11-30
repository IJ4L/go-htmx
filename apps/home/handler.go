package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/external/public/landing"
)

type HomeHandler struct {
	hs HomeService
}

func NewHomeHandler(hs HomeService) HomeHandler {
	return HomeHandler{hs: hs}
}

func (hh HomeHandler) GetIndex(ctx *gin.Context) {
	products, err := hh.hs.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "", landing.Index(products))
}

func (hh HomeHandler) GetHome(ctx *gin.Context) {
	products, err := hh.hs.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "", landing.Home(products))
}
