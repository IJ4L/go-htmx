package apps

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func serveImage(c *gin.Context) {
	imageName := c.Param("imageName")

	imagePath := filepath.Join("/home/ijulyy/go-htmx/external/public/assets/images/", imageName)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "gambar tidak ditemukan"})
		return
	}

	c.File(imagePath)
}

func serveIcon(c *gin.Context) {
	iconName := c.Param("iconName")

	iconPath := filepath.Join("/home/ijulyy/go-htmx/external/public/assets/icons/", iconName)

	if _, err := os.Stat(iconPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "gambar tidak ditemukan"})
		return
	}

	c.File(iconPath)
}

func serveStorage(c *gin.Context) {
	fileName := c.Param("fileName")

	filePath := filepath.Join("/home/ijulyy/go-htmx/external/public/assets/storages/", fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "gambar tidak ditemukan"})
		return
	}

	c.File(filePath)
}
