package middleware

import (
	"context"
	"server/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		authHeader := c.GetHeader("Authorization")
		bearerString := string(authHeader)
		if bearerString == "" {
			util.NewResponse(c).Error("authorization header is required", "", 400)
			return
		}
		token, found := strings.CutPrefix(bearerString, "Bearer ")
		if !found {
			util.NewResponse(c).Error("undefined token", "", 400)
			return

		}
		claims, errT := util.NewToken().ParseAccess(&token)
		if errT != nil {
			util.NewResponse(c).Error(errT.Error(), "", 401)
			return
		}
		if errT := util.NewToken().ValidateAccess(&ctx, claims); errT != nil {
			util.NewResponse(c).Error(errT.Error(), "", 401)
			return
		}

		c.Set("claim", claims)
		c.Next()
	}
}
