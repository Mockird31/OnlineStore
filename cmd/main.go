package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte{'h','e','l','l','o'})
	})
	srv := &http.Server{
		Addr: ":8081",
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}