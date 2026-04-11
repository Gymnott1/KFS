package apiinternal

import (
	"net/http"

	"clientcompany/handlers"
	"clientcompany/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterInternalRoutes mounts all internal API routes under /api/v1
func RegisterInternalRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", handlers.Logout)
	}

	protected := rg.Group("/")
	protected.Use(middleware.RequireAuth())
	{
		protected.GET("/me", handlers.Me)
		// Add more protected routes here
	}

	admin := rg.Group("/admin")
	admin.Use(middleware.RequireAuth(), middleware.RequireRole("admin"))
	{
		admin.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "admin users list"})
		})
	}
}
