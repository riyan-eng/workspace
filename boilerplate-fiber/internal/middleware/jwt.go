package middleware

import (
	"context"
	"server/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Jwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		authHeader := c.Get("Authorization")
		bearerString := string(authHeader)
		if bearerString == "" {
			return util.NewResponse(c).Error("authorization header is required", "", 400)

		}
		token, found := strings.CutPrefix(bearerString, "Bearer ")
		if !found {
			return util.NewResponse(c).Error("undefined token", "", 400)

		}
		claims, errT := util.NewToken().ParseAccess(&token)
		if errT != nil {
			return util.NewResponse(c).Error(errT.Error(), "", 401)

		}
		if errT := util.NewToken().ValidateAccess(&ctx, claims); errT != nil {
			return util.NewResponse(c).Error(errT.Error(), "", 401)

		}

		c.Locals("claim", claims)
		return c.Next()
	}
}
