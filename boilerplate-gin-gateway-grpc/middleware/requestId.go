package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return requestid.New(
		requestid.WithGenerator(func() string {
			newRequestID := uuid.NewString()
			return newRequestID
		}),
	)
}
