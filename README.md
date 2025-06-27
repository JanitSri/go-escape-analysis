
# 🧠 Go Compiler Internals: Escape Analysis, SSA IR, and Heap Allocation

This guide helps you analyze heap allocations, escape behavior, SSA Intermediate Representation (SSA IR), and disassembled machine code in Go programs using various compiler tools and environment variables.

---

## 📄 Example Program: `main.go`

```go
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
```

---

## 📄 Benchmark Program: `main_test.go`

```go
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
```

---

## 🔍 Escape Analysis

Escape analysis determines whether a variable should be allocated on the heap or the stack.

### 🔧 Command

```bash
go build -gcflags="-m -l -N" main.go
```

### 🔧 Flags

- `-m`: Show escape analysis diagnostics and inlining decisions
- `-l`: Disable inlining
- `-N`: Disable optimizations

### 🧾 Sample Output

```
main.go:6:12: moved to heap: x
```

---

## 🧠 SSA IR (Static Single Assignment Form)

Visualize SSA form for a specific function.

### 🔧 Command

```bash
GOSSAFUNC=escape go build main.go
```

Or for all functions:

```bash
GOSSAFUNC=* go build main.go
```

- Generates HTML in `/tmp`
- Includes control flow graphs, phi nodes, and allocation hints

---

## 🧬 Compiler Assembly Output

```bash
go tool compile -S main.go
```

Or:

```bash
go build -gcflags="-S" main.go
```

- `-S`: Output Go compiler assembly

---

## 🧼 Disassemble Binary

```bash
go build -o mybin main.go
go tool objdump -s main.main mybin
```

- `-s`: Restrict to specific symbol

---

## 📊 Run Benchmark

```bash
go test -bench=. -benchmem
```

### 🧾 Sample Output

```
BenchmarkEscape-8       10000000    120 ns/op    8 B/op    1 allocs/op
BenchmarkNoEscape-8     20000000      5 ns/op    0 B/op    0 allocs/op
```

---

## 📋 Summary

| Purpose                      | Command                                                                    |
|------------------------------|----------------------------------------------------------------------------|
| Escape analysis              | `go build -gcflags="-m -l -N" main.go`                                     |
| SSA IR (specific function)   | `GOSSAFUNC=escape go build main.go`                                        |
| SSA IR (all functions)       | `GOSSAFUNC=* go build main.go`                                             |
| Compiler-generated assembly  | `go tool compile -S main.go` or `go build -gcflags="-S" main.go`           |
| Disassemble compiled binary  | `go tool objdump -s main.main ./mybin`                                     |
| Benchmark heap allocations   | `go test -bench=. -benchmem`                                               |

---

## 📌 Notes

- The compiler avoids heap allocations unless necessary.
- Using a global `sink` variable forces escape.
- `go test -benchmem` reveals allocations at runtime.
