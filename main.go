package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// PageData はテンプレートに渡すデータ構造体です
type PageData struct {
	Title   string
	Message string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := PageData{
		Title:   "Goウェブアプリ",
		Message: "Hello, World!",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Printf("テンプレート実行エラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()
	// 静的ファイルのサーブ（オプション）
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", homeHandler).Methods("GET")

	// サーバー起動
	log.Println("サーバーを起動します。http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
