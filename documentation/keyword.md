# Go Language Keywords Reference

The Go programming language (Golang) is designed for simplicity and readability. It defines exactly **25 keywords** in its language specification.

---

## 📋 Complete List of Keywords

| # | Keyword | Category | Description |
|---|---|---|---|
| 1 | `break` | Control Flow | Terminates loop or switch execution |
| 2 | `case` | Control Flow | Specifies a block in switch / select statements |
| 3 | `chan` | Composite Types | Defines a channel type for concurrency |
| 4 | `const` | Declarations | Declares a constant value |
| 5 | `continue` | Control Flow | Skips to the next iteration of a loop |
| 6 | `default` | Control Flow | Specifies default case in switch / select |
| 7 | `defer` | Control Flow | Defers execution of a function until surrounding function returns |
| 8 | `else` | Control Flow | Specifies alternative branch for `if` statement |
| 9 | `fallthrough` | Control Flow | Forces execution to flow to the next case in switch |
| 10 | `for` | Control Flow | Defines a loop (Go's only looping construct) |
| 11 | `func` | Declarations | Declares a function or method |
| 12 | `go` | Control Flow | Starts a new goroutine (concurrent execution) |
| 13 | `goto` | Control Flow | Transfers control to a labeled statement |
| 14 | `if` | Control Flow | Conditional execution statement |
| 15 | `import` | Declarations | Imports package dependencies |
| 16 | `interface` | Program Structure | Defines an interface type |
| 17 | `map` | Composite Types | Defines a hash map / key-value pair type |
| 18 | `package` | Declarations | Defines the package name for the current file |
| 19 | `range` | Control Flow | Iterates over elements in arrays, slices, maps, or channels |
| 20 | `return` | Control Flow | Returns values from a function |
| 21 | `select` | Control Flow | Waits on multiple channel operations |
| 22 | `struct` | Composite Types | Defines a structured custom type |
| 23 | `switch` | Control Flow | Multi-way conditional branch statement |
| 24 | `type` | Program Structure | Declares a new custom type |
| 25 | `var` | Declarations | Declares a variable |

---

## 🏷️ Categorized View

### 1. Declarations (5)
* `const`
* `func`
* `import`
* `package`
* `var`

### 2. Composite Types (3)
* `chan`
* `map`
* `struct`

### 3. Control Flow (15)
* `break`
* `case`
* `continue`
* `default`
* `defer`
* `else`
* `fallthrough`
* `for`
* `go`
* `goto`
* `if`
* `range`
* `return`
* `select`
* `switch`

### 4. Program Structure (2)
* `interface`
* `type`

---

## 💡 Notes & Language Design

* **Minimalist Design**: Go intentionally maintains a minimal set of 25 keywords (compared to 50+ in Java, 90+ in C++, and 35+ in Python).
* **Pre-declared Identifiers**: Types like `int`, `string`, `bool`, and built-in functions like `len`, `append`, `make`, `new` are **pre-declared identifiers**, not keywords.