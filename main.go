package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
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

func tacticsHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートファイルのパスを明示的に指定
	tmplPath := filepath.Join("templates", "tactics", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("テンプレート読み込みエラー: %v", err)
		http.Error(w, "Template Not Found", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:   "隅の形",
		Message: "Hello, World!",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("テンプレート実行エラー: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()

	// 静的ファイルの提供を有効化
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// ルートの設定
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/tactics", tacticsHandler).Methods("GET")

	// ルーターを使用してサーバーを起動
	http.Handle("/", r)

	// サーバー起動
	log.Println("サーバーを起動します。http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
