package author

import (
	"bookApp/internal/domain/book"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Books []book.Book `gorm:"foreignKey:AuthorID;references:ID"`
}

// ToString: Convert author data into more readable string
func (a *Author) ToString() string {
	return fmt.Sprintf("ID: %s, Name: %s", a.ID, a.Name)
}

// ToStringWithBooks: Convert author data into more readable string with book more readable book info
func (a *Author) ToStringWithBooks() string {
	authorInfo := fmt.Sprintf("ID: %s, Name: %s", a.ID, a.Name)
	bookInfo := []string{}
	for _, book := range a.Books {
		bookInfo = append(bookInfo, book.ToString())
	}
	bookInfoJoined := strings.Join(bookInfo, "\n")

	return fmt.Sprintf("Author: %s has books:\n%s", authorInfo, bookInfoJoined)
}
