package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Отображает домашнюю страницу"))
}

func showShippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображает определенную заметку"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Создает новую заметку"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showShippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
