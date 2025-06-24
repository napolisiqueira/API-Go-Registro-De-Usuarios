package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	Name  string `json:"nome"`
	ID    int    `json:"id"`
	Email string `json:"email"`
	Done  bool   `json:"done"`
}

var persons []Person
var currentID = 0

func main() {

	router := mux.NewRouter()
	// Feito
	router.HandleFunc("/listar", get_persons).Methods("GET")
	// Feito
	router.HandleFunc("/listar/{id}", get_person_by_id).Methods("GET")
	// TODO
	router.HandleFunc("/cadastrar", create_persons).Methods("POST")
	// TODO
	router.HandleFunc("/deletar", delete_persons).Methods("DELETE")
	// TODO
	router.HandleFunc("/atualizar", update_persons).Methods("PATCH")

	fmt.Println("Iniciando servidor http//:3000")
	http.ListenAndServe(":3000", router)

}

func get_persons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func get_person_by_id(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for _, person := range persons {
		if person.ID == id {
			json.NewEncoder(w).Encode(persons)
		}
	}
	http.Error(w, "Not finalized", http.StatusNotFound)
}

func create_persons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro nos dados da requisição", http.StatusPartialContent)
		return
	}
	err = json.Unmarshal(body, &person)
	if err != nil {
		http.Error(w, "Erro ao decodificar corpo da requisição", http.StatusBadRequest)
		return
	}

	person.ID = currentID
	currentID++
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

func delete_persons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	indexRemoved := -1
	var personRemoved Person
	for i, p := range persons {
		if p.ID == id {
			indexRemoved = i
			personRemoved = p
			break
		}
	}
	if indexRemoved == -1 {
		http.Error(w, "ID invalido.", http.StatusBadRequest)
		return
	}

	persons = append(persons[:indexRemoved], persons[indexRemoved+1:]...)
	json.NewEncoder(w).Encode(personRemoved)
}

func update_persons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	var updateFields Person
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	foudIndex := -1
	for i, p := range persons {
		if p.ID == id {
			foudIndex = i
			break
		}
	}
	if foudIndex == -1 {
		http.Error(w, "ID invalido.", http.StatusBadRequest)
		return
	}

	if updateFields.Name != "" {
		persons[foudIndex].Name = updateFields.Name
	}

	if updateFields.Done != true {
		updateFields.Done = true
		persons[foudIndex].Done = updateFields.Done
	}

	if updateFields.Email != "" {
		persons[foudIndex].Email = updateFields.Email
	}

	json.NewEncoder(w).Encode(updateFields)
}
