package posts

import "net/http"

type PostsController struct {
	// тут могут быть различные зависимости, которые требуются для обработки запросов
	// например, сервисы о которых речь пойдёт ниже
	postsService *Service
}

func (c *PostsController) PostsGet(r *http.Request, rw http.ResponseWriter) {
	// обработка запроса GET /api/v1/posts
}

func (c *PostsController) PostsPost(r *http.Request, rw http.ResponseWriter) {
	// обработка запроса POST /api/v1/posts
}

func (c *PostsController) PostsDelete(r *http.Request, rw http.ResponseWriter) {
	// обработка запроса DELETE /api/v1/posts/{id}
}
