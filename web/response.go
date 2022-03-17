package web

import (
	"context"
	"encoding/json"
	"github.com/AlexRipoll/go-pod/logger"
	"net/http"
)

// Response parses the data to JSON and sets the HTTP Headers and sends it to the client.
func Response(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		//return errorFlag.New(err, errorFlag.Internal)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}

// ErrorResponse checks if the received error has errorFlag.flagged type, parses to JSON and
// returns it to the client.
func ErrorResponse(ctx context.Context, log *logger.Logger, w http.ResponseWriter, error error) {

	errorData := ErrorMatch(ctx, log, error)

	jsonData, err := json.Marshal(errorData)
	if err != nil {
		//return errorFlag.New(err, errorFlag.Internal)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(errorData.Status)

	if _, err := w.Write(jsonData); err != nil {
		return
	}

	return
}
