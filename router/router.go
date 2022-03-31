package router

import (
	"bookApp/domain/author"
	"bookApp/domain/book"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// var r mux.Router{}
// type MuxRouter struct {
// 	MuxRouter *mux.Router
// }

var (
	BookRepo   *book.BookRepository
	AuthorRepo *author.AuthorRepository
)

func Handle(mr *mux.Router) {
	mr.HandleFunc("/", HomeHandler)

	b := mr.PathPrefix("/books").Subrouter()
	b.HandleFunc("/", GetBooks).Methods("GET")
	b.HandleFunc("/all", GetBooksInludingDeleted).Methods("GET")
	b.HandleFunc("/stock", GetBooksInludingDeleted).Methods("GET")
	b.HandleFunc("/price/{priceunder}", GetBooksUnderPrice).Methods("GET")
	b.HandleFunc("", GetBookByBookID).Methods("GET").Queries("id", "{id}")
	// burayı farklı şekilde handle edbeiliriz
	b.HandleFunc("", GetBookByISBN).Methods("GET").Queries("isbn", "{isbn}")
	b.HandleFunc("", GetBookByName).Methods("GET").Queries("name", "{name}")
	b.HandleFunc("", GetBooksByAuthorName).Methods("GET").Queries("author", "{author}")
	b.HandleFunc("", DeleteBookById).Methods("DELETE").Queries("id", "{id}")
	b.HandleFunc("/order", DeleteBookById).Methods("PATCH").Queries("id", "{id}", "quantity", "{quantity}")

	a := mr.PathPrefix("/authors").Subrouter()
	a.HandleFunc("/", GetAuthorsWithBookInfo).Methods("GET")
	a.HandleFunc("/*", GetAuthorsWithoutBookInfo).Methods("GET")
	a.HandleFunc("", GetAuthorByID).Methods("GET").Queries("id", "{id}")
	a.HandleFunc("", GetAuthorByName).Methods("GET").Queries("name", "{name}")

	// aşağıdaki fonksiyon çalışmıyor
	// a.HandleFunc("", GetBooksOfAuthorByName).Methods("GET").Queries("name", "{name}")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the book store\n"))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// r.Header.Get("Content-Type") != "application/json"

	books, err := BookRepo.FindAll()
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
	// bytes, err := json.Marshal(books)
	// if err != nil {
	// 	w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	// }
	respondWithJson(w, http.StatusOK, books)
}

func GetBooksInludingDeleted(w http.ResponseWriter, r *http.Request) {

	books, err := BookRepo.FindAllIncludingDeleted()
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, books)
}

func GetBooksInStock(w http.ResponseWriter, r *http.Request) {

	books, err := BookRepo.FindAllInStock()
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, books)
}

func GetBooksUnderPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	price, _ := strconv.ParseFloat(vars["priceunder"], 32)
	books, err := BookRepo.FindAllBooksUnderPrice(float32(price))
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, books)
}

func GetBookByBookID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// vars := r.URL.Query()
	id := vars["id"]

	book, err := BookRepo.FindByBookID(id)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	// dereference operatörüne bak
	respondWithJson(w, http.StatusOK, *book)
}

func GetBookByISBN(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	book, err := BookRepo.FindByBookISBN(isbn)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	// dereference operatörüne bak
	respondWithJson(w, http.StatusOK, *book)
}
func GetBookByName(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	vars := mux.Vars(r)
	name := vars["name"]
	books, err := BookRepo.FindByBookName(name)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	// dereference operatörüne bak
	respondWithJson(w, http.StatusOK, books)
}

func GetBooksByAuthorName(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	vars := mux.Vars(r)
	author := vars["author"]
	books, err := BookRepo.FindByAuthorName(author)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	// dereference operatörüne bak
	respondWithJson(w, http.StatusOK, books)
}
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	vars := mux.Vars(r)
	id := vars["id"]
	err := BookRepo.DeleteByBookID(id)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func BuyBookById(w http.ResponseWriter, r *http.Request) {
	// vars := r.URL.Query()
	vars := mux.Vars(r)
	id := vars["id"]
	quantiy, err := strconv.Atoi(vars["quantity"])
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
	err = BookRepo.BuyByBookID(id, quantiy)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func GetAuthorsWithBookInfo(w http.ResponseWriter, r *http.Request) {

	authors, err := AuthorRepo.FindAuthorsWithBookInfo()
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, authors)
}

func GetAuthorsWithoutBookInfo(w http.ResponseWriter, r *http.Request) {

	authors, err := AuthorRepo.FindAuthorsWithoutBookInfo()
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, authors)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	author, err := AuthorRepo.FindByAuthorID(id)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, author)
}
func GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	authors, err := AuthorRepo.FindByAuthorName(name)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, authors)
}

func GetBooksOfAuthorByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	authors, err := AuthorRepo.FindBooksOfAuthorByName(name)
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}

	respondWithJson(w, http.StatusOK, authors)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		// respond with error'a çevir
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}