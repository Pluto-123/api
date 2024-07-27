package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Down struct {
}

func (*Down) DownloadFile(c *gin.Context) {
	filePath := "/app/deploy/test.txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"message": "File not found"})
		return
	}
	c.File(filePath)
}
