package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUser(r *gin.Engine) {

	r.GET("/user/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong user",
			})
	})
}