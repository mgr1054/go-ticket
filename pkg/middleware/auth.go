package middlewares

import (
	"github.com/mgr1054/go-ticket/pkg/utils"
	"github.com/gin-gonic/gin"
)


// validate token from gin http request 
func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		
		c.Set("role", claims.Role)
		c.Next()
	}
}