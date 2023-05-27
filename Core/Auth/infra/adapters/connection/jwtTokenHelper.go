package connection

import (
	"auth/tokens"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func setJwtCookies(ctx *gin.Context, accessToken string, refreshToken string) {
	http.SetCookie(ctx.Writer, &http.Cookie{Name: "access_token", Value: accessToken, Expires: time.Now().Add(time.Hour * 24), Path: "/"})
	http.SetCookie(ctx.Writer, &http.Cookie{Name: "refresh_token", Value: refreshToken, Expires: time.Now().Add(time.Hour * 24), Path: "/"})
}

func respondWithHello(ctx *gin.Context, status int, login string) {
	ctx.JSON(status, gin.H{
		"text": fmt.Sprintf("Hello, %s", login),
	})
}

func (a *Adapter) GetJwtTokens(ctx *gin.Context) (*jwt4.Token, *jwt4.Token, error) {
	accessTokenRaw, refreshTokenRaw, err := getRawTokens(ctx)
	if err != nil {
		return nil, nil, err
	}
	return parseTokens(accessTokenRaw, refreshTokenRaw, a.cfg.JwtSecret)
}

func getRawTokens(ctx *gin.Context) (string, string, error) {
	accessTokenRaw, errAccess := ctx.Cookie("access_token")
	refreshTokenRaw, errRefresh := ctx.Cookie("refresh_token")

	if errAccess != nil || errRefresh != nil {
		return "", "", errors.New("jwt-tokens not found: make login")
	}
	return accessTokenRaw, refreshTokenRaw, nil
}

func parseTokens(accessTokenRaw string, refreshTokenRaw string, secret string) (*jwt4.Token, *jwt4.Token, error) {
	accessToken, errAccess := tokens.ParseToken(accessTokenRaw, secret)
	refreshToken, errRefresh := tokens.ParseToken(refreshTokenRaw, secret)

	if errAccess != nil || errRefresh != nil {
		return nil, nil, errors.New("jwt-tokens broken: make login")
	}

	return accessToken, refreshToken, nil
}
