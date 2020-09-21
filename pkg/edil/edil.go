package edil

import (
	"context"
	"fmt"
	"net/http"

	"github.com/padwalab/mcsvcs/internal"
)

type edilService struct{}

func NewService() Service { return &edilService{} }

func (w *edilService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	doc := internal.Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rowling",
		Topic:   "Fiction and Magic",
	}
	return []internal.Document{doc}, nil
}

func (w *edilService) ServiceStatus(_ context.Context) (int, error) {
	fmt.Println("Checking the Service health...")
	return http.StatusOK, nil
}
