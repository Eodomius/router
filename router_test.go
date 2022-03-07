package router

import (
	"net/http"
	"testing"
)

const PORT = "8080"

func TestRouter(t *testing.T){

	var router = New()
	http.Handle("/", router)
	router.Get("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Get : Test"))
	})
	router.Post("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Post : Test"))
	})
	router.Patch("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Patch : Test"))
	})
	router.Put("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Put : Test"))
	})
	router.Delete("/test", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Delete : Test"))
	})
	http.ListenAndServe(":"+PORT, nil)
}