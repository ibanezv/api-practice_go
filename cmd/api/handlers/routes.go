package handlers

import (
	"github.com/gorilla/mux"
	"github.com/ibanezv/fly-devs-practice_go/internal/services"
)

func ApiRoutesMapper(bookService services.BookService) *mux.Router {
	h := NewBookHandler(bookService)
	router := mux.NewRouter()
	router.HandleFunc("/book/{id}", h.Get).Methods("GET")
	return router
}
