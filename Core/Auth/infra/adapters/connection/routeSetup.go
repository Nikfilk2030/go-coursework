package connection

import (
	"net/http"
	"os"
	"time"

	"auth/logger"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
)

func initRouter(a *Adapter, r *gin.Engine) {
	r.Use(loggerContext(a))
	r.Use(corsMiddleware())
	r.Use(logger.SetCookieMiddleware(a.log))
	r.Use(ginzap.Ginzap(a.log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(a.log, true))

	setupRoutes(a, r)
}

func loggerContext(a *Adapter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lCtx := zapctx.WithLogger(ctx.Request.Context(), a.log)
		ctx.Request = ctx.Request.WithContext(lCtx)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		debug := os.Getenv("DEBUG")
		origin := ""
		if debug == "True" {
			origin = "http://localhost:3000"
		} else {
			origin = "http://gigachat.site"
		}
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Date, Content-Length")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == http.MethodOptions {
			ctx.Status(http.StatusOK)
			return
		}
		ctx.Next()
	}
}

func setupRoutes(a *Adapter, r *gin.Engine) {
	api := r.Group("/gigachat/auth/api/v1")
	{
		api.GET("auth", a.handleGetAuth)
		api.GET("login", a.handleGetLogin)
		api.POST("register", a.handlePostRegister)
	}
}
