package handlers

import (
	"go_quizzer/internal/models"
	"reflect"
	"testing"
)

func TestGetRandomQuestions(t *testing.T) {
	type args struct {
		questions []models.Question
		n         int
	}
	tests := []struct {
		name string
		args args
		want []models.Question
	}{
		{
			name: "Обычный случай",
			args: args{
				questions: []models.Question{
					{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
					{Text: "Вопрос 2", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 2"},
					{Text: "Вопрос 3", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
				},
				n: 2,
			},
			want: []models.Question{
				{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
				{Text: "Вопрос 2", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 2"},
			},
		},
		{
			name: "Точное совпадение",
			args: args{
				questions: []models.Question{
					{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
					{Text: "Вопрос 2", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 2"},
					{Text: "Вопрос 3", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
				},
				n: 3,
			},
			want: []models.Question{
				{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
				{Text: "Вопрос 2", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 2"},
				{Text: "Вопрос 3", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
			},
		},
		{
			name: "Больше доступного",
			args: args{
				questions: []models.Question{
					{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
					{Text: "Вопрос 2", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 2"},
				},
				n: 5,
			},
			want: []models.Question{
				{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
				{Text: "Вопрос 2", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 2"},
			}, // Ожидаемый результат - все доступные вопросы
		},
		{
			name: "Пустой срез",
			args: args{
				questions: []models.Question{},
				n:         0,
			},
			want: []models.Question{}, // Ожидаемый результат - пустой срез
		},
		{
			name: "Один вопрос",
			args: args{
				questions: []models.Question{
					{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
				},
				n: 1,
			},
			want: []models.Question{
				{Text: "Вопрос 1", Options: []string{"Ответ 1", "Ответ 2"}, Answer: "Ответ 1"},
			}, // Ожидаемый результат - единственный вопрос
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandomQuestions(tt.args.questions, tt.args.n)
			// Поскольку результат рандомизирован, проверим только длину и наличие элементов
			if len(got) != len(tt.want) {
				t.Errorf("GetRandomQuestions() = %v, want %v", got, tt.want)
			}
			// Проверяем, что все элементы из want присутствуют в got
			for _, w := range tt.want {
				found := false
				for _, g := range got {
					if reflect.DeepEqual(g, w) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("GetRandomQuestions() missing element: %v", w)
				}
			}
		})
	}
}
