package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBlog(r *gin.Engine) {
	r.GET("/blog/posts/newest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "newestPosts.html", nil)
	})
}