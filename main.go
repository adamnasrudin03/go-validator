package main

import (
	"log"
	"net/http"

	"github.com/adamnasrudin03/go-validator/controllers"
	"github.com/adamnasrudin03/go-validator/helpers/response"
	"github.com/go-playground/validator/v10"
)

func main() {
	port := "1200"
	mux := http.NewServeMux()
	validator := validator.New()
	userController := controllers.NewUserController(validator)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		defer response.PanicRecover(w, r)

		switch {
		case r.URL.Path == "/" && r.Method == "GET":
			response.APIResponse(w, r, "Welcome my application", http.StatusOK)

		case r.URL.Path == "/custom-validation" && r.Method == "POST":
			userController.HandlerCustomValidation(w, r)

		default:
			response.APIResponseNotFound(w, r)
		}
	})

	log.Printf("Server is running on port %v\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
