package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
	Problem: Concurrency Patterns
	Notes:
		-> There are so many interesting Concurrency Patterns
		out there
		-> Toy programs running on a toy machines get 'toy results'.
			> Not Closing the results channel correctly >>> "sync" pkg
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
	close(results)
}

func main() {
	fmt.Println("Version", runtime.Version())
    fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(1)
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))


	const workersCount = 8;

	jobs := make(chan int, 100)
	results := make(chan feb_num, 100)

	startTime := time.Now()

	for id := 1; id <= workersCount; id++ {
		go worker(id, jobs, results)
	}

	for i := 1; i <= 42; i++ {
		jobs <- i
	}
	close(jobs)

	for result := range results {
		fmt.Println(result)
	}

	duration := time.Since(startTime).Milliseconds()
	var numberOfWorkers string = "<" + strconv.Itoa(workersCount) + ">:"
	fmt.Println("Total time taken" + numberOfWorkers, duration, "msec")
}