package main

import "testing"

func BenchmarkSumSequential(b *testing.B) {
	numbers := buildArray()
	for n := 0; n < b.N; n++ {
		sumSequentially(numbers)
	}
}

func BenchmarkSumConcurrently(b *testing.B) {
	numbers := buildArray()
	for n := 0; n < b.N; n++ {
		sumConcurrently(numbers)
	}
}

func BenchmarkSumConcurrentlyDataRace(b *testing.B) {
	numbers := buildArray()
	for n := 0; n < b.N; n++ {
		sumConcurrentlyDataRace(numbers)
	}
}
