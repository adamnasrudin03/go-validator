package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/adamnasrudin03/go-validator/dto"
	"github.com/adamnasrudin03/go-validator/helpers/response"
	help_validator "github.com/adamnasrudin03/go-validator/helpers/validator"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	HandlerCustomValidation(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	Validate *validator.Validate
}

func NewUserController(validate *validator.Validate) UserController {
	return &userHandler{
		Validate: validate,
	}
}

func (h *userHandler) HandlerCustomValidation(w http.ResponseWriter, r *http.Request) {
	defer response.PanicRecover(w, r)
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.APIResponse(w, r, errors.New("invalid request"), http.StatusBadRequest)
		return
	}

	input := dto.UserRequest{}
	err = json.Unmarshal(reqBody, &input)
	if err != nil {
		response.APIResponse(w, r, errors.New("invalid request"), http.StatusBadRequest)
		return
	}

	// validate input
	err = h.Validate.Struct(input)
	if err != nil {
		response.APIResponse(w, r, help_validator.FormatValidationError(err), http.StatusBadRequest)
		return
	}

	response.APIResponse(w, r, input, http.StatusOK)
}
