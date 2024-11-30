package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/apps"
)

func InitializeAdminHandler(router *gin.Engine, repo *apps.AppRepository) {
	ctx := &gin.Context{}
	er := NewAdminRepository(repo)
	as := NewAdminService(ctx, er)
	hh := NewAdminHandler(as)

	router.GET("/admin", hh.GetIndex)
	router.POST("/admin/add-product", hh.PostAddProduct)
	router.DELETE("/admin/delete-product/:id", hh.PostDeleteProduct)
}
