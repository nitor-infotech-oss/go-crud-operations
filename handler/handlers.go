package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-api-patient/db"
	"rest-api-patient/model"

	"github.com/gorilla/mux"
)

//gorilla mux used for simplifying api code , response will be in json format

//Returns list of patients
func GetPatients(w http.ResponseWriter, _ *http.Request) {
	patients := db.FindAll()

	bytes, err := json.Marshal(patients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

//Returns single entry of Patient
func GetPatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	patient, ok := db.FindBy(name)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := json.Marshal(patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

//Save Patient's record
func SavePatient(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	patient := new(model.Patient)
	err = json.Unmarshal(body, patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(patient.Name, patient)

	w.Header().Set("Location", r.URL.Path+"/"+patient.Name)
	w.WriteHeader(http.StatusCreated)
}

//Updates Patient's record 
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	patient := new(model.Patient)
	err = json.Unmarshal(body, patient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(name, patient)
}

//Deletes patient's record
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	db.Remove(name)
	w.WriteHeader(http.StatusNoContent)
}

//Json response formatter
func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
