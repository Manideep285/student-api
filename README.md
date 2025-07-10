# Student API

A simple REST API built with Go that performs CRUD operations on student records and integrates with Ollama for AI-generated student summaries.

## Project Structure

```
student-api/
├── go.mod             # Go module definition
├── main.go            # Main application entry point
├── handlers/          # HTTP request handlers
│   ├── errors.go      # Error definitions
│   └── student_handlers.go # Student-related handlers
├── models/            # Data models
│   └── student.go     # Student model and store
└── utils/             # Utility functions
    └── ollama.go      # Ollama API integration
```

## Prerequisites

- Go 1.16 or higher
- [Ollama](https://github.com/ollama/ollama) installed locally
- Llama3 model pulled in Ollama (`ollama pull llama3`)

## Installation

1. Clone the repository
2. Install dependencies:
   ```
   go mod download
   ```
3. Build the application:
   ```
   go build
   ```
4. Run the application:
   ```
   ./student-api
   ```

The server will start on port 8080.

## API Endpoints

### Create a new student
```
POST /students
```
Example request:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe","age":20,"email":"john@example.com"}' http://localhost:8080/students
```

### Get all students
```
GET /students
```
Example request:
```bash
curl http://localhost:8080/students
```

### Get a student by ID
```
GET /students/{id}
```
Example request:
```bash
curl http://localhost:8080/students/1
```

### Update a student by ID
```
PUT /students/{id}
```
Example request:
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"John Smith","age":21,"email":"john.smith@example.com"}' http://localhost:8080/students/1
```

### Delete a student by ID
```
DELETE /students/{id}
```
Example request:
```bash
curl -X DELETE http://localhost:8080/students/1
```

### Generate a summary of a student by ID using Ollama
```
GET /students/{id}/summary
```
Example request:
```bash
curl http://localhost:8080/students/1/summary
```

## Input Validation

The API validates the following:
- Name is required
- Age must be a positive number
- Email is required

## Concurrency

The API uses a mutex to ensure thread-safe access to the student data store, making it safe for concurrent requests.

## Ollama Integration

The API integrates with Ollama to generate AI-based summaries of student profiles. Make sure Ollama is running locally and the Llama3 model is available before using the summary endpoint.
