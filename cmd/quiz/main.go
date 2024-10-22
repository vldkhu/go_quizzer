package main

import (
	"go_quizzer/internal/handlers"
	"go_quizzer/internal/models"
)

func main() {
	// Получение вопросов из модели
	sections := models.GetQuestions()

	// Запуск теста
	handlers.StartQuiz(sections)
}

//TODO доработать логику структур вопросов
//TODO создать простой restful сервис
//TODO подумать какую БД выбрать
