package middleware

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var skipAuthRoutes = map[string]bool{
	"/auth/login":     true,
	"/auth/login/":    true,
	"/auth/register":  true,
	"/auth/register/": true,
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//skip
		fmt.Println(c.FullPath())
		if skipAuthRoutes[c.FullPath()] { //faster lookup for using map
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utility.JwtTokenErrorResponse(c, "Missing Token.", nil)
			c.Abort() // Stop further processing
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, claims, err := utility.ValidateToken(tokenString)
		if err != nil {
			utility.JwtTokenErrorResponse(c, "Invalid Token.", nil)
			c.Abort() // Stop further processing
			return
		}

		c.Set("userName", claims["userName"])
		c.Set("userId", claims["userId"])
		c.Set("userRole", claims["userRole"])

		c.Next()
	}
}
