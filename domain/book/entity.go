package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string
	Name        string
	PageNumber  uint
	StockNumber int
	StockID     string
	Price       float32
	ISBN        string
	AuthorID    string
	AuthorName  string
}

// ToString: Convert book data into more readable string
func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %s, Name: %s, Page Number: %d, Stock Number: %d, StockID: %s, Price: %.2f, ISBN: %s, Author ID: %s, Author Name: %s\n", b.ID, b.Name, b.PageNumber, b.StockNumber, b.StockID, b.Price, b.ISBN, b.AuthorID, b.AuthorName)
}

// BeforeDelete: Print book name before deleting.
func (b *Book) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s is deleting...\n", b.Name)
	fmt.Printf(b.ToString())
	return nil
}

// AfterDelete: Print book name after it is deleted with a success message.
func (b *Book) AfterDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s is succesfully deleted...\n", b.Name)
	return nil
}

// AfterDelete: Print book name after it is deleted with a success message.
func (b *Book) AfterOrder(num int) {
	fmt.Printf("Book %s of quantity %d is succesfully ordered...\n", b.Name, num)
}
