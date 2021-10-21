package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Get("/timeout/{seconds}", timeoutHandler)
	router.Get("/health", livenessProbe)
	log.Println(http.ListenAndServe(":8080", router))
}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	startedAt := time.Now()
	seconds, err := strconv.Atoi(chi.URLParam(r, "seconds"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, _ = w.Write([]byte("Timeout is: "))
	<-time.After(time.Duration(seconds) * time.Second)
	_, _ = w.Write([]byte(time.Since(startedAt).String()))
}

func livenessProbe(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("ok"))
}
