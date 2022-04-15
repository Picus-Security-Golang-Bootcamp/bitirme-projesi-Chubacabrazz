package category

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var CategoryList []Category

func ReadCategoryWithWorkerPool(path string) error {
	const numJobs = 5
	jobs := make(chan []string, numJobs)
	results := make(chan Category, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}

	go func() {
		fmt.Println("open file running...")
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()

		for _, line := range lines[1:] {
			jobs <- line
		}

		close(jobs)
	}()

	go func() {
		fmt.Println("wait")
		wg.Wait()
		close(results)
	}()

	for v := range results {
		CategoryList = append(CategoryList, v)
	}
	return nil
}

//func toStruct reads given csv file by line and pushes them to result channel.
func toStruct(jobs <-chan []string, results chan<- Category, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		category := Category{}
		category.ID = j[0]
		category.Name = j[1]
		category.Desc = j[2]
		category.IsActive, _ = strconv.ParseBool(j[3])

		results <- category
	}
}
