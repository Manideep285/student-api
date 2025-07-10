package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"student-api/models"
	"student-api/utils"

	"github.com/gorilla/mux"
)

// StudentHandler handles student-related HTTP requests
type StudentHandler struct {
	store *models.StudentStore
}

// NewStudentHandler creates a new student handler
func NewStudentHandler(store *models.StudentStore) *StudentHandler {
	return &StudentHandler{
		store: store,
	}
}

// GetStudents returns all students
func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students := h.store.GetAll()
	respondWithJSON(w, http.StatusOK, students)
}

// GetStudent returns a student by ID
func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	student, err := h.store.Get(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Student not found")
		return
	}

	respondWithJSON(w, http.StatusOK, student)
}

// CreateStudent creates a new student
func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate student data
	if err := validateStudent(student); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdStudent := h.store.Create(student)
	respondWithJSON(w, http.StatusCreated, createdStudent)
}

// UpdateStudent updates a student by ID
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Validate student data
	if err := validateStudent(student); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	updatedStudent, err := h.store.Update(id, student)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Student not found")
		return
	}

	respondWithJSON(w, http.StatusOK, updatedStudent)
}

// DeleteStudent deletes a student by ID
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	if err := h.store.Delete(id); err != nil {
		respondWithError(w, http.StatusNotFound, "Student not found")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// GetStudentSummary generates a summary for a student using Ollama
func (h *StudentHandler) GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	student, err := h.store.Get(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Student not found")
		return
	}

	summary, err := utils.GenerateStudentSummary(student)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate summary: "+err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"summary": summary})
}

// Helper functions

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func validateStudent(student models.Student) error {
	if student.Name == "" {
		return ErrNameRequired
	}
	if student.Age <= 0 {
		return ErrInvalidAge
	}
	if student.Email == "" {
		return ErrEmailRequired
	}
	return nil
}
