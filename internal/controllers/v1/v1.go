package v1

import (
	db "screening/db/sqlc"
	v1 "screening/internal/handlers/v1"
	"screening/pkg/util"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *v1.RouteHandler
	store   db.Store
}

func NewRouter(store db.Store) *Router {
	return &Router{
		handler: v1.NewRouteHandler(store),
	}
}

func (r *Router) RegisterRoutes(router *gin.RouterGroup) {
	v1Router := router.Group("/v1")

	// USER ROUTES
	v1Router.POST("/user", util.HandlerWrapper[string](r.handler.CreateUser))
	v1Router.PUT("/user", util.HandlerWrapper[string](r.handler.UpdateUser))
}
