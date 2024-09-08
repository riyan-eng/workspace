package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/requestid"
)

func RequestId() app.HandlerFunc {
	return requestid.New(requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
		newRequestID := uuid.NewString()
		return newRequestID
	}))
}
