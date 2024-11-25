package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/external/public/landing"
)

func GetHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", landing.Index)
}
