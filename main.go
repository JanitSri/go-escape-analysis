package main

var sink any

func escape() *int {
	x := 42
	return &x // Escapes to heap
}

func noEscape() int {
	y := 99
	return y // Stays on stack
}

func forceEscape() {
	ptr := escape()
	sink = ptr // Store pointer in global to force heap allocation
}

func forceNoEscape() {
	val := noEscape()
	sink = val // Still stack-allocated, as it's a value not a pointer
}

func main() {
	forceEscape()
	forceNoEscape()
}
