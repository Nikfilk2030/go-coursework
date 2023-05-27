package connection

import (
	"auth/tokens"
	"fmt"
	"github.com/gin-gonic/gin"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func (a *Adapter) handleGetLogin(ctx *gin.Context) {
	login, password, authExists := ctx.Request.BasicAuth()
	if !authExists {
		respondWithError(ctx, http.StatusBadRequest, "Basic-Auth expected")
		return
	}
	if err := a.auth.Login(login, password); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	a.createAndSetTokens(ctx, login)
}

func (a *Adapter) handleGetAuth(ctx *gin.Context) {
	accessToken, refreshToken, err := a.GetJwtTokens(ctx)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err = a.auth.Auth(); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if accessToken != nil {
		login := accessToken.Claims.(jwt4.MapClaims)["sub"].(string)
		respondWithHello(ctx, http.StatusOK, login)
		setJwtCookies(ctx, accessToken.Raw, refreshToken.Raw)
	} else if refreshToken != nil {
		login := refreshToken.Claims.(jwt4.MapClaims)["sub"].(string)
		accessTokenNew, err := tokens.CreateToken(login, time.Now().Add(time.Minute), a.cfg.JwtSecret)
		if err != nil {
			respondWithError(ctx, http.StatusForbidden, err.Error())
			return
		}
		refreshTokenNew, err := tokens.CreateToken(login, time.Now().Add(time.Hour), a.cfg.JwtSecret)
		if err != nil {
			respondWithError(ctx, http.StatusForbidden, err.Error())
			return
		}
		setJwtCookies(ctx, accessTokenNew, refreshTokenNew)
		respondWithHello(ctx, http.StatusOK, login)
	} else {
		respondWithError(ctx, http.StatusForbidden, "make login")
	}
}

func (a *Adapter) handlePostRegister(ctx *gin.Context) {
	login, password, authExists := ctx.Request.BasicAuth()
	if !authExists {
		respondWithError(ctx, http.StatusBadRequest, "Basic-Auth header expected")
		return
	}
	if err := a.auth.Register(login, password); err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	respondWithSuccess(ctx, http.StatusOK, fmt.Sprintf("Success register, %s! :)", login))
}

func (a *Adapter) createAndSetTokens(ctx *gin.Context, login string) {
	accessToken, refreshToken, err := createTokens(login, a.cfg.JwtSecret)
	if err != nil {
		respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	setCookies(ctx, accessToken, refreshToken)
	respondWithSuccess(ctx, http.StatusOK, fmt.Sprintf("Hello, %s", login))
}

func createTokens(login string, secret string) (string, string, error) {
	accessToken, err := tokens.CreateToken(login, time.Now().Add(time.Minute), secret)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := tokens.CreateToken(login, time.Now().Add(time.Hour), secret)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func setCookies(ctx *gin.Context, accessToken string, refreshToken string) {
	http.SetCookie(ctx.Writer, &http.Cookie{Name: "access_token", Value: accessToken, Expires: time.Now().Add(time.Hour * 24), Path: "/"})
	http.SetCookie(ctx.Writer, &http.Cookie{Name: "refresh_token", Value: refreshToken, Expires: time.Now().Add(time.Hour * 24), Path: "/"})
}

func respondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"text": message,
	})
}

func respondWithSuccess(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"text": message,
	})
}
