package handlers

import (
	"fmt"
	"go_quizzer/internal/models"
	"math/rand"
	"strconv"
	"time"
)

func GetRandomQuestions(questions []models.Question, n int) []models.Question {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	if n > len(questions) {
		n = len(questions)
	}
	return questions[:n]
}

func StartQuiz() {
	selectedQuestions := GetRandomQuestions(models.Questions, 20)
	score := 0

	for _, q := range selectedQuestions {
		fmt.Println(q.Text)
		for i, option := range q.Options {
			fmt.Printf("%d: %s\n", i+1, option)
		}

		var userAnswer string
		fmt.Print("Ваш ответ (введите номер): ")
		fmt.Scanln(&userAnswer)

		// Преобразуем ответ пользователя в число
		answerIndex, err := strconv.Atoi(userAnswer)
		if err != nil || answerIndex < 1 || answerIndex > len(q.Options) {
			fmt.Println("Неверный ввод. Пожалуйста, введите номер варианта.")
			continue // Пропускаем итерацию, если ввод неверный
		}

		// Проверяем правильность ответа
		correctAnswer := q.Answer
		if q.Options[answerIndex-1] == correctAnswer {
			fmt.Println("Правильно!")
			score++
		} else {
			fmt.Printf("Неправильно! Правильный ответ: %s\n", correctAnswer)
		}
	}
	percentage := (float64(score) / float64(len(selectedQuestions))) * 100
	fmt.Printf("Ваш итоговый счет: %d из %d\n", score, len(selectedQuestions))
	fmt.Printf("Процент верных ответов: %.2f%%\n", percentage)
}
