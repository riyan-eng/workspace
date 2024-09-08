package middleware

import (
	"server/config"
	"server/util"

	"github.com/gofiber/fiber/v2"
)

func Permission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := util.CurrentUser(c)
		if user.UserId == "" {
			return util.NewResponse(c).Error("current logged in user not found", "", 401)
		}

		enforcer := config.NewEnforcer()
		if err := enforcer.LoadPolicy(); err != nil {
			return util.NewResponse(c).Error("failed to load policy", "", 500)
		}

		uri := c.OriginalURL()
		method := c.Method()
		accepted, err := enforcer.Enforce(user.UserId, uri, method) // userID - url - method
		if err != nil {
			return util.NewResponse(c).Error(err.Error(), "error when authorizing user's accessibility", 500)
		}

		if !accepted {
			return util.NewResponse(c).Error("you are not allowed", "kamu tidak diizinkan", 403)
		}

		return c.Next()
	}
}
