package routes

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/handlers"
	"go-cource-api/application/middlewares"
)

func InitApiV1Routes(e *echo.Echo, postsHandlers *handlers.PostHandlers, commentHandlers *handlers.CommentHandlers) {
	apiV1 := e.Group("/api/v1")
	apiV1.Use(middlewares.SecurityContextMiddleware)
	restrictedApiV1 := apiV1.Group("")
	restrictedApiV1.Use(middlewares.AuthMiddleware())

	// Public API
	apiV1.GET("/posts", postsHandlers.List)
	apiV1.GET("/posts/:id", postsHandlers.FindOne)
	apiV1.GET("/posts/:postId/comments", commentHandlers.FindByPostId)

	// Private API
	restrictedApiV1.POST("/posts", postsHandlers.Save)
	restrictedApiV1.DELETE("/posts/:id", postsHandlers.Delete)
	restrictedApiV1.PUT("/posts/:id", postsHandlers.Update)

	restrictedApiV1.POST("/posts/:postId/comments", commentHandlers.Save)
	restrictedApiV1.DELETE("/comments/:id", commentHandlers.Delete)
	restrictedApiV1.PUT("/comments/:id", commentHandlers.Update)
}
