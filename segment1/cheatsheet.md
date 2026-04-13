### Program structure
```golang
package main

import "fmt"

func main() {
    fmt.Println("hello, world")
}
```

### Declaring variables
```golang
// explicit type
var age int = 17
var name string = "alex"

// short declaration (inside functions)
score := 95.5
passed := true

// declare without value (zero value)
var count int   // count == 0
```

### Basic types
```golang
int        // whole numbers: 0, -3, 42
float64    // decimals:  3.14, -0.5
string     // text:      "hello"
bool       // true or false

// type conversion
x := 5
y := float64(x)
s := fmt.Sprintf("%d", x) // int → string
```

### Printing output
```golang
fmt.Println("hello")        // adds newline
fmt.Print("no newline")    // no newline

// formatted output
fmt.Printf("name: %s\n", name)
fmt.Printf("score: %f\n", score)
fmt.Printf("age: %d\n", age)
fmt.Printf("pass: %v\n", passed) // %v = any
```

### Reading input
```golang
var name string
fmt.Scan(&name)              // reads one word

var age int
fmt.Scan(&age)               // reads an int

var score float64
fmt.Scan(&score)             // reads a float

// read multiple values at once
var x, y int
fmt.Scan(&x, &y)             // user types: 3 7

// read a full line (including spaces)
reader := bufio.NewReader(os.Stdin)
line, _ := reader.ReadString('\n')
```

### if/else
```golang
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("C or below")
}

// condition must be bool — no implicit truthy
```

### switch
```golang
switch grade {
case "A":
    fmt.Println("excellent")
case "B", "C":
    fmt.Println("good")
default:
    fmt.Println("keep going")
}
// no break needed — no fallthrough by default
```

### loops (only for)
```golang
// classic (like a for loop in other languages)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while-style
for count > 0 {
    count--
}

// infinite loop
for {
    // use break to exit
}
```