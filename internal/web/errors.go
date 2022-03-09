package web

import (
	"github.com/AlexRipoll/go-pod/internal/sys/errorFlag"
	"net/http"
)

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// ErrorMatch checks the error's flag and creates a REST error response to return for the API request.
func ErrorMatch(err error) ErrorResponse {
	var er ErrorResponse
	flaggedErr := err.(errorFlag.Flagged)

	switch flaggedErr.Flag() {
	case errorFlag.InvalidData, errorFlag.AlreadyExists:
		er = ErrorResponse{
			Message: flaggedErr.Unwrap().Error(),
			Status:  http.StatusBadRequest,
		}
	case errorFlag.NotFound:
		er = ErrorResponse{
			Message: flaggedErr.Unwrap().Error(),
			Status:  http.StatusNotFound,
		}
	case errorFlag.Internal:
		er = ErrorResponse{
			Message: flaggedErr.Unwrap().Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return er
}
