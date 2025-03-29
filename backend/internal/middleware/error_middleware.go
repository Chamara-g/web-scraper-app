package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gihanc.dev/web-scraper-app/backend/internal/errors"
	"github.com/gihanc.dev/web-scraper-app/backend/internal/logger"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func WriteErrorResponse(w http.ResponseWriter, errorCode int, message string) {
	errorMessage := "MESSAGE: " + message + ", CODE: " + strconv.Itoa(errorCode)

	logger.Logger.Error(errorMessage, "error", "response")

	w.Header().Set("Content-Type", "application/json")

	err := errors.New(errorCode, message)

	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Message, Code: err.Code})

}
