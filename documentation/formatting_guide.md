# Deep Dive: String Formatting & Verbs in Go (`fmt` Package)

In Go (Golang), formatted output and string interpolation are handled primarily by functions in the standard library's `fmt` package—such as `fmt.Printf`, `fmt.Sprintf`, `fmt.Fprintf`, and `fmt.Errorf`.

This guide provides an exhaustive breakdown of formatting verbs, specifiers, flags, precision controls, and type-specific rules.

---

## 📌 Core Distinction: `fmt.Println` vs `fmt.Printf`

Before diving into verbs, it is essential to understand why formatting verbs work only in `Printf`-family functions:

| Function          | Accepts Verbs (`%g`, `%s`)? | Appends Newline automatically? | Purpose                                                                                                   |
| :---------------- | :-------------------------: | :----------------------------: | :-------------------------------------------------------------------------------------------------------- |
| **`fmt.Println`** |            ❌ No            |             ✅ Yes             | Prints values separated by spaces with an automatic trailing `\n`. Treats `%` characters as literal text. |
| **`fmt.Printf`**  |           ✅ Yes            |             ❌ No              | Formats according to specifiers and prints to standard output. Requires explicit `\n`.                    |
| **`fmt.Sprintf`** |           ✅ Yes            |             ❌ No              | Returns the formatted string instead of printing it to stdout.                                            |
| **`fmt.Fprintf`** |           ✅ Yes            |             ❌ No              | Writes formatted output to any `io.Writer` (e.g., files, network streams, HTTP responses).                |

---

## 📊 Summary Table of Common Verbs

|        Verb         | Name / Type          | Output Description                            | Code Example                | Rendered Output                     |
| :-----------------: | :------------------- | :-------------------------------------------- | :-------------------------- | :---------------------------------- |
|      **`%v`**       | Default Value        | Natural default format for any value          | `fmt.Printf("%v", 42)`      | `42`                                |
|      **`%+v`**      | Struct Fields        | Adds field names for structs                  | `fmt.Printf("%+v", p)`      | `{Name:Alice Age:30}`               |
|      **`%#v`**      | Go Syntax            | Valid Go code representation                  | `fmt.Printf("%#v", p)`      | `main.Person{Name:"Alice", Age:30}` |
|      **`%T`**       | Type Name            | Go type of the variable                       | `fmt.Printf("%T", 3.14)`    | `float64`                           |
|      **`%%`**       | Literal Percent      | Literal `%` character                         | `fmt.Printf("100%%")`       | `100%`                              |
|      **`%s`**       | String / Bytes       | Raw unquoted characters                       | `fmt.Printf("%s", "Hello")` | `Hello`                             |
|      **`%q`**       | Double-quoted        | Quoted string with escape sequences           | `fmt.Printf("%q", "a\\nb")` | `"a\\nb"`                           |
|      **`%x`**       | Hexadecimal (string) | Lowercase hex encoding per byte               | `fmt.Printf("%x", "Go")`    | `476f`                              |
|      **`%X`**       | Hexadecimal (string) | Uppercase hex encoding per byte               | `fmt.Printf("%X", "Go")`    | `476F`                              |
|      **`%d`**       | Base-10 Integer      | Decimal integer notation                      | `fmt.Printf("%d", 255)`     | `255`                               |
|      **`%b`**       | Base-2 Integer       | Binary representation                         | `fmt.Printf("%b", 5)`       | `101`                               |
|      **`%o`**       | Base-8 Integer       | Octal representation                          | `fmt.Printf("%o", 255)`     | `377`                               |
|      **`%x`**       | Base-16 Integer      | Lowercase hex integer                         | `fmt.Printf("%x", 255)`     | `ff`                                |
|      **`%X`**       | Base-16 Integer      | Uppercase hex integer                         | `fmt.Printf("%X", 255)`     | `FF`                                |
|      **`%c`**       | Character (Rune)     | Unicode character for integer                 | `fmt.Printf("%c", 65)`      | `A`                                 |
|      **`%U`**       | Unicode Format       | `U+XXXX` format                               | `fmt.Printf("%U", 65)`      | `U+0041 'A'`                        |
| **`%f`** / **`%F`** | Decimal Float        | Standard decimal float notation               | `fmt.Printf("%f", 12.34)`   | `12.340000`                         |
| **`%g`** / **`%G`** | Compact Float        | Removes trailing zeros; uses exponent if huge | `fmt.Printf("%g", 212.0)`   | `212`                               |
| **`%e`** / **`%E`** | Scientific Float     | Exponent notation (`1.234e+01`)               | `fmt.Printf("%e", 12.34)`   | `1.234000e+01`                      |
|      **`%t`**       | Boolean              | `true` or `false`                             | `fmt.Printf("%t", true)`    | `true`                              |
|      **`%p`**       | Pointer              | Memory address (hexadecimal)                  | `fmt.Printf("%p", &var)`    | `0xc00008a008`                      |

---

## 🔍 In-Depth Breakdown by Data Category

### 1. Floating-Point Numbers (`%g` vs `%f` vs `%e`)

Formatting floating-point numbers (`float32`, `float64`) depends on the desired precision and formatting style:

#### `%f` — Fixed Decimal Representation

Always displays a fixed number of decimal digits (default is 6 digits after the decimal point).

```go
num := 212.0
fmt.Printf("%f\n", num)   // Output: 212.000000
fmt.Printf("%.2f\n", num) // Output: 212.00 (controlled width/precision)
```

#### `%g` — Compact / Adaptive Representation

`%g` automatically chooses between standard decimal notation and scientific exponent notation depending on the magnitude of the number. Most importantly, it suppresses unnecessary trailing zeros.

```go
temp1 := 212.0
temp2 := 0.00000000123

fmt.Printf("%g\n", temp1) // Output: 212
fmt.Printf("%g\n", temp2) // Output: 1.23e-09 (switches to scientific notation when tiny)
```

#### `%e` — Scientific / Exponent Representation

Always displays numbers in scientific exponential form.

```go
val := 12345.67
fmt.Printf("%e\n", val) // Output: 1.234567e+04
```

### 2. General / Universal Verbs (`%v`, `%T`, `%%`)

These verbs work with variables of any type:

```go
type User struct {
    ID   int
    Name string
}

u := User{ID: 101, Name: "Gopher"}

// %v: Standard value print
fmt.Printf("%v\n", u) // Output: {101 Gopher}

// %+v: Struct field names included
fmt.Printf("%+v\n", u) // Output: {ID:101 Name:Gopher}

// %#v: Full Go syntax representation
fmt.Printf("%#v\n", u) // Output: main.User{ID:101, Name:"Gopher"}

// %T: Type inspection
fmt.Printf("%T\n", u) // Output: main.User

// %%: Printing literal percent sign
fmt.Printf("Progress: 95%%\n") // Output: Progress: 95%
```

### 3. Strings and Byte Slices (`%s`, `%q`, `%x`)

```go
msg := "Hello, 世界"

// %s: Unquoted UTF-8 string
fmt.Printf("%s\n", msg) // Output: Hello, 世界

// %q: Safely double-quoted string with escape sequences
fmt.Printf("%q\n", "Hello\nWorld") // Output: "Hello\nWorld"

// %x / %X: Hexadecimal bytes
fmt.Printf("%x\n", "Go") // Output: 476f
```

### 4. Integers (`%d`, `%b`, `%o`, `%x`, `%c`)

Go supports formatting integers into various numeric bases:

```go
val := 255

fmt.Printf("Decimal:     %d\n", val) // Output: Decimal:     255
fmt.Printf("Binary:      %b\n", val) // Output: Binary:      11111111
fmt.Printf("Octal:       %o\n", val) // Output: Octal:       377
fmt.Printf("Hex (lower): %x\n", val) // Output: Hex (lower): ff
fmt.Printf("Hex (upper): %X\n", val) // Output: Hex (upper): FF

charVal := 65
fmt.Printf("Character:   %c\n", charVal) // Output: Character:   A
```

### 5. Pointers (`%p`)

To view memory addresses of variables, slices, or pointers, use `%p`:

```go
x := 42
fmt.Printf("Pointer address: %p\n", &x) // Output: Pointer address: 0xc000012088
```

## 🛠️ Width, Alignment, and Precision Modifiers

You can fine-tune formatting by adding numbers between `%` and the verb letter.

### Width & Alignment

- `%5d`: Right-align integer in a field of width 5.
- `%-5d`: Left-align integer in a field of width 5.
- `%05d`: Pad number with leading zeros up to width 5.

```go
fmt.Printf("|%5d|\n", 42)  // Output: |   42|
fmt.Printf("|%-5d|\n", 42) // Output: |42   |
fmt.Printf("|%05d|\n", 42) // Output: |00042|
```

### Precision Controls

- `%.2f`: Print float with exactly 2 decimal places.
- `%.4s`: Truncate string to a maximum of 4 characters.

```go
fmt.Printf("%.2f\n", 3.14159) // Output: 3.14
fmt.Printf("%.4s\n", "Golang") // Output: Gola
```

## 💡 Practical Code Example

Save and run the following standalone script to test all major verbs:

```go
package main

import "fmt"

type Account struct {
    Username string
    Balance  float64
}

func main() {
    acc := Account{Username: "alex_go", Balance: 1540.50}

    fmt.Println("=== Default vs Field-aware Struct Formatting ===")
    fmt.Printf("%%v  : %v\n", acc)
    fmt.Printf("%%+v : %+v\n", acc)
    fmt.Printf("%%#v : %#v\n", acc)

    fmt.Println("\n=== Float Formatting Comparison ===")
    fmt.Printf("%%f  : %f\n", acc.Balance)
    fmt.Printf("%%g  : %g\n", acc.Balance)
    fmt.Printf("%%.2f: %.2f\n", acc.Balance)

    fmt.Println("\n=== String and Pointer Formatting ===")
    fmt.Printf("%%s  : %s\n", acc.Username)
    fmt.Printf("%%q  : %q\n", acc.Username)
    fmt.Printf("%%p  : %p\n", &acc)
}
```
