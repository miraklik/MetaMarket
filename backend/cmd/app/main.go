package main

import (
	"net/http"
)

func HandlerHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func HandleRegistet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi new user, let s registration!"))
}

func main() {
	http.HandleFunc("/", HandlerHi)
	http.HandleFunc("/registr", HandleRegistet)

	http.ListenAndServe(":8080", nil)

}
