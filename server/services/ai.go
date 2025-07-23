package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"google.golang.org/genai"
)

type OpenAIRequest struct {
	Model    string      `json:"model"`
	Messages []AIMessage `json:"messages"`
}

type AIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message AIMessage `json:"message"`
	} `json:"choices"`
}

func GetAIInsights(title, description, affectedService string) (string, string, error) {
	prompt := fmt.Sprintf(`
	You are an incident triage assistant. Given the following incident details, classify:

- Severity: One of ["Low", "Medium", "High", "Critical"]
- Category: One of ["Network", "Software", "Hardware", "Security"]

Respond strictly in the following JSON format:
{
  "severity": "<one of the above>",
  "category": "<one of the above>"
}
Incident Title: %s
Incident Description: %s
Affected Service: %s
`, title, description, affectedService)

	ctx := context.Background()
	// The client gets the API key from the environment variable `GEMINI_API_KEY`.
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", "", err
	}
	if len(result.Candidates) == 0 {
		return "", "", fmt.Errorf("empty response")
	}
	// 1. Extract the text content (adjust field access as per actual SDK)
	text := result.Candidates[0].Content.Parts[0].Text

	// 2. Unmarshal the JSON string
	cleaned := strings.TrimSpace(text)
	cleaned = strings.TrimPrefix(cleaned, "```json")
	cleaned = strings.TrimPrefix(cleaned, "```")
	cleaned = strings.TrimSuffix(cleaned, "```")

	var response map[string]string
	err = json.Unmarshal([]byte(cleaned), &response)
	if err != nil {
		return "", "", err
	}

	return response["severity"], response["category"], nil
}
