package main

import (
	"os"
	"log"
	"net/http"
)

func HandlerHomeRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	templates := http.FileServer(http.Dir("./templates"))
	static := http.FileServer(http.Dir("./static"))

	http.Handle("/", templates)
	http.Handle("/static/", http.StripPrefix("/static", static))

	log.Println("The server is running in " + port)
	http.ListenAndServe(":" + port, nil)
}