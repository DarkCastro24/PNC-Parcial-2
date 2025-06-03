package controllers

import (
	dto "PNC-Parcial-2/entities/dto"
	"PNC-Parcial-2/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type BookController struct {
	Service *services.BookService
}

func (uc *BookController) Routes(router *mux.Router) {
	router.HandleFunc("/books", uc.CreateBook).Methods("POST")
	router.HandleFunc("/books", uc.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", uc.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", uc.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", uc.DeleteBook).Methods("DELETE")
}

func (uc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var input dto.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := uc.Service.CreateBook(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (uc *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (uc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := uc.Service.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (uc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var input dto.BookDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := uc.Service.UpdateUser(id, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario actualizado correctamente"})
}

func (uc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := uc.Service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario eliminado correctamente"})
}
