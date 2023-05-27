package logger

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger = *zap.Logger

func New() (Logger, error) {
	zapCfg := zap.NewDevelopmentConfig()
	l, err := zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("create logger failed: %w", err)
	}
	zap.ReplaceGlobals(l)
	return l, nil
}

func SetCookieMiddleware(log Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info("Cookies check", zap.String("cookies", getCookiesString(ctx)))
		ctx.Next()
	}
}

func getCookiesString(ctx *gin.Context) string {
	var cookies []string
	for _, cookie := range ctx.Request.Cookies() {
		cookies = append(cookies, fmt.Sprintf("(%s=%s)", cookie.Name, cookie.Value))
	}
	return strings.Join(cookies, " ")
}
