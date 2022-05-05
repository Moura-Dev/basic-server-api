package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	fmt.Println("Server is Online")
	http.ListenAndServe(":5050", r)
	println("Hello, world.")
}
