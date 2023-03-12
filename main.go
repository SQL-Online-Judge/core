package main

import (
	"net/http"
	gateway "github.com/SQL-Online-Judge/core/service/gateway"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, SQL-Online-Judge!"))
	})
	http.ListenAndServe(":3000", r)
}
