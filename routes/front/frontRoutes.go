package front

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleFrontend(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	})
}