package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	const PORT = "3000"

	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		// Указываем браузеру, что это HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		helloMessage := fmt.Sprintf("<h1>Hello from the %s</h1>", hostname)

		fmt.Println("Запрос на / с пода:", hostname)
		fmt.Fprint(w, helloMessage)
	})

	// Запрос к стороннему сервису nginx внутри K8s
	http.HandleFunc("/nginx", func(w http.ResponseWriter, r *http.Request) {
		// Делаем запрос к сервису nginx
		resp, err := http.Get("http://nginx")
		if err != nil {
			http.Error(w, "Ошибка запроса к nginx: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Читаем тело ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения ответа", http.StatusInternalServerError)
			return
		}

		// Пробрасываем ответ от nginx пользователю
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, string(body))
	})

	// Новый эндпоинт /posts
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		// 1. Делаем запрос к внешнему API (JSONPlaceholder)
		// Это аналог: await fetch('https://jsonplaceholder.typicode.com')
		resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			http.Error(w, "Ошибка запроса к API: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// 2. Читаем тело ответа
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
			return
		}

		// 3. Указываем, что отправляем JSON (чтобы браузер подсветил синтаксис)
		w.Header().Set("Content-Type", "application/json")

		// Отправляем JSON пользователю
		fmt.Fprint(w, string(body))
	})

	fmt.Printf("Web server is listening at port %s\n", PORT)

	// Запуск сервера
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		fmt.Printf("Ошибка запуска: %s\n", err)
	}
}
