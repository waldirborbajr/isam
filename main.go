package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Heroku")
}
