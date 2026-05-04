### What is a module
```golang
// a module is a collection of Go files
// that belong together as a project
// every module has a go.mod file at its root

// go.mod looks like this:
module github.com/you/myproject

go 1.22

require (
    github.com/fatih/color v1.16.0
)

// it tracks:
// - the module name
// - the Go version
// - all third party dependencies
```

### Creating a module
```bash
# initialise a module
go mod init myproject

# this creates go.mod in your folder
# you only need to do this once per project

# build and run your program
go run main.go

# or build an executable
go build
```

### The standard library
```golang
// standard library packages are built in —
// no installation needed, just import
import (
    "fmt"       // printing
    "strings"   // string manipulation
    "strconv"   // type conversion
    "math"      // math functions
    "math/rand" // random numbers
    "os"        // operating system
    "bufio"     // buffered input
    "errors"    // error creation
)

// browse all standard packages at:
// pkg.go.dev/std
```

### Installing third party packages
```bash
# find packages at pkg.go.dev
# then install with go get:
go get github.com/fatih/color

# go get will:
# 1. download the package
# 2. add it to go.mod
# 3. update go.sum

# to remove a package you no longer use:
go mod tidy

# go.sum is generated automatically — you never edit it by hand
# it records a fingerprint of every downloaded package
# so Go can verify nothing has been tampered with

# both go.mod and go.sum should be committed to version control
# the downloaded packages themselves (in the module cache) should not
```

### Using a third party package
```golang
// after go get, import it just like a standard library package
import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    // use the package name (last part of the import path)
    color.Red("this prints in red")
    color.Green("this prints in green")

    // or create a reusable printer
    highlight := color.New(color.FgYellow, color.Bold)
    highlight.Println("this is bold and yellow")
}
```
