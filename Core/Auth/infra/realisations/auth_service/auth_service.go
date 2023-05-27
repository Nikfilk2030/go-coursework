package auth_service

import (
	"auth/infra/headers"
	"auth/infra/realisations/user"
	"fmt"
	"github.com/caarlos0/env"
)

type Auth struct {
	cfg     *Config
	storage headers.Store
}

type Config struct {
}

func New(storage headers.Store) (*Auth, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse domain config failed: %w", err)
	}
	return &Auth{storage: storage, cfg: &cfg}, nil
}

func (auth *Auth) Login(login string, password string) error {
	_, err := auth.storage.GetUser(login, password)
	return err
}

func (auth *Auth) Auth() error {
	return nil
}

func (auth *Auth) Register(login string, password string) error {
	return auth.storage.InsertUser(user.NewUser(login, password))
}

func (c *Config) Parse() error {
	return env.Parse(c)
}
