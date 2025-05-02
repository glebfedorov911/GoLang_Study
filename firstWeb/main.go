package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world!</h1>"))
}

func main() {
	port := "8000"
	fmt.Printf("Веб-сервер успешно запущен на порте %s", port)

	mux := http.NewServeMux()

	mux.HandleFunc("/hello-world", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
