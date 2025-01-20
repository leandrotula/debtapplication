package service

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrorResponse(message string, status int) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Status:  status,
	}
}
