package main

import (
	"encoding/json"
	"net/http"
)

// Структура для ответа
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Роут для обработки GET-запроса на /hello
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Создание ответа
		response := Response{Message: "Hello, World!"}

		// Установка заголовков
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Кодирование ответа в JSON и отправка его
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Goodbye!"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Запуск сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}
