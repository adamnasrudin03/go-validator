package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Response struct {
	Message interface{} `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(input interface{}) (resp Response) {
	switch input := input.(type) {
	case []string:
		resp = Response{
			Errors: input,
		}
	case error:
		resp = Response{
			Errors: input.Error(),
		}
	case string:
		resp = Response{
			Message: input,
		}
	default:
		resp = Response{
			Data: input,
		}
	}

	return resp
}

func APIResponse(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	jsonBytes, err := json.Marshal(NewResponse(data))
	if err != nil {
		APIResponseInternalServerError(w, r)
		return
	}
	w.WriteHeader(code)
	w.Write(jsonBytes)
}

func APIResponseInternalServerError(w http.ResponseWriter, r *http.Request) {
	APIResponse(w, r, errors.New(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
}

func APIResponseNotFound(w http.ResponseWriter, r *http.Request) {
	APIResponse(w, r, errors.New(http.StatusText(http.StatusNotFound)), http.StatusNotFound)
}

func PanicRecover(w http.ResponseWriter, r *http.Request) {
	if rc := recover(); rc != nil {
		APIResponseInternalServerError(w, r)
	}
}
