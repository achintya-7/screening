package v1

import db "screening/db/sqlc"

type RouteHandler struct {
	store db.Store
}

func NewRouteHandler(store db.Store) *RouteHandler {
	return &RouteHandler{
		store: store,
	}
}
