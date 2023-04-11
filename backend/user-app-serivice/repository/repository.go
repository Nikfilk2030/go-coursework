package posts

import "context"

type Post struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	// ...
}

type Repository interface {
	ReadAll(ctx context.Context) ([]Posts, error)
	Upsert(ctx context.Context, post Post) error
	Delete(ctx context.Context, id uuid.UUID) error
}

/*
В итоге проект будет выглядеть как-то так
|- internal
|     |-posts
|     |       |- controller.go
|     |       |- service.go
|     |       |- repository.go
|     |-db
|          |- posts_storage_pg.go
|- main.go
*/
