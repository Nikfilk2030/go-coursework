package posts

import "context"

type PostsService struct {
	// тут уже будут репозитории в качестве зависимостей
	postsRepo Repository
}

func (c *PostsService) CreatePost(ctx context.Context, post Post) (uuid.UUID, error) {
	// создаёт новый пост и возвращает его ID или ошибку
}

func (c *PostsService) ReadPosts(ctx context.Context) ([]Post, error) {
	// возвращает список постов или ошибку
}

func (c *PostsController) DeletePost(ctx context.Context, id uuid.UUID) error {
	// удаляет пост по ID
}
