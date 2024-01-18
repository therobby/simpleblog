package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/therobby/simpleblog/routes/blog"
	"github.com/therobby/simpleblog/routes/front"
	"github.com/therobby/simpleblog/routes/user"
//	"path/filepath"
)

func HandleRoutes(r *gin.Engine) {
	r.Static("/assets", "frontend/assets")
	//	r.LoadHTMLGlob("frontend/templates/**")

	// Load templates with parts
	//	templatesDir := "frontend/templates"
	//	partsDir := filepath.Join(templatesDir, "parts")
	//	r.LoadHTMLGlob(filepath.Join(templatesDir, "*.html"))
	//	r.LoadHTMLGlob(filepath.Join(partsDir, "*.html"))

	r.LoadHTMLGlob("frontend/templates/**/*.html")
//	r.LoadHTMLGlob("frontend/templates/*.html")

	blog.HandleBlog(r)
	user.HandleUser(r)
	front.HandleFrontend(r)
}
