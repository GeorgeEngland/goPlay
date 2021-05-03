package main

import (
	"fmt"
)

func main() {
	const number = 10

	jobs := make(chan int, number)
	results := make(chan int, number)
	//var res [number]int
	var slice = make([]int, number)
	slice[0] = 0
	slice[1] = 1
	go worker(jobs, results, slice)
	go worker(jobs, results, slice)
	go worker(jobs, results, slice)
	go worker(jobs, results, slice)
	for i := 0; i < number; i++ {
		jobs <- i
	}

	close(jobs)
	for j := 0; j < number; j++ {
		fmt.Println(j, <-results)
	}
}

func worker(jobs <-chan int, results chan<- int, res []int) {
	for n := range jobs {
		results <- fib(n, res)
	}
}

func fib(n int, res []int) int {
	if n <= 1 {
		return n
	}
	if res[n-1] == 0 {
		res[n-1] = fib(n-1, res)
	}
	if res[n-2] == 0 {
		res[n-2] = fib(n-2, res)
	}
	return res[n-1] + res[n-2]
}
