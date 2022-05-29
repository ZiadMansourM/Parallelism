package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
	Problem: Concurrency Patterns
	Note: Toy programs running on a toy machines get 'toy results'.
	> Not Closing the results channel >>> The only factor is # of workers no "Sync"
	> The way I recieve is hardcoded
*/

type feb_num struct {
    worker_id int
    num int
    result int
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func worker(id int, jobs <-chan int, results chan<- feb_num) {
	for n := range jobs {
		results <- feb_num{id, n, fib(n)}
	}
}

func main() {
	fmt.Println("Version", runtime.Version())
    fmt.Println("NumCPU", runtime.NumCPU())
	workersCount := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 50};
	for _, workerCount := range workersCount {
		jobs := make(chan int, 100)
		results := make(chan feb_num, 100)

		startTime := time.Now()

		for id := 1; id <= workerCount; id++ {
			go worker(id, jobs, results)
		}

		for i := 1; i <= 42; i++ {
			jobs <- i
		}
		close(jobs)

		for i :=0; i != 42; i++ {
			<-results
		}
		var numberOfWorkers string = "<" + strconv.Itoa(workerCount) + ">:"
		fmt.Println("Total time taken" + numberOfWorkers, time.Since(startTime).Milliseconds(), "msec")
		time.Sleep(time.Second)
	}
}