package author

import (
	"fmt"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

// SetupDatabase: automatically migrates database of Authors with gorm and insert author data to database by the given input path
func (a *AuthorRepository) SetupDatabase(path string) {
	a.Migrations()
	a.InsertAuthorData(path)
}

// Migrations: automatically migrates database of Authors
func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

// InsertAuthorData: insert author data to database by the given input path
func (a *AuthorRepository) InsertAuthorData(path string) error {

	authors, err := readAuthorsWithWorkerPool(path)
	if err != nil {
		return err
	}
	for _, author := range authors {
		a.db.Where(Author{ID: author.ID}).Attrs(Author{ID: author.ID, Name: author.Name}).FirstOrCreate(&author)
	}
	return nil
}

// FindAuthorsWithBookInfo: Find all the authors with their book data
func (a *AuthorRepository) FindAuthorsWithBookInfo() ([]Author, error) {
	authors := []Author{}
	result := a.db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

// FindAuthorsWithBookInfo: Find all the authors without their book data
func (a *AuthorRepository) FindAuthorsWithoutBookInfo() ([]Author, error) {
	authors := []Author{}
	result := a.db.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

// FindByAuthorID: returns the author with given ID input
// the search is elastic and case insensitive
func (a *AuthorRepository) FindByAuthorID(ID string) (*Author, error) {
	author := Author{}
	result := a.db.Where(&Author{ID: ID}).First(&author)
	if result.Error != nil {
		return nil, result.Error
	}
	return &author, nil
}

// FindByAuthorName: returns the author with given name input
// the search is elastic and case insensitive
func (a *AuthorRepository) FindByAuthorName(name string) ([]Author, error) {
	authors := []Author{}
	nameString := fmt.Sprintf("%%%s%%", name)

	result := a.db.Where("name ILIKE ?", nameString).Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

// FindBooksOfAuthorByName: returns the author with given name input as well as his/her books
// the search is elastic and case insensitive
func (a *AuthorRepository) FindBooksOfAuthorByName(name string) ([]Author, error) {
	authors := []Author{}
	nameString := fmt.Sprintf("%%%s%%", name)

	result := a.db.Preload("Books").Where("name ILIKE ?", nameString).Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}
