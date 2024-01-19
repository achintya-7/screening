package dto

type ErrorResponse struct {
	Code           int
	Message        string
	HttpStatusCode int
}

type ApiError struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int
}

type ApiResponse[T any] struct {
	Status bool      `json:"status"`
	Data   *T        `json:"data"`
	Error  *ApiError `json:"error"`
}

type CreateUserRequest struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
}

type UpdateUserRequest struct {
	Name  string `form:"name" binding:"required"`
	Email string `form:"email" binding:"required"`
	Id    int64  `form:"id" binding:"required"`
}
