package main

import (
	"html/template"
	"log"
	"net/http"
)

// PageData はテンプレートに渡すデータ構造体です
type PageData struct {
	Title   string
	Message string
}

func main() {
	// テンプレートファイルのパス設定
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// ハンドラーの設定
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:   "Goウェブアプリ",
			Message: "Hello, World!",
		}

		// テンプレートの実行
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("テンプレート実行エラー: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// 静的ファイルのサーブ（オプション）
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// サーバー起動
	log.Println("サーバーを起動します。http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
