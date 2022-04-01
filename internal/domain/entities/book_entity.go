package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string  `json:"ID" gorm:"unique"`
	Name        string  `json:"name"`
	PageNumber  uint    `json:"pageNumber"`
	StockNumber int     `json:"stockNumber"`
	StockID     string  `json:"stockId" gorm:"unique"`
	Price       float32 `json:"price"`
	ISBN        string  `json:"isbn"`
	AuthorID    string  `json:"authorID"`
	Author      *Author `json:",omitempty" gorm:"OnDelete:SET NULL"`
}

// ToString: Convert book data into more readable string
func (b *Book) ToString() string {

	return fmt.Sprintf("ID: %s, Name: %s, Page Number: %d, Stock Number: %d, StockID: %s, Price: %.2f, ISBN: %s, Author ID: %s\n", b.ID, b.Name, b.PageNumber, b.StockNumber, b.StockID, b.Price, b.ISBN, b.AuthorID)
}

// BeforeDelete: Print book name before deleting.
func (b *Book) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s is deleting...\n", b.Name)
	fmt.Println(b.ToString())
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
