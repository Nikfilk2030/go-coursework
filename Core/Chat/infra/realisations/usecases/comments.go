package usecases

import (
	"chat/infra/headers"
	"fmt"

	"github.com/caarlos0/env"
)

type CommentsUsecase struct {
	config  *Configuration
	storage headers.Storage
}

type Configuration struct {
}

func (commentsUsecase *CommentsUsecase) GetComments(parentId int) (string, error) {
	jsonComments, err := commentsUsecase.storage.GetComments(parentId)
	if err != nil {
		return "", err
	}
	return jsonComments, nil
}

func (config *Configuration) Parse() error {
	if err := env.Parse(config); err != nil {
		return err
	}
	return nil
}

func New(storage headers.Storage) (*CommentsUsecase, error) {
	var config Configuration
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to parse domain config: %w", err)
	}
	return &CommentsUsecase{
		storage: storage,
	}, nil
}
