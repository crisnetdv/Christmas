package main

import (
	"os"
	"log"
	"net/http"
	"github.com/Cristiano67867/christmas/handlers"
)

const filename string = "./comments.json"

func init() {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, _ := os.Create(filename)
		defer file.Close()
		file.WriteString("[]")
	}
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

	http.HandleFunc("/comment", handlers.HandlerCommentRouter)
	http.HandleFunc("/comments", handlers.HandlerGetAllCommentRouter)

	log.Println("The server is running in " + port)
	http.ListenAndServe(":" + port, nil)
}