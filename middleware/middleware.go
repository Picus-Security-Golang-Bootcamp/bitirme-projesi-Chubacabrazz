package middlewares

import (
	"net/http"

	"github.com/Chubacabrazz/picus-storeApp/storage/helper"
	"github.com/gin-gonic/gin"
)

//Middleware for checking user token is valid.
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
