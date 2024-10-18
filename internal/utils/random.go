package utils

import (
	"go_quizzer/internal/models"
	"math/rand"
	"time"
)

func ShuffleQuestions(questions []models.Question) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
