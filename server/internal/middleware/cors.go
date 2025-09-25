package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		log.Printf("CORS Request: %s %s from Origin: %s", c.Request.Method, c.Request.URL.Path, origin)

		// Allow requests from localhost:4200 (Angular dev server) or no origin (direct access)
		if origin == "http://localhost:4200" || origin == "" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		}

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Max-Age", "86400") // 24 hours

		if c.Request.Method == "OPTIONS" {
			log.Printf("Preflight OPTIONS request handled for %s", c.Request.URL.Path)
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
