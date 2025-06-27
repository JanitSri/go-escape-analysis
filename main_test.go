package main

import "testing"

func BenchmarkEscape(b *testing.B) {
	for b.Loop() {
		forceEscape()
	}
}

func BenchmarkNoEscape(b *testing.B) {
	for b.Loop() {
		forceNoEscape()
	}
}
