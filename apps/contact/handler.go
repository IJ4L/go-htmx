package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ij4l/go-htmx/external/public/landing"
)

func GetContact(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", landing.Contact())
}
