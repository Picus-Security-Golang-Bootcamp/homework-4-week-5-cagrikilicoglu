package main

import (
	postgres "bookApp/common/db"
	"bookApp/domain/author"
	"bookApp/domain/book"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// Set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")

	}

	// Initialize database
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot be initalized.")
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Database connection cannot be closed.")
	}
	defer sqlDb.Close()

	log.Println("Postgress connected")

	// Repositories
	bookRepo := book.NewBookRepository(db)
	bookRepo.SetupDatabase("./files/books.csv")

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.SetupDatabase("./files/authors.csv")

	// Initalize sample queries
	// SampleQueries(*bookRepo, *authorRepo)
}

// SampleQueries: to test the function of the queries defined in the app
func SampleQueries(bookRepo book.BookRepository, authorRepo author.AuthorRepository) {

	// Find all books in the book list (excluding soft-deleted)
	books, _ := bookRepo.FindAll()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Find all books in the book list including soft-deleted
	books, _ = bookRepo.FindAllIncludingDeleted()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Find all books currently in stock (excluding soft deleted)
	books, _ = bookRepo.FindAllInStock()
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Find all books under a certain price & currently in stock
	var samplePriceInput float32 = 16
	books, _ = bookRepo.FindAllBooksUnderPrice(float32(samplePriceInput))
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Find a book by ID
	var sampleIDInput string = "1"
	book, _ := bookRepo.FindByBookID(sampleIDInput)
	fmt.Println(book.ToString())

	// Find a book by ISBN
	var sampleISBNInput string = "9780385093798"
	book, _ = bookRepo.FindByBookISBN(sampleISBNInput)
	fmt.Println(book.ToString())

	// Find books by name (elastic search)
	var sampleBookNameInput string = "the"
	books, _ = bookRepo.FindByBookName(sampleBookNameInput)
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Find books by author's name (elastic search)
	var sampleAuthorNameInput string = "dickens"
	books, _ = bookRepo.FindByAuthorName(sampleAuthorNameInput)
	for _, book := range books {
		fmt.Println(book.ToString())
	}

	// Delete books by ID input
	var sampleDeleteInputID string = "3"
	bookRepo.DeleteByBookID(sampleDeleteInputID)

	// Buy books by ID input and quantity input
	var sampleBuyInputID string = "5"
	var sampleQuantity int = 10
	var sampleQuantityForNotEnoughStock int = 4
	bookRepo.BuyByBookID(sampleBuyInputID, sampleQuantity)
	// // To check "not enough stock error"
	err := bookRepo.BuyByBookID(sampleBuyInputID, sampleQuantityForNotEnoughStock)
	fmt.Println(err)

	// Find all authors with their books info
	authors, _ := authorRepo.FindAuthorsWithBookInfo()
	for _, author := range authors {
		fmt.Println(author.ToStringWithBooks())
	}

	// Find all authors without their books info
	authors, _ = authorRepo.FindAuthorsWithoutBookInfo()
	for _, author := range authors {
		fmt.Println(author.ToString())
	}

	// Find an author with ID input
	var sampleAuthorID string = "404"
	author, _ := authorRepo.FindByAuthorID(sampleAuthorID)
	fmt.Println(author.ToString())

	// Find authors by name (elastic search)
	sampleAuthorName := "j."
	authors, _ = authorRepo.FindByAuthorName(sampleAuthorName)
	for _, author := range authors {
		fmt.Println(author.ToString())
	}

	// Find books of an author by giving name input (of author's name) (elastic search)
	sampleAuthorName = "cao"
	authors, _ = authorRepo.FindBooksOfAuthorByName(sampleAuthorName)
	for _, author := range authors {
		fmt.Println(author.ToStringWithBooks())
	}

}
