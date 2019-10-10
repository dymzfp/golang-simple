package main

import (
	"log"
	"net/http"

	// ndb "github.com/haidlir/x-golang-course/021-simple-rest-api/db"
	ndb "github.com/dymzfp/golang-simple/db"
	// nhandler "github.com/haidlir/x-golang-course/021-simple-rest-api/handler"
	nhandler "github.com/dymzfp/golang-simple/handler"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8081"
)

// Handler hold the function handler for API's endpoint.
type Handler interface {
	AddSiswa(w http.ResponseWriter, r *http.Request)
	GetAllSiswa(w http.ResponseWriter, r *http.Request)
}

// Router
func NewRouter(handler Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/siswa", handler.GetAllSiswa).Methods(http.MethodGet)
	r.HandleFunc("/api/siswa", handler.AddSiswa).Methods(http.MethodPost)
	return r
}

func main() {
	log.Println("Start service...")

	db := ndb.NewDummyDB()
	log.Println("Successfully Conneceted to DB")

	handler := nhandler.NewHandler(db)
	r := NewRouter(handler)

	log.Printf("Starting http server at %v", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatalf("Unable to start server, %v", err)
	}
	log.Println("Stopping API Service...")
}

