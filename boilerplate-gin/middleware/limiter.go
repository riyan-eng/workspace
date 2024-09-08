package middleware

import (
	"context"
	"log"
	"server/config"

	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		_, _, _, ok, err := config.LimiterStore.Take(ctx, "key")
		if err != nil {
			log.Fatal(err)
		}
		if !ok {
			c.AbortWithStatusJSON(429, "too many request")
		}
		c.Next()
	}
}
