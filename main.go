package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	rand.Seed(123456789)
	numbers := buildArray()
	fmt.Println(sumSequentially(numbers))
	fmt.Println(sumConcurrently(numbers))
	fmt.Println(sumConcurrentlyDataRace(numbers))

}

func buildArray() []int {
	const length = 1e6
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
}

func sumSequentially(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func sumConcurrently(numbers []int) int {
	var sum int64

	goroutines := runtime.NumCPU()
	sliceLength := len(numbers) / goroutines

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		numbersSlice := numbers[(0+i)*sliceLength : (1+i)*sliceLength]
		go func() {
			var intermediateSum int64
			for _, num := range numbersSlice {
				intermediateSum += int64(num)
			}
			atomic.AddInt64(&sum, intermediateSum)
			wg.Done()
		}()
	}
	wg.Wait()
	return int(sum)
}

func sumConcurrentlyDataRace(numbers []int) int {
	var sum int64

	goroutines := runtime.NumCPU()
	sliceLength := len(numbers) / goroutines

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		numbersSlice := numbers[(0+i)*sliceLength : (1+i)*sliceLength]
		go func() {
			var intermediateSum int64
			for _, num := range numbersSlice {
				intermediateSum += int64(num)
			}
			sum += intermediateSum
			wg.Done()
		}()
	}
	wg.Wait()
	return int(sum)
}
