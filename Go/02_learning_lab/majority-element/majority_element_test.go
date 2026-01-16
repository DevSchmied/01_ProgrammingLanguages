package main

import "testing"

// helper: large input with guaranteed majority element
func generateTestData(size int, majority int) []int {
	data := make([]int, size)

	for i := 0; i < (size/2)+1; i++ {
		data[i] = majority
	}

	for i := (size / 2) + 1; i < size; i++ {
		data[i] = i
	}

	return data
}
func BenchmarkMajorityElement_Map(b *testing.B) {
	data := generateTestData(100_000, 42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = majorityElement(data)
	}
}

func BenchmarkMajorityElement_BoyerMoore(b *testing.B) {
	data := generateTestData(100_000, 42)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = majorityElementOptimized(data)
	}
}
