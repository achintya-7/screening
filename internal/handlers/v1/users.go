package v1

import (
	"database/sql"
	"net/http"
	db "screening/db/sqlc"
	"screening/internal/constants"
	"screening/internal/dto"
	"screening/pkg/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (rh *RouteHandler) CreateUser(c *gin.Context) (*string, *dto.ErrorResponse) {
	logger := util.NewLoggerWithCorrelationId(c)

	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("error while binding request", zap.Error(err))
		return nil, &dto.ErrorResponse{
			Code:           http.StatusBadRequest,
			HttpStatusCode: http.StatusBadRequest,
			Message:        err.Error(),
		}
	}

	args := db.InsertUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	logger.Info("inserting user")
	_, err := rh.store.InsertUser(c, args)
	if err != nil {
		logger.Error("error while inserting user", zap.Error(err))
		return nil, &dto.ErrorResponse{
			Code:           http.StatusInternalServerError,
			HttpStatusCode: http.StatusInternalServerError,
			Message:        err.Error(),
		}
	}

	return &constants.CREATE_USER_SUCCESS_MSG, nil
}

func (rh *RouteHandler) UpdateUser(c *gin.Context) (*string, *dto.ErrorResponse) {
	logger := util.NewLoggerWithCorrelationId(c)

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("error while binding request", zap.Error(err))
		return nil, &dto.ErrorResponse{
			Code:           http.StatusBadRequest,
			HttpStatusCode: http.StatusBadRequest,
			Message:        err.Error(),
		}
	}

	args := db.UpdateUserParams{
		Name:  req.Name,
		Email: req.Email,
		ID:    req.Id,
	}

	logger.Info("updating user")
	_, err := rh.store.UpdateUser(c, args)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("user not found", zap.Error(err))
			return nil, &dto.ErrorResponse{
				Code:           http.StatusNotFound,
				HttpStatusCode: http.StatusNotFound,
				Message:        err.Error(),
			}
		}

		logger.Error("error while updating user", zap.Error(err))
		return nil, &dto.ErrorResponse{
			Code:           http.StatusInternalServerError,
			HttpStatusCode: http.StatusInternalServerError,
			Message:        err.Error(),
		}
	}

	return &constants.UPDATE_USER_SUCCESS_MSG, nil
}
