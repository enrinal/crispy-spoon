package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, ErrorResponse{
				Error: "Unauthorized",
			})
			c.Abort()
			return
		}

		// remove Bearer
		token = token[7:]

		// validate token
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			c.JSON(401, ErrorResponse{
				Error: "Unauthorized",
			})
			c.Abort()
			return
		}

		if !tkn.Valid {
			c.JSON(401, ErrorResponse{
				Error: "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}

func AuthzAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get role from context
		role := c.MustGet("role")

		// check if role is admin
		if role != "admin" {
			c.JSON(401, ErrorResponse{
				Error: "Unauthorized",
			})
			c.Abort()
			return
		}

		// continue
		c.Next()
	}
}
