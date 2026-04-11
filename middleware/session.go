package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// TrackSession logs request metadata for session tracking.
// Extend this to persist to DB using models.Session if needed.
func TrackSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[%s] %s %s | %d | %s | %s",
			time.Now().Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
			c.ClientIP(),
		)
	}
}
