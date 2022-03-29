package author

import (
	"encoding/csv"
	"os"
	"sync"
)

// readAuthorsWithWorkerPool: Reading a csv file concurrently and returns an author slice with the authors in the file
func readAuthorsWithWorkerPool(path string) ([]Author, error) {
	const numJobs = 5
	authors := []Author{}
	jobs := make(chan []string, numJobs)
	results := make(chan Author, numJobs)
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

	for a := range results {
		authors = append(authors, a)
	}

	return authors, nil
}

// toStruct: creates an author struct as the data from the file is read and send the struct to results channel
func toStruct(jobs <-chan []string, results chan<- Author, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {

		author := Author{ID: j[0],
			Name: j[1]}

		results <- author
	}
}
