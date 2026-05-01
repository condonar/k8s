package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	const PORT = "3000"

	// Аналог app.get("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Получаем имя хоста (в K8s это будет имя Пода)
		hostname, _ := os.Hostname()
		
		helloMessage := fmt.Sprintf("<h1>VERSION 2: Hello from the %s</h1>", hostname)
		
		// Печать в консоль 
		fmt.Println(helloMessage)
		
		// Отправка ответа
		fmt.Fprint(w, helloMessage)
	})

	fmt.Printf("Web server is listening at port %s\n", PORT)
	
	// Запуск сервера
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		fmt.Printf("Ошибка запуска: %s\n", err)
	}
}