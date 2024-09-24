# For Loop

> The `for` loop is the only looping construct in Go, but it is versatile enough to be used in different ways, including traditional loops, infinite loops, and range-based iterations.

## 1. Basic `for` Loop Structure
The syntax of the `for` loop in Go is similar to other C-like languages but more flexible.

```go
for initialization; condition; post {
    // code to be executed
}
```
+ **Initialization**: Typically used to initialize a counter variable.
+ **Condition**: The loop runs as long as this condition is `true`.
+ **Post**: Executes after each iteration, often used to update the counter.

### Example:
```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
    }
}
```

### Explanation:
- The loop starts with `i = 0`, and as long as `i < 5`, it prints the current value of `i`.
- The `i++` increments `i` by 1 after each iteration.