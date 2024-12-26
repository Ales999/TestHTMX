package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)

}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	// Обработчик для динамической загрузки контента
	w.Write([]byte("<p>Это Загруженный контент</p>"))
}

func main() {
	// Локальные JS надо загружать отдельно от темплейтов.
	// Описано тут хорошо: https://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//  Загрузка шаблонов.
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/load", loadHandler)

	err := http.ListenAndServe("localhost:8083", nil)
	if err != nil {
		log.Fatal(err)
	}
}
