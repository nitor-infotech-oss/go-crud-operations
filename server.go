package main

import (
	"log"
	"net/http"
	"rest-api-patient/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/patient").HandlerFunc(handler.GetPatients)
	sub.Methods("POST").Path("/patient").HandlerFunc(handler.SavePatient)
	sub.Methods("GET").Path("/patient/{name}").HandlerFunc(handler.GetPatient)
	sub.Methods("PUT").Path("/patient/{name}").HandlerFunc(handler.UpdatePatient)
	sub.Methods("DELETE").Path("/patient/{name}").HandlerFunc(handler.DeletePatient)

	log.Fatal(http.ListenAndServe(":8089", router))//local port for api end point like http://localhost:8089/api/v1/patient
}
