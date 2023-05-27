package connection

import (
	"chat/infra/headers"
	"chat/logger"
	"context"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/juju/zaputil/zapctx"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type ChatServer struct {
	server   *http.Server
	listener net.Listener
	comments headers.Comments
	log      logger.Logger
	config   ServerConfig
}

type ServerConfig struct {
	Port    int    `env:"HTTP_PORT" envDefault:"3002"`
	AuthUrl string `env:"AUTH_URL"`
}

func NewChatServer(comments headers.Comments, log logger.Logger) (*ChatServer, error) {
	var config ServerConfig
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("parse server http adapter configuration failed: %w", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return nil, fmt.Errorf("server start failed: %w", err)
	}

	router := gin.Default()
	server := http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	chatServer := ChatServer{
		server:   &server,
		listener: listener,
		comments: comments,
		log:      log,
		config:   config,
	}

	initRouter(&chatServer, router)

	return &chatServer, nil
}

func (c *ChatServer) Start() error {
	var err error
	go func() {
		err = c.server.Serve(c.listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			err = fmt.Errorf("server start failed: %w", err)
		}
		err = nil
	}()

	if err != nil {
		return err
	}
	return nil
}

func (c *ChatServer) Stop(ctx context.Context) error {
	var (
		err  error
		once sync.Once
	)
	once.Do(func() {
		err = c.server.Shutdown(ctx)
	})
	return err
}

func (c *ChatServer) authMiddleWare(ctx *gin.Context) {
	req, err := http.NewRequest("GET", c.config.AuthUrl, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"text": err.Error(),
		})
		return
	}
	addCookieToReq(ctx, req, "access_token", "/gigachat")
	addCookieToReq(ctx, req, "refresh_token", "/gigachat")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"text": err.Error(),
		})
		return
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"text": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"text": string(body),
		})
		return
	}

	for _, cookie := range resp.Cookies() {
		ctx.Request.AddCookie(cookie)
	}
	ctx.Next()
}

func addCookieToReq(ctx *gin.Context, req *http.Request, cookieName, path string) {
	if cookieValue, err := ctx.Cookie(cookieName); err == nil {
		req.AddCookie(&http.Cookie{Name: cookieName, Value: cookieValue, Path: path})
	}
}

func (c *ChatServer) getComments(ctx *gin.Context) {
	parentId, err := getIDFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"text": err.Error(),
		})
		return
	}

	comments, err := c.comments.GetComments(parentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"text": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func getIDFromQuery(ctx *gin.Context) (int, error) {
	parentIdQuery := ctx.Query("id")
	if len(parentIdQuery) == 0 {
		return 0, errors.New("expected id")
	}
	return strconv.Atoi(parentIdQuery)
}

func initRouter(c *ChatServer, r *gin.Engine) {
	r.Use(loggerContext(c.log))
	r.Use(corsMiddleware())
	r.Use(logger.SetCookieMiddleware(c.log))
	r.Use(ginzap.Ginzap(c.log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(c.log, true))
	r.Use(c.authMiddleWare)

	api := r.Group("/gigachat/comments/api/v1")
	{
		api.GET("get_comments", c.getComments)
	}
}

func loggerContext(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lCtx := zapctx.WithLogger(ctx.Request.Context(), log)
		ctx.Request = ctx.Request.WithContext(lCtx)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		debug := os.Getenv("DEBUG")
		origin := "http://gigachat.site"
		if debug == "True" {
			origin = "http://localhost:3000"
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
