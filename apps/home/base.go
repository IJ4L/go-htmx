package home

import (
	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/apps"
)

func InitializeHomeHandler(router *gin.Engine, repo *apps.AppRepository) {
	ctx := &gin.Context{}
	er := NewHomeRepository(repo)
	as := NewHomeService(ctx, er)
	hh := NewHomeHandler(as)

	router.GET("/", hh.GetIndex)
	router.GET("/home", hh.GetHome)
}
