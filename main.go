package main

import (
	"fmt"
	"log"
	"net/http"
	"student-api/handlers"
	"student-api/models"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new student store
	store := models.NewStudentStore()

	// Create a new student handler
	studentHandler := handlers.NewStudentHandler(store)

	// Create a new router
	router := mux.NewRouter()

	// Define API routes
	router.HandleFunc("/students", studentHandler.GetStudents).Methods("GET")
	router.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id:[0-9]+}", studentHandler.GetStudent).Methods("GET")
	router.HandleFunc("/students/{id:[0-9]+}", studentHandler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id:[0-9]+}", studentHandler.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id:[0-9]+}/summary", studentHandler.GetStudentSummary).Methods("GET")

	// Add middleware for logging
	router.Use(loggingMiddleware)

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// loggingMiddleware logs all requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
