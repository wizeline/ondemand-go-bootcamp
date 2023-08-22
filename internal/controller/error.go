package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service"
)

type errHTTP struct {
	Code      int    `json:"code"`
	ErrorType string `json:"status"`
	Message   string `json:"message"`
}

func errHTTPResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	errHttp := newErrHTTP(err)
	w.WriteHeader(errHttp.Code)
	_ = json.NewEncoder(w).Encode(errHttp)
}

func newErrHTTP(err error) errHTTP {
	var (
		repoCsvErr   *repository.CsvErr
		svcFilterErr *service.FilterErr
	)

	switch {

	// ###########  REPOSITORY ERRORS ###########
	case errors.As(err, &repoCsvErr):
		return errHTTP{
			Code:      http.StatusInternalServerError,
			ErrorType: "RepositoryCSVError",
			Message:   err.Error(),
		}

	// ########### SERVICE ERRORS ###########
	case errors.As(err, &svcFilterErr):
		return errHTTP{
			Code:      http.StatusUnprocessableEntity,
			ErrorType: "ServiceFilterError",
			Message:   err.Error(),
		}

	// ########### CONTROLLER ERRORS ###########

	// ########### DEFAULT ERRORS ###########

	default:
		return errHTTP{
			Code:      http.StatusBadRequest,
			ErrorType: reflect.TypeOf(err).String(),
			Message:   err.Error(),
		}
	}
}
