package test

import (
	"go_quizzer/internal/handlers"
	"go_quizzer/internal/models"
	"testing"
)

func TestGetRandomQuestions(t *testing.T) {
	questions := models.Questions
	selected := handlers.GetRandomQuestions(questions, 20)

	if len(selected) != 20 {
		t.Errorf("Expected 5 questions, got %d", len(selected))
	}
}
