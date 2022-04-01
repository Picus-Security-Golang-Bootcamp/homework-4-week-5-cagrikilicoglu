package repos

import (
	"bookApp/internal/domain/entities"
	"encoding/csv"
	"os"
	"strconv"
	"sync"
)

// readDataWithWorkerPool: Reading a csv file concurrently and returns books and authors slices from the data in the file
func readDataWithWorkerPool(path string) ([]entities.Book, []entities.Author, error) {
	const numJobs = 5
	books := []entities.Book{}
	authors := []entities.Author{}
	jobs := make(chan []string, numJobs)
	resultsBooks := make(chan entities.Book, numJobs)
	resultsAuthors := make(chan entities.Author, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go toStruct(jobs, resultsBooks, resultsAuthors, &wg)
	}
	go func() {
		f, err := os.Open(path)
		if err != nil {
			return
		}
		defer f.Close()

		lines, err := csv.NewReader(f).ReadAll()
		if err != nil {
			return
		}
		for _, line := range lines[1:] {
			jobs <- line
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(resultsBooks)
		close(resultsAuthors)
	}()

	for b := range resultsBooks {
		books = append(books, b)
	}
	for a := range resultsAuthors {
		authors = append(authors, a)
	}

	return books, authors, nil
}

// ToStruct: creates book and author structs as the data from the file is read and send the structs to respective results channels
func toStruct(jobs <-chan []string, resultsBooks chan<- entities.Book, resultsAuthors chan<- entities.Author, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		// parse string data read from csv file to the matching types with book struct
		pageNumberParsed, err := strconv.Atoi(j[2])
		if err != nil {
			return
		}
		stockNumberParsed, err := strconv.Atoi(j[3])
		if err != nil {
			return
		}
		priceParsed, err := strconv.ParseFloat(j[5], 32)
		if err != nil {
			return
		}

		book := entities.Book{ID: j[0],
			Name:        j[1],
			PageNumber:  uint(pageNumberParsed),
			StockNumber: stockNumberParsed,
			StockID:     j[4],
			Price:       float32(priceParsed),
			ISBN:        j[6],
			AuthorID:    j[7],
			Author: &entities.Author{ID: j[7],
				Name: j[8]}}

		resultsAuthors <- *book.Author
		resultsBooks <- book
	}
}
