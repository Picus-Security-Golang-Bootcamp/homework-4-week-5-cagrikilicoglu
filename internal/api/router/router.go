package router

import (
	"bookApp/internal/domain/repos"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	BookRepo   *repos.BookRepository
	AuthorRepo *repos.AuthorRepository
)

func Handle(mr *mux.Router) {

	// home handler
	mr.HandleFunc("/", HomeHandler)

	// handlers regarding books
	b := mr.PathPrefix("/books").Subrouter()
	b.HandleFunc("/", GetBooks).Methods(http.MethodGet)
	b.HandleFunc("/all", GetBooksInludingDeleted).Methods(http.MethodGet)
	b.HandleFunc("/stock", GetBooksInStock).Methods(http.MethodGet)
	b.HandleFunc("/price/{priceunder}", GetBooksUnderPrice).Methods(http.MethodGet)
	b.HandleFunc("", GetBookByBookID).Methods(http.MethodGet).Queries("id", "{id}")
	b.HandleFunc("", GetBookByISBN).Methods(http.MethodGet).Queries("isbn", "{isbn}")
	b.HandleFunc("", GetBookByName).Methods(http.MethodGet).Queries("name", "{name}")
	b.HandleFunc("/delete", DeleteBookById).Methods(http.MethodDelete).Queries("id", "{id}")
	b.HandleFunc("/order", BuyBookById).Methods(http.MethodPatch).Queries("id", "{id}", "quantity", "{quantity}")
	b.HandleFunc("/add", AddBookToDatabase).Methods(http.MethodPost)

	// handlers regarding authors
	a := mr.PathPrefix("/authors").Subrouter()
	a.HandleFunc("/", GetAuthorsWithBookInfo).Methods(http.MethodGet)
	a.HandleFunc("/*", GetAuthorsWithoutBookInfo).Methods(http.MethodGet)
	a.HandleFunc("", GetAuthorByID).Methods(http.MethodGet).Queries("id", "{id}")
	a.HandleFunc("", GetAuthorByName).Methods(http.MethodGet).Queries("name", "{name}")
	a.HandleFunc("/books", GetBooksOfAuthorByName).Methods(http.MethodGet).Queries("name", "{name}")
}
