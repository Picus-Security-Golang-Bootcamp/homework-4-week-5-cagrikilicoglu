package router

import (
	"bookApp/internal/api/router/httpErrors"
	"bookApp/internal/domain/entities"
	"bookApp/internal/domain/repos"
	"encoding/json"
	"net/http"
	"strconv"

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

type ApiResponse struct {
	Payload interface{} `json:"data"`
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data := ApiResponse{
		Payload: payload,
	}
	response, err := json.Marshal(data)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, a httpErrors.ApiErr) {
	respondWithJson(w, a.Status(), a.Error())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	welcomeMessage := "Welcome to the book store"
	respondWithJson(w, http.StatusOK, welcomeMessage)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := BookRepo.FindAll()
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func GetBooksInludingDeleted(w http.ResponseWriter, r *http.Request) {
	books, err := BookRepo.FindAllIncludingDeleted()
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func GetBooksInStock(w http.ResponseWriter, r *http.Request) {
	books, err := BookRepo.FindAllInStock()
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func GetBooksUnderPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	price, _ := strconv.ParseFloat(vars["priceunder"], 32)
	books, err := BookRepo.FindAllBooksUnderPrice(float32(price))
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func GetBookByBookID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book, err := BookRepo.FindByBookID(id)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	// dereference operatörüne bak
	respondWithJson(w, http.StatusOK, *book)
}

func GetBookByISBN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	book, err := BookRepo.FindByBookISBN(isbn)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	// dereference operatörüne bak
	respondWithJson(w, http.StatusOK, *book)
}
func GetBookByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	books, err := BookRepo.FindByBookName(name)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := BookRepo.DeleteByBookID(id)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func BuyBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	quantiy, err := strconv.Atoi(vars["quantity"])
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	err = BookRepo.BuyByBookID(id, quantiy)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func AddBookToDatabase(w http.ResponseWriter, r *http.Request) {
	var newBook entities.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	err = BookRepo.AddBook(newBook)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, newBook)
}

func GetAuthorsWithBookInfo(w http.ResponseWriter, r *http.Request) {

	authors, err := AuthorRepo.FindAuthorsWithBookInfo()
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}

	respondWithJson(w, http.StatusOK, authors)
}

func GetAuthorsWithoutBookInfo(w http.ResponseWriter, r *http.Request) {

	authors, err := AuthorRepo.FindAuthorsWithoutBookInfo()
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}

	respondWithJson(w, http.StatusOK, authors)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	author, err := AuthorRepo.FindByAuthorID(id)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}

	respondWithJson(w, http.StatusOK, author)
}
func GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	authors, err := AuthorRepo.FindByAuthorName(name)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, authors)
}

func GetBooksOfAuthorByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	authors, err := AuthorRepo.FindBooksOfAuthorByName(name)
	if err != nil {
		respondWithError(w, httpErrors.ParseErrors(err))
		return
	}
	respondWithJson(w, http.StatusOK, authors)
}
