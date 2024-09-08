package util

import (
	"math"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gofiber/fiber/v2"
)

func CurrentUser(c *fiber.Ctx) *AccessTokenClaims {
	a:= c.Locals("claim")
	return a.(*AccessTokenClaims)
}

func GetRequestError(c *app.RequestContext) any {
	a, _ := c.Get("error")
	return a
}

func RoundFloat(val float64, precision uint) float64 {
	if math.IsNaN(val) {
		return 0
	}
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Average(total, n float64) float64 {
	if total == 0 || n == 0 {
		return 0
	}
	avg := total / n
	return RoundFloat(avg, 2)
}

func Percentage(val, total float64) float64 {
	if val == 0 || total == 0 {
		return 0
	}

	per := val / total * 100
	return RoundFloat(per, 2)
}
