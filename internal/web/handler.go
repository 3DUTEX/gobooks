package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
)

type BookHandlers struct {
	service *service.BookService
}

func (handler *BookHandlers) GetBooks(res http.ResponseWriter, req http.Request) {
	books, err := handler.service.GetBooks()

	if (err != nil) {
		http.Error(res, "faield to get books!", 500)
		return
	}

	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(books)
}