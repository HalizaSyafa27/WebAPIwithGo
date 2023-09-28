package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define your API endpoints here
	r.HandleFunc("/api/users", getUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUserByIDHandler).Methods("GET")
	r.HandleFunc("/api/users", createUserHandler).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUserHandler).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get all users")
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "Get user with ID: %s", userID)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user")
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "Update user with ID: %s", userID)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	vars := mux.Vars(r)
	userID := vars["id"]

	fmt.Fprintf(w, "Delete user with ID: %s", userID)
}
