package middleware

import (
	"net/http"
	"thunder_hoster/handler"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessControlMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if handler.ValidTime.Before(time.Now()) {
			// 已经超时
			ctx.AbortWithStatus(http.StatusServiceUnavailable)
		}
		ctx.Next()
	}
}
