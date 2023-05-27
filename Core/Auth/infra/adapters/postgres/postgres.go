package postgres

import (
	"auth/infra/headers"
	"auth/infra/realisations/user"
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4"
)

func (s *PostgresStorage) makeQuery(query string) ([]*user.User, error) {
	db, err := pgx.Connect(context.Background(), s.cfg.Postgres_url)
	if err != nil {
		return nil, err
	}
	defer db.Close(context.Background())
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var users []*user.User
	for rows.Next() {
		var login string
		var password string
		err = rows.Scan(
			&login,
			&password,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user.NewUser(login, password))
	}
	return users, nil
}

func (s *PostgresStorage) FindUser(login string) (bool, error) {
	users, err := s.makeQuery(fmt.Sprintf("SELECT * FROM users WHERE login = '%s';", login))
	if err != nil {
		return false, err
	}
	return len(users) == 1, nil
}

func (s *PostgresStorage) GetUser(login string, password string) (headers.User, error) {
	ok, err := s.FindUser(login)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("GetUser: user not found")
	}
	users, err := s.makeQuery(fmt.Sprintf("SELECT * FROM users WHERE login = '%s';", login))
	if err != nil {
		return nil, err
	}
	user := users[0]
	if password != user.GetPassword() {
		return nil, fmt.Errorf("GetUser: invalid password")
	}
	return user, nil
}

func (s *PostgresStorage) InsertUser(user headers.User) error {
	ok, err := s.FindUser(user.GetLogin())
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("InsertUser: user arleady pushed")
	}
	ctx := context.Background()
	db, err := pgx.Connect(context.Background(), s.cfg.Postgres_url)
	if err != nil {
		return err
	}
	defer db.Close(context.Background())
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	batch := new(pgx.Batch)
	batch.Queue("INSERT INTO users (login, password) VALUES ($1, $2);", user.GetLogin(), user.GetPassword())
	res := tx.SendBatch(ctx, batch)
	err = res.Close()
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

type PostgresStorage struct {
	cfg Config
}

type Config struct {
	Postgres_url string `env:"POSTGRES_URL"`
}

func New() (*PostgresStorage, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parse postgres adapter configuration failed: %w", err)
	}
	return &PostgresStorage{
		cfg: cfg,
	}, nil
}

func (s *PostgresStorage) Stop(ctx context.Context) error {
	return nil
}
