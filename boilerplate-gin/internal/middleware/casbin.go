package middleware

import (
	"server/config"
	"server/util"

	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := util.CurrentUser(c)
		if user.UserId == "" {
			util.NewResponse(c).Error("current logged in user not found", "", 401)
			return
		}

		enforcer := config.NewEnforcer()
		if err := enforcer.LoadPolicy(); err != nil {
			util.NewResponse(c).Error("failed to load policy", "", 500)
			return
		}

		uri := c.Request.URL.Path
		method := c.Request.Method

		accepted, err := enforcer.Enforce(user.UserId, uri, method) // userID - url - method
		if err != nil {

			util.NewResponse(c).Error(err.Error(), "error when authorizing user's accessibility", 500)
			return
		}

		if !accepted {
			util.NewResponse(c).Error("you are not allowed", "kamu tidak diizinkan", 403)
			return
		}

		c.Next()
	}
}
