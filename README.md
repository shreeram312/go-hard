# Go (Golang) Basics - Quick Reference Notes

## 1. Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

- Every Go file starts with `package` declaration
- Entry point is `func main()` in `package main`
- `fmt.Println()` prints with spaces and a newline

---

## 2. Variables & Data Types

### Data Types

| Type | Description |
|------|-------------|
| `int` | Signed integer (positive + negative), default is **int64** |
| `uint` | Unsigned integer (positive only), default is **uint64** |
| `float64` | Floating point number |
| `bool` | `true` or `false` |
| `string` | Text |
| `rune` | Character (alias for `int32`) |
| `byte` | Byte of memory (alias for `uint8`) |

### Variable Declarations

```go
var x string = "hello"       // explicit type
var m int = 10                // explicit type
z := 3                       // shorthand (type inferred)
const y uint8 = 244          // constant (cannot be changed)
```

**Key notes:**
- `:=` is shorthand for `var` with type inference (only inside functions)
- `const` values cannot be reassigned
- Use `fmt.Printf("%T", z)` to print the type of a variable

---

## 3. Console I/O (Formatting)

```go
fmt.Println("hello", x, 2)    // adds spaces + newline
fmt.Printf("%T", y)            // type
fmt.Printf("%b", z)            // binary
fmt.Printf("%.2f %%\n", a)     // 2 decimal places + literal %
fmt.Printf("%s", b)            // string
fmt.Printf("%e", a)            // scientific notation
```

### Format Verbs Cheat Sheet

| Verb | Purpose |
|------|---------|
| `%T` | Type of variable |
| `%b` | Binary format |
| `%f` | Float (decimal) |
| `%.2f` | Float rounded to 2 decimals |
| `%e` | Scientific notation |
| `%s` | String |
| `%%` | Literal `%` |

### Reading Input

```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
```

---

## 4. Arithmetic & Type Conversion

```go
z := float64(y) / float64(x)       // explicit conversion required
c := a + fmt.Sprint(b)              // int to string (proper concatenation)
og, err := strconv.ParseInt("2324234", 10, 64)  // string to int
```

**Key notes:**
- Go requires **explicit type conversion** (no implicit conversion)
- Always convert to the **smaller** type to avoid overflow
- `string(65)` gives ASCII character, not `"65"` — use `fmt.Sprint()` for proper conversion
- `strconv` package for string <-> number conversions
- `math` package: `math.Min()`, `math.Round()`, etc.

---

## 5. Conditions

```go
if abc > 100 {
    fmt.Println("greater")
} else {
    fmt.Println("lesser")
}
```

### Switch

```go
switch a {
case 10:
    fmt.Println("ten")
case 20:
    fmt.Println("twenty")
default:
    fmt.Println("unknown")
}
```

- `break` is **implicit** in Go switch (no need to write it)
- `fallthrough` keyword forces execution of the next case regardless of condition
- **Naked switch** (no variable): `switch { case a > 6: ... }`

---

## 6. Loops

Go has **only `for`** — no `while` or `do-while`.

```go
// Classic for loop
for i := 0; i < 10; i++ { }

// While-style
a := 1
for a <= 10 { a++ }

// Range loop (for collections)
for idx, val := range slice { }
for _, val := range str { }    // _ to discard index
```

### String Looping

- `len(str)` returns **bytes**, not characters
- ASCII = 1 byte (256 chars), UTF-8 = up to 4 bytes (emojis, non-Latin)
- Use `range` for proper UTF-8 iteration over strings
- `string(str[i])` gives one byte — may split multi-byte characters

---

## 7. Arrays

```go
var arr [3]int                // zero-valued array: [0, 0, 0]
arr := [3]int{1, 2, 4}       // implicit
arr2 := [3]int{1: 2, 2: 3}  // explicit indices
```

- Arrays have **fixed size** — size is part of the type (`[3]int != [5]int`)
- Passed by **value** (copied) when sent to functions
- 2D arrays: `[...][3]int{{1,2,3}, {4,5,6}}`
- `len()` gives number of elements

---

## 8. Slices

```go
sl := []string{"hello", "world"}
sl := make([]int, 10, 10)       // make(type, length, capacity)
sl := arr[1:3]                   // slice from array
```

### Slice internals: pointer + length + capacity

- `len(sl)` — number of elements
- `cap(sl)` — capacity of underlying array
- `append()` adds elements

**Critical:**
- If capacity is available → `append` uses the **same** underlying array
- If capacity is full → `append` creates a **new** underlying array
- Slices are passed by **reference** (modifications reflect in caller)
- Use `slice...` to unpack a slice into variadic args: `sum([]int{1,2,3}...)`

---

## 9. Maps

```go
mp := map[string]int{"a": 3}          // shorthand
mp2 := map[string][]int{"a": {1,2,3}} // slice as value
mp3 := map[uint]uint{}                 // empty map
```

### Map Operations

```go
mp["b"] = 10              // add/update
delete(mp, "b")           // delete key
val, ok := mp["b"]        // check existence (ok is bool)
```

- Maps are **reference types** (like slices)
- Accessing a missing key returns the **zero value** (not an error)
- Use `val, ok` pattern to distinguish missing key from zero value

---

## 10. Functions

```go
// Multiple return values
func add(a int, b int) (int, string) {
    return a + b, "hello world"
}

// Named return values (no explicit return needed)
func sum(nums ...int) (s int) {
    for _, v := range nums {
        s += v
    }
    return
}

// Functions as parameters (callbacks)
func callFunc(callable func(int) int) int {
    return callable(10)
}

// Closures / anonymous functions
value := callFunc(func(x int) int { return x * 4 })

// Currying / higher-order functions
func firstFunc(str string) func(string) string {
    return func(str2 string) string {
        return str + " " + str2
    }
}
```

**Key notes:**
- **Variadic params**: `func sum(nums ...int)` — receives as slice
- **Named returns**: declare return var in signature, bare `return` works
- Functions are **first-class citizens** (can be passed as arguments, returned)
- **Variadic unpacking**: `sum(slice...)` spreads slice into individual args

---

## 11. Structs (Go's "Classes")

```go
type Person struct {
    Name  string    // Capital = exported (public)
    Age   uint
    Sport []Sport
}

// Methods (value receiver — receives a COPY)
func (p Person) GetName() string {
    return p.Name
}
```

- Struct fields starting with **capital letter** are exported (public)
- Lowercase fields are unexported (private to package)
- **Value receiver** = method gets a **copy** — cannot modify original
- To modify original, use **pointer receiver**: `func (p *Person) SetName(name string)`
- Struct embedding for composition (no inheritance in Go)

---

## 12. Generics

```go
func add[T int | float64](x T, y T) T {
    return x + y
}

func getValues[K comparable, V any](mp map[K]V) []V {
    values := []V{}
    for _, v := range mp {
        values = append(values, v)
    }
    return values
}
```

- Type parameters in square brackets: `[T int | float64]`
- `comparable` — any type that supports `==` and `!=`
- `any` — alias for `interface{}` (any type at all)
- Go infers type arguments at call site: `add(1, 2)` (no need for `add[int](1,2)`)

---

## 13. Error Handling & Panic/Recover

Go uses **explicit error returns** instead of exceptions.

```go
og, err := strconv.ParseInt("2324234", 10, 64)
if err != nil {
    fmt.Println(err)
}
```

### Panic & Recover

```go
func deferredFunc() {
    r := recover()    // catches panic
    fmt.Println(r)
}

func main() {
    defer deferredFunc()   // runs when function exits
    panic("hello world")   // stops program, triggers deferred funcs
}
```

- `panic()` — stops execution, runs deferred functions
- `recover()` — catches a panic (only useful inside `defer`)
- `defer` — schedules function call for when surrounding function returns
- **Best practice**: use error returns, not panic, for expected failures

---

## 14. Interfaces

```go
type Shape interface {
    getPerimeter() uint
}

func (s Square) getPerimeter() uint {
    return 4 * s.width
}
```

- Interfaces are **implicitly implemented** (no `implements` keyword)
- A struct satisfies an interface by implementing **all** its methods
- Interface variables can hold any concrete type that satisfies the interface
- `interface{}` (or `any`) is the empty interface — all types satisfy it

---

## 15. CLI Arguments

```go
args := os.Args     // []string of all CLI args
args[0]             // program name
args[1]             // first argument
```

---

## Quick Rules of Thumb

| Rule | Detail |
|------|--------|
| No unused variables | Code won't compile |
| No unused imports | Code won't compile |
| Package visibility | Capital = exported, lowercase = private |
| Error handling | Check `err != nil` always, don't ignore errors |
| No inheritance | Use composition and interfaces |
| No while loop | Use `for condition { }` |
| No generics until Go 1.18 | Now supported |
| Pass by value | Arrays, structs, basics are copied; slices, maps, channels are references |
| `go run main.go` | Run a file |
| `go build main.go` | Compile to binary |

---

## Project Structure

```
basics-of-go/
  01-hello-world/
  02-variables/
  03-console/
  04-arithmetic/
  05-conditions/
  06-loops/
  07-arrays/
  08-slices/
  09-maps/
  10-functions/
  11-structs/
  12-generics/
  13-error-handling/
  interfaces/
  scan-go/
cli/
  print-cli/
```
