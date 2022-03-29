package book

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
)

// readBooksWithWorkerPool: Reading a csv file concurrently and returns a book slice with the books in the file
func readBooksWithWorkerPool(path string) ([]Book, error) {
	const numJobs = 5
	books := []Book{}
	jobs := make(chan []string, numJobs)
	results := make(chan Book, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg)
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
		close(results)
	}()

	for b := range results {
		books = append(books, b)
	}

	return books, nil
}

// toStruct: creates a book struct as the data from the file is read and send the struct to results channel
func toStruct(jobs <-chan []string, results chan<- Book, wg *sync.WaitGroup) {
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

		book := Book{ID: j[0],
			Name:        j[1],
			PageNumber:  uint(pageNumberParsed),
			StockNumber: stockNumberParsed,
			StockID:     j[4],
			Price:       float32(priceParsed),
			ISBN:        j[6],
			AuthorID:    j[7],
			AuthorName:  j[8]}

		results <- book
	}
}
