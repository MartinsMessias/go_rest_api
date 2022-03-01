package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/martinsmessias/go_rest_api/database"
	"github.com/martinsmessias/go_rest_api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// Create a new personality
func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	_ = json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

// Return all the personalities
func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

// Search for a personality by ID and return it
func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var p models.Personality

	database.DB.First(&p, key)
	json.NewEncoder(w).Encode(p)
}

// Delete a personality by ID
func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var p models.Personality

	database.DB.First(&p, key)
	database.DB.Delete(&p)
	json.NewEncoder(w).Encode(p)
}

// Update a personality by ID
func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var p models.Personality

	database.DB.First(&p, key)
	_ = json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
