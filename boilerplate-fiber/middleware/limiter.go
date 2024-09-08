package middleware

import (
	"context"
	"server/config"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
)

func Limiter() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		_, _, _, ok, err := config.LimiterStore.Take(ctx, "key")
		if err != nil {
			log.Fatal(err)
		}
		if !ok {
			c.AbortWithMsg("too many request", 429)
		}
		c.Next(ctx)
	}
}
