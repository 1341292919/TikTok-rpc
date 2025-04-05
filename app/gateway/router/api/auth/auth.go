package auth

import (
	"TikTok-rpc/app/gateway/middleware/jwt"
	"TikTok-rpc/app/gateway/pack"
	"TikTok-rpc/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Auth() []app.HandlerFunc {
	//为了有扩展性
	return append(make([]app.HandlerFunc, 0),
		DoubleTokenAuthFunc(),
	)
}

func DoubleTokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.IsAccessTokenAvailable(ctx, c) {
			if !jwt.IsRefreshTokenAvailable(ctx, c) {
				pack.SendFailResponse(c, errno.AuthInvalid)
				c.Abort()
				return
			}
			jwt.GenerateAccessToken(c)
		}
		c.Next(ctx)
	}
}
