package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
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
	var repoCsvErr *repository.CsvErr

	switch {

	// ###########  REPOSITORY ERRORS ###########
	case errors.As(err, &repoCsvErr):
		return errHTTP{
			Code:      http.StatusInternalServerError,
			ErrorType: "RepositoryCSVError",
			Message:   err.Error(),
		}
	// TODO: migrate this validation to the controller
	//case errors.Is(err, repository.ErrInvalidFilter):
	//	return errHTTP{
	//		Code:      http.StatusBadRequest,
	//		ErrorType: "RepositoryError",
	//		Message:   err.Error(),
	//	}

	// ########### SERVICE ERRORS ###########

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
