package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(jwtSecret string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication logic here
		c.Next()
	}
}
