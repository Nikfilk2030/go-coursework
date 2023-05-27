package connection

import (
	"auth/infra/headers"
	"auth/logger"
	"context"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"sync"
	"time"
)

type Adapter struct {
	s    *http.Server
	l    net.Listener
	auth headers.Auth
	log  logger.Logger
	cfg  Config
}

type Config struct {
	Port      int    `env:"HTTP_PORT" envDefault:"3001"`
	JwtSecret string `env:"JWT_SECRET"`
}

func New(auth headers.Auth, log logger.Logger) (*Adapter, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse server http adapter configuration failed: %w", err)
	}

	l, err := initListener(cfg.Port)
	if err != nil {
		return nil, fmt.Errorf("server start failed: %w", err)
	}

	router := gin.Default()
	server := initServer(router)

	a := Adapter{
		s:    server,
		l:    l,
		auth: auth,
		log:  log,
		cfg:  cfg,
	}
	initRouter(&a, router)

	return &a, nil
}

func parseConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func initListener(port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf(":%d", port))
}

func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func (a *Adapter) Start() error {
	var err error
	go func() {
		err = a.s.Serve(a.l)
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

func (a *Adapter) Stop(ctx context.Context) error {
	var (
		err  error
		once sync.Once
	)
	once.Do(func() {
		err = a.s.Shutdown(ctx)
	})
	return err
}
