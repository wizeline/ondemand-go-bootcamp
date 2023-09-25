package controller

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service"

	"github.com/go-chi/render"
)

const (
	repoCsvErrType     errType = "RepositoryCSVError"
	repoDataApiErrType errType = "RepositoryDataAPIError"
	svcFilterErrType   errType = "ServiceFilterError"
)

var _ fmt.Stringer = errType("")

type errHTTP struct {
	Code      int     `json:"code"`
	ErrorType errType `json:"status"`
	Message   string  `json:"message"`
}

type errType string

func (e errType) String() string {
	return string(e)
}

func errJSON(w http.ResponseWriter, r *http.Request, err error) {
	errHttp := newErrHTTP(err)
	render.Status(r, errHttp.Code)
	render.JSON(w, r, errHttp)
}

func newErrHTTP(err error) errHTTP {
	var (
		repoCsvErr     *repository.CsvErr
		repoDataApiErr *repository.DataApiErr
		svcFilterErr   *service.FilterErr
	)

	switch {

	// ###########  REPOSITORY ERRORS ###########

	case errors.As(err, &repoCsvErr):
		return errHTTP{
			Code:      http.StatusInternalServerError,
			ErrorType: repoCsvErrType,
			Message:   err.Error(),
		}
	case errors.As(err, &repoDataApiErr):
		return errHTTP{
			Code:      http.StatusBadGateway,
			ErrorType: repoDataApiErrType,
			Message:   err.Error(),
		}

	// ########### SERVICE ERRORS ###########

	case errors.As(err, &svcFilterErr):
		return errHTTP{
			Code:      http.StatusUnprocessableEntity,
			ErrorType: svcFilterErrType,
			Message:   err.Error(),
		}

	// ########### CONTROLLER ERRORS ###########

	// ########### DEFAULT ERRORS ###########

	default:
		return errHTTP{
			Code:      http.StatusBadRequest,
			ErrorType: errType(reflect.TypeOf(err).String()),
			Message:   err.Error(),
		}
	}
}
