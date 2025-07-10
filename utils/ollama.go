package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"student-api/models"
)

const (
	// OllamaURL is the URL for the Ollama API
	OllamaURL = "http://localhost:11434/api/generate"
)

// OllamaRequest represents the request body for Ollama API
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// OllamaResponse represents the response from Ollama API
type OllamaResponse struct {
	Model     string `json:"model"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
	Context   []int  `json:"context,omitempty"`
	TotalDuration int64 `json:"total_duration,omitempty"`
}

// GenerateStudentSummary generates a summary for a student using Ollama
func GenerateStudentSummary(student models.Student) (string, error) {
	// Create the prompt for Ollama
	prompt := fmt.Sprintf(
		"Generate a brief summary for a student with the following information:\n"+
			"Name: %s\n"+
			"Age: %d\n"+
			"Email: %s\n"+
			"Please provide a concise professional summary in 2-3 sentences.",
		student.Name, student.Age, student.Email)

	// Create the request body
	reqBody := OllamaRequest{
		Model:  "llama3",
		Prompt: prompt,
	}

	// Convert request body to JSON
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", OllamaURL, bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Ollama: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error from Ollama API (status %d): %s", resp.StatusCode, string(body))
	}

	// Read and parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Ollama returns multiple JSON objects, one per line
	// We'll collect all responses and combine them
	lines := bytes.Split(body, []byte("\n"))
	var summary string

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var ollamaResp OllamaResponse
		if err := json.Unmarshal(line, &ollamaResp); err != nil {
			return "", fmt.Errorf("error unmarshaling response: %w", err)
		}

		summary += ollamaResp.Response
	}

	return summary, nil
}
