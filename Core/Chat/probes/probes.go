package probes

import (
	"chat/logger"
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type Probes struct {
	isReady     bool
	isStarted   bool
	readyOnce   sync.Once
	startedOnce sync.Once
	l           logger.Logger
	server      *http.Server
}

type Config struct {
	Port int `env:"PROBES_PORT" envDefault:"3030"`
}

func New(l logger.Logger) (*Probes, error) {
	return &Probes{
		l: l,
	}, nil
}

func (p *Probes) Start() error {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("configuration parsing failed: %w", err)
	}

	if cfg.Port <= 0 || cfg.Port > 65535 {
		return fmt.Errorf("invalid port number: %d", cfg.Port)
	}

	r := gin.Default()
	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusOK)
	})
	r.GET("/ready", func(ctx *gin.Context) {
		if p.isReady {
			ctx.Writer.WriteHeader(http.StatusOK)
		} else {
			ctx.Writer.WriteHeader(http.StatusLocked)
		}
	})
	r.GET("/startup", func(ctx *gin.Context) {
		if p.isStarted {
			ctx.Writer.WriteHeader(http.StatusOK)
		} else {
			ctx.Writer.WriteHeader(http.StatusLocked)
		}
	})

	p.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	go func() {
		if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			p.l.Sugar().Errorf("start probes failed: %w", err)
		}
	}()

	return nil
}

func (p *Probes) SetReady() {
	p.readyOnce.Do(func() {
		p.isReady = true
	})
}

func (p *Probes) SetStarted() {
	p.startedOnce.Do(func() {
		p.isStarted = true
	})
}

func (p *Probes) Shutdown(ctx context.Context) error {
	return p.server.Shutdown(ctx)
}
