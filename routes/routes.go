package routes

import (
	"net/http"
	"os"
	"path/filepath"

	apiinternal "clientcompany/api/apiinternal"
	"clientcompany/middleware"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.Use(middleware.TrackSession())

	// API always first
	api := r.Group("/api/v1")
	apiinternal.RegisterInternalRoutes(api)

	// Serve Vue dist if built, otherwise fall back to templates
	if _, err := os.Stat("./dist/index.html"); err == nil {
		serveVue(r)
	} else {
		serveTemplates(r)
	}
}

func serveVue(r *gin.Engine) {
	r.Static("/assets", "./dist/assets")
	r.NoRoute(func(c *gin.Context) {
		path := filepath.Join("dist", c.Request.URL.Path)
		if _, err := os.Stat(path); err == nil {
			c.File(path)
			return
		}
		c.File("./dist/index.html")
	})
}

func serveTemplates(r *gin.Engine) {
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
}
