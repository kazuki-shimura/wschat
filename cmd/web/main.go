package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("ポート8080にてサーバーを起動します")

	_ = http.ListenAndServe(":8080", mux)
}
