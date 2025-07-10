package models

import (
	"errors"
	"sync"
)

// Student represents a student entity
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// StudentStore manages the student data
type StudentStore struct {
	sync.RWMutex
	students map[int]Student
	nextID   int
}

// NewStudentStore creates a new student store
func NewStudentStore() *StudentStore {
	return &StudentStore{
		students: make(map[int]Student),
		nextID:   1,
	}
}

// GetAll returns all students
func (s *StudentStore) GetAll() []Student {
	s.RLock()
	defer s.RUnlock()

	students := make([]Student, 0, len(s.students))
	for _, student := range s.students {
		students = append(students, student)
	}
	return students
}

// Get returns a student by ID
func (s *StudentStore) Get(id int) (Student, error) {
	s.RLock()
	defer s.RUnlock()

	student, exists := s.students[id]
	if !exists {
		return Student{}, errors.New("student not found")
	}
	return student, nil
}

// Create adds a new student
func (s *StudentStore) Create(student Student) Student {
	s.Lock()
	defer s.Unlock()

	student.ID = s.nextID
	s.students[student.ID] = student
	s.nextID++
	return student
}

// Update updates an existing student
func (s *StudentStore) Update(id int, student Student) (Student, error) {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.students[id]; !exists {
		return Student{}, errors.New("student not found")
	}

	student.ID = id
	s.students[id] = student
	return student, nil
}

// Delete removes a student
func (s *StudentStore) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.students[id]; !exists {
		return errors.New("student not found")
	}

	delete(s.students, id)
	return nil
}
