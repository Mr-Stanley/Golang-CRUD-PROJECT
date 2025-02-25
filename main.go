package main

import (
	"GO-CRUD-PROJECT/internal/handler"
	_ "GO-CRUD-PROJECT/internal/handler"
	"GO-CRUD-PROJECT/internal/repository"
	_ "GO-CRUD-PROJECT/internal/repository"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

func main() {
	repository.InitMongoDB("mongodb://localhost:27017", "crud_db", "users")
	defer repository.DisconnectMongoDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
