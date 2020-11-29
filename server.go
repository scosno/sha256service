package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	r.GET("/sha256/:item", func(c *gin.Context) {
		item := c.Param("item")
		hash := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(item))))
		//hash := sha256.Sum256([]byte(item))
		c.String(http.StatusOK, "%s", hash)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
