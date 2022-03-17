package web

import (
	"context"
	"github.com/AlexRipoll/go-pod/internal/sys/errorFlag"
	"github.com/AlexRipoll/go-pod/logger"
	"net/http"
)

// Error is the form used for API responses from failures in the API.
type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// ErrorMatch checks the error's flag and creates a REST error response to return for the API request.
func ErrorMatch(ctx context.Context, log *logger.Logger, err error) Error {
	var er Error
	flaggedErr := err.(errorFlag.Flagged)

	switch flaggedErr.Flag() {
	case errorFlag.InvalidData, errorFlag.AlreadyExists:
		er = Error{
			Message: flaggedErr.Unwrap().Error(),
			Status:  http.StatusBadRequest,
		}
	case errorFlag.NotFound:
		er = Error{
			Message: flaggedErr.Unwrap().Error(),
			Status:  http.StatusNotFound,
		}
	case errorFlag.Internal:
		er = Error{
			Message: "internal server error",
			//Message: flaggedErr.Unwrap().Error(),
			Status:  http.StatusInternalServerError,
		}
		log.Error(flaggedErr.Unwrap().Error())
	}

	return er
}
