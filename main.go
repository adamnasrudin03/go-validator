package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/adamnasrudin03/go-validator/dto"
	"github.com/adamnasrudin03/go-validator/helpers"
)

func ValidateInput(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.APIResponse(w, r, errors.New("invalid request"), http.StatusBadRequest)
	}

	input := dto.User{}
	err = json.Unmarshal(reqBody, &input)
	if err != nil {
		helpers.APIResponse(w, r, errors.New("invalid request"), http.StatusBadRequest)
	}

	helpers.APIResponse(w, r, input, http.StatusOK)
}

func main() {
	port := "1200"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		switch {
		case r.URL.Path == "/" && r.Method == "GET":
			helpers.APIResponse(w, r, "Welcome my application", http.StatusOK)

		case r.URL.Path == "/default-validation" && r.Method == "POST":
			ValidateInput(w, r)

		case r.URL.Path == "/custom-validation" && r.Method == "POST":
			ValidateInput(w, r)

		default:
			helpers.APIResponseNotFound(w, r)
		}
	})

	log.Printf("Server is running on port %v\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
