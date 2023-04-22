package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PageData struct {
	Title   string
	Snippet []string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)

		return
	}

	//
	db, err := sql.Open("mysql", "root:@/snippetbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT title FROM snippets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var snippets []string
	for rows.Next() {
		var title string
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		snippets = append(snippets, title)
	}
	//
	data := PageData{
		Title:   "Домашняя страница",
		Snippet: snippets,
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type PageAbout struct {
	Title       string
	SnipContent []string
}

func about(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/about.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	db, err := sql.Open("mysql", "root@/snippetbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT content FROM snippets")
	if err != nil {
		log.Fatal(err)
	}

	var snipContents []string
	for rows.Next() {
		var content string
		err := rows.Scan(&content)
		if err != nil {
			log.Fatal(err)
		}
		snipContents = append(snipContents, content)
	}

	data := PageAbout{
		"Данные с базы О нас",
		snipContents,
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type PageAllContacts struct {
	Id    int
	Title string
}

func contacts(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/contacts.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	db, err := sql.Open("mysql", "root@/snippetbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title FROM snippets")
	if err != nil {
		log.Fatal(err)
	}

	var data []PageAllContacts
	for rows.Next() {
		var id int
		var title string

		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		item := PageAllContacts{id, title}
		data = append(data, item)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение определенной заметки с ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", 405)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}
