package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ibanezv/fly-devs-practice_go/internal/services"
	"gorm.io/gorm"
)

type Book struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *Book {
	return &Book{service: service}
}

func (s *Book) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var id = params["id"]
	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "id not valid", http.StatusBadRequest)
		return
	}

	book, err := s.service.GetById(r.Context(), bookId)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
