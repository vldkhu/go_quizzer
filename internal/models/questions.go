package models

type Question struct {
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}

var Questions = []Question{
	{Text: "Какой язык программирования используется для разработки Go?", Options: []string{"Python", "Go", "Java"}, Answer: "Go"},
	{Text: "Что такое переменная?", Options: []string{"Хранилище данных", "Функция", "Класс"}, Answer: "Хранилище данных"},
	{Text: "Что такое слайс", Options: []string{"Массив", "Мапа", "Динамический массив"}, Answer: "Динамический массив"},
	{Text: "Какой язык программирования используется для разработки Go?", Options: []string{"Python", "Go", "Java"}, Answer: "Go"},
	{Text: "Что такое переменная?", Options: []string{"Хранилище данных", "Функция", "Класс"}, Answer: "Хранилище данных"},
	{Text: "Какой из следующих языков является языком разметки?", Options: []string{"HTML", "Python", "Java"}, Answer: "HTML"},
	{Text: "Что такое ООП?", Options: []string{"Объектно-ориентированное программирование", "Операционная система", "Объектно-ориентированная платформа"}, Answer: "Объектно-ориентированное программирование"},
	{Text: "Какой из следующих языков является функциональным?", Options: []string{"C", "Haskell", "Java"}, Answer: "Haskell"},
	{Text: "Что такое массив?", Options: []string{"Структура данных", "Функция", "Класс"}, Answer: "Структура данных"},
	{Text: "Какой из следующих языков является языком программирования для веб-разработки?", Options: []string{"HTML", "CSS", "JavaScript"}, Answer: "JavaScript"},
	{Text: "Что такое SQL?", Options: []string{"Язык программирования", "Язык запросов к базам данных", "Язык разметки"}, Answer: "Язык запросов к базам данных"},
	{Text: "Какой из следующих языков является языком программирования общего назначения?", Options: []string{"HTML", "Python", "CSS"}, Answer: "Python"},
	{Text: "Что такое Git?", Options: []string{"Система контроля версий", "Язык программирования", "Редактор кода"}, Answer: "Система контроля версий"},
	{Text: "Какой из следующих языков является языком разметки стилей?", Options: []string{"HTML", "CSS", "JavaScript"}, Answer: "CSS"},
	{Text: "Что такое API?", Options: []string{"Интерфейс программирования приложений", "Программное обеспечение", "Язык программирования"}, Answer: "Интерфейс программирования приложений"},
	{Text: "Какой из следующих языков является языком программирования для разработки мобильных приложений?", Options: []string{"Swift", "HTML", "CSS"}, Answer: "Swift"},
	{Text: "Что такое компилятор?", Options: []string{"Программа для перевода кода", "Редактор кода", "Система контроля версий"}, Answer: "Программа для перевода кода"},
	{Text: "Какой из следующих языков является языком программирования для анализа данных?", Options: []string{"Python", "HTML", "CSS"}, Answer: "Python"},
	{Text: "Что такое фреймворк?", Options: []string{"Библиотека для разработки", "Язык программирования", "Редактор кода"}, Answer: "Библиотека для разработки"},
	{Text: "Какой из следующих языков является языком программирования для разработки игр?", Options: []string{"C#", "HTML", "CSS"}, Answer: "C#"},
	{Text: "Что такое JSON?", Options: []string{"Формат обмена данными", "Язык программирования", "Система контроля версий"}, Answer: "Формат обмена данными"},
	{Text: "Какой из следующих языков является языком программирования для разработки серверов?", Options: []string{"Java", "HTML", "CSS"}, Answer: "Java"},
	{Text: "Что такое Docker?", Options: []string{"Платформа для контейнеризации", "Язык программирования", "Редактор кода"}, Answer: "Платформа для контейнеризации"},
	{Text: "Какой из следующих языков является языком программирования для разработки веб-приложений?", Options: []string{"Ruby", "HTML", "CSS"}, Answer: "Ruby"},
	{Text: "Что такое Agile?", Options: []string{"Методология разработки", "Язык программирования", "Система контроля версий"}, Answer: "Методология разработки"},
	{Text: "Какой из следующих языков является языком программирования для разработки искусственного интеллекта?", Options: []string{"Python", "HTML", "CSS"}, Answer: "Python"},
	{Text: "Что такое машинное обучение?", Options: []string{"Метод анализа данных", "Язык программирования", "Система контроля версий"}, Answer: "Метод анализа данных"},
	{Text: "Какой из следующих языков используется для разработки приложений на Android?", Options: []string{"Java", "HTML", "CSS"}, Answer: "Java"},
	{Text: "Что такое переменная в программировании?", Options: []string{"Хранилище данных", "Функция", "Класс"}, Answer: "Хранилище данных"},
	{Text: "Какой из следующих языков является языком программирования для разработки веб-сайтов?", Options: []string{"HTML", "Python", "Java"}, Answer: "HTML"},
	{Text: "Что такое CSS?", Options: []string{"Язык стилей", "Язык программирования", "Система контроля версий"}, Answer: "Язык стилей"},
	{Text: "Какой из следующих языков является языком программирования для разработки игр?", Options: []string{"C++", "HTML", "CSS"}, Answer: "C++"},
	{Text: "Что такое SQL?", Options: []string{"Язык запросов к базам данных", "Язык программирования", "Язык разметки"}, Answer: "Язык запросов к базам данных"},
	{Text: "Какой из следующих языков используется для создания веб-приложений?", Options: []string{"JavaScript", "HTML", "CSS"}, Answer: "JavaScript"},
	{Text: "Что такое фреймворк?", Options: []string{"Библиотека для разработки", "Язык программирования", "Редактор кода"}, Answer: "Библиотека для разработки"},
	{Text: "Какой из следующих языков является языком программирования для разработки серверов?", Options: []string{"Node.js", "HTML", "CSS"}, Answer: "Node.js"},
	{Text: "Что такое REST API?", Options: []string{"Стандарт для веб-сервисов", "Язык программирования", "Система контроля версий"}, Answer: "Стандарт для веб-сервисов"},
	{Text: "Какой из следующих языков используется для анализа данных?", Options: []string{"R", "HTML", "CSS"}, Answer: "R"},
	{Text: "Что такое Big Data?", Options: []string{"Обработка больших объемов данных", "Язык программирования", "Система контроля версий"}, Answer: "Обработка больших объемов данных"},
	{Text: "Какой из следующих языков является языком программирования для разработки мобильных приложений?", Options: []string{"Swift", "HTML", "CSS"}, Answer: "Swift"},
	{Text: "Что такое DevOps?", Options: []string{"Методология разработки", "Язык программирования", "Система контроля версий"}, Answer: "Методология разработки"},
	{Text: "Какой из следующих языков используется для создания игр?", Options: []string{"C#", "HTML", "CSS"}, Answer: "C#"},
	{Text: "Что такое Agile?", Options: []string{"Методология разработки", "Язык программирования", "Система контроля версий"}, Answer: "Методология разработки"},
	{Text: "Какой из следующих языков используется для создания веб-сайтов?", Options: []string{"PHP", "HTML", "CSS"}, Answer: "PHP"},
	{Text: "Что такое блокчейн?", Options: []string{"Технология распределенного реестра", "Язык программирования", "Система контроля версий"}, Answer: "Технология распределенного реестра"},
	{Text: "Какой из следующих языков является языком программирования для разработки искусственного интеллекта?", Options: []string{"Python", "HTML", "CSS"}, Answer: "Python"},
	{Text: "Что такое виртуальная машина?", Options: []string{"Эмуляция аппаратного обеспечения", "Язык программирования", "Система контроля версий"}, Answer: "Эмуляция аппаратного обеспечения"},
	{Text: "Какой из следующих языков используется для создания веб-приложений?", Options: []string{"Ruby", "HTML", "CSS"}, Answer: "Ruby"},
}
