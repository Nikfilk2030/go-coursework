package postgres

import (
	"chat/infra/realisations/comment_service"
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v4"

	"github.com/caarlos0/env"
)

type Config struct {
	PostgresURL string `env:"POSTGRES_URL"`
}

type PostgresStorage struct {
	cfg Config
}

func New() (*PostgresStorage, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("parsing postgres adapter configuration failed: %w", err)
	}
	return &PostgresStorage{
		cfg: cfg,
	}, nil
}

func (s *PostgresStorage) Stop(ctx context.Context) error {
	// Placeholder for stop process
	return nil
}

func (s *PostgresStorage) GetComments(parentId int) (string, error) {
	query := fmt.Sprintf("SELECT * FROM comments WHERE parent = %d;", parentId)
	db, err := pgx.Connect(context.Background(), s.cfg.PostgresURL)
	if err != nil {
		return "", err
	}
	defer db.Close(context.Background())
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return "", err
	}
	var comments []*comment_service.Comment
	for rows.Next() {
		comment := new(comment_service.Comment)
		err = rows.Scan(
			&comment.Id,
			&comment.OwnerLogin,
			&comment.CommentType,
			&comment.Text,
			&comment.Parent,
		)
		if err != nil {
			return "", err
		}
		comments = append(comments, comment)
	}
	jsonBytes, err := json.Marshal(comments)
	fmt.Println(string(jsonBytes), comments)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
