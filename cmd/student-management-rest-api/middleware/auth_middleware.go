package middleware

import (
	"Student_Management_Rest_API_GO/cmd/student-management-rest-api/utility"
	"github.com/gin-gonic/gin"
	"strings"
)

var skipAuthRoutes = []string{"/login", "/login/", "/dashboard", "/dashboard/"}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//skip
		for _, route := range skipAuthRoutes {
			if route == c.FullPath() {
				c.Next()
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utility.JwtTokenErrorResponse(c, "Missing Token.", nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, claims, err := utility.ValidateToken(tokenString)
		if err != nil {
			utility.JwtTokenErrorResponse(c, "Invalid Token.", nil)
		}

		c.Set("userName", claims["userName"])
		c.Set("userId", claims["userId"])
		c.Set("userRole", claims["userRole"])

		c.Next()
	}
}
