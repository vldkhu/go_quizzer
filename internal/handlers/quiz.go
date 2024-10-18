package handlers

import (
	"fmt"
	"go_quizzer/internal/models"
	"math/rand"
	"os"
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
	fmt.Println("Тестирование по общим знаниям программирования.")
	fmt.Println("*****")
	for {
		selectedQuestions := GetRandomQuestions(models.Questions, 20)
		score := 0

		for _, q := range selectedQuestions {
			fmt.Println(q.Text)
			for i, option := range q.Options {
				fmt.Printf("%d: %s\n", i+1, option)
			}

			var userAnswer string
			fmt.Println("*****")
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
				fmt.Println("*****************************")
			} else {
				fmt.Printf("Неправильно! Правильный ответ: %s\n", correctAnswer)
				fmt.Println("*****************************")
			}
		}
		percentage := (float64(score) / float64(len(selectedQuestions))) * 100
		fmt.Println("*****")
		fmt.Printf("Ваш итоговый счет: %d из %d\n", score, len(selectedQuestions))
		fmt.Printf("Процент верных ответов: %.2f%%\n", percentage)
		fmt.Println("*****")

		fmt.Println("Тест завершен.")
		var choice string
		fmt.Println("Введите 'y' для выхода или 'n' для перезапуска теста:")
		fmt.Scan(&choice)

		if choice == "y" {
			os.Exit(0)
		} else if choice == "n" {
			fmt.Println("Перезапуск приложения...")
		} else {
			fmt.Println("Неверный ввод...")
		}
	}
}
