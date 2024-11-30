package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/apps"
)

func InitializeShopHandler(router *gin.Engine, repo *apps.AppRepository) {
	ctx := &gin.Context{}
	er := NewShopRepository(repo)
	as := NewShopService(ctx, er)
	hh := NewShopHandler(as)

	router.GET("/shop", hh.GetShop)
}
