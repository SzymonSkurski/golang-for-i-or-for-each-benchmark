package main

import "testing"

func BenchmarkCloneForI(b *testing.B) {
	c := getMyColors()

	for i := 0; i < b.N; i++ {
		cloneForI(c)
	}
}

func BenchmarkCloneForEach(b *testing.B) {
	c := getMyColors()

	for i := 0; i < b.N; i++ {
		cloneForEach(c)
	}
}
