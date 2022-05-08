package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"db-test/app"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", app.CreateProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", app.GetAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", app.GetUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", app.UpdateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", app.DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
