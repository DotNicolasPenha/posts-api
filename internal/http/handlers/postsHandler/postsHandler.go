package postsHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/middleware"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service *post.Service
	g       *gin.Engine
}

func NewPostHandler(service *post.Service, g *gin.Engine) {
	postsHandler := g.Group("/posts")
	handler := handler{service: service, g: g}
	{
		postsHandler.POST("/", middleware.UserMiddleware(), handler.createPostHandler)
		postsHandler.GET("/", handler.getPostsHandler)
		postsHandler.GET("/:id", handler.getOnePostHandler)
		postsHandler.DELETE("/:id", middleware.UserMiddleware(), handler.deletePostHandler)
	}
}
