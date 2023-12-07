package main

import (
	httphandler "go_spiral/spiral_client/http_handler"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var Server httphandler.HandlerHub
	r := mux.NewRouter()
	r.HandleFunc("/", Server.HomeHandler)
	r.HandleFunc("/login/{username}", Server.AuthHandler)
	r.HandleFunc("/border", Server.BorderHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
