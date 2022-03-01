package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/martinsmessias/go_rest_api/controllers"
	"github.com/martinsmessias/go_rest_api/middleware"
)

// HandleRequests handles all incoming requests and routes them to the appropriate controller
func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonality).Methods("GET")
	r.HandleFunc("/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
