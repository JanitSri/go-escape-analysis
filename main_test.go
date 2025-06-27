package main

import "testing"

func BenchmarkEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forceEscape()
	}
}

func BenchmarkNoEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		forceNoEscape()
	}
}
