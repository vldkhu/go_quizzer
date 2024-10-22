package handlers

import (
	"fmt"
	"go_quizzer/internal/models"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GetRandomQuestions(sections []models.Section, n int, questionsPerSection int) []models.Section {
	// Инициализация генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Перемешивание секций
	r.Shuffle(len(sections), func(i, j int) {
		sections[i], sections[j] = sections[j], sections[i]
	})

	// Ограничение количества секций
	if n > len(sections) {
		n = len(sections)
	}

	// Создание нового среза для случайных секций с вопросами
	randomSections := make([]models.Section, n)

	for i := 0; i < n; i++ {
		// Перемешивание вопросов в секции
		r.Shuffle(len(sections[i].Questions), func(a, b int) {
			sections[i].Questions[a], sections[i].Questions[b] = sections[i].Questions[b], sections[i].Questions[a]
		})

		// Ограничение количества вопросов в секции
		if questionsPerSection > len(sections[i].Questions) {
			questionsPerSection = len(sections[i].Questions)
		}

		// Добавление случайных вопросов в новую секцию
		randomSections[i] = models.Section{
			Title:     sections[i].Title,
			Questions: sections[i].Questions[:questionsPerSection],
		}
	}

	return randomSections
}

func StartQuiz(sections []models.Section) {
	fmt.Println("Тестирование по общим знаниям программирования.")
	fmt.Println("*****")
	for {
		selectedQuestions := GetRandomQuestions(sections, 3, 5) // Укажите количество секций и вопросов на секцию
		score := 0

		for _, section := range selectedQuestions {
			for _, q := range section.Questions {
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
		}

		// После завершения теста
		totalQuestions := 0 // Общее количество вопросов
		for _, section := range selectedQuestions {
			totalQuestions += len(section.Questions) // Считаем количество вопросов в каждой секции
		}

		percentage := (float64(score) / float64(totalQuestions)) * 100 // Учитываем фактическое количество вопросов
		fmt.Println("*****")
		fmt.Printf("Ваш итоговый счет: %d из %d\n", score, totalQuestions)
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
