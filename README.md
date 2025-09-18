# LeetCode Solutions Repository

This repository contains solutions to LeetCode problems organized by algorithmic patterns and data structures. Each solution includes the original problem link.

## Running the Programs

### JavaScript Programs

JavaScript solutions can be run directly using Node.js:

```bash
# Navigate to the specific problem directory
cd path/to/leet-code/array-string

# Run a JavaScript solution
node merge-sorted-array.js
```

### Go Programs

Import the function you wish to run in `main.go`

```go
package main

import (
    "fmt"
    "leet-code/array-string"
)

func main() {
    result := arraystring.StrStr("hello", "ll") // Edit here
    fmt.Println(result)
}
```

```bash
# Navigate to the root folder
cd path/to/leet-code

# Run the main file
go run .
```

## Debugging the Programs (VS Code & Forks)

### JavaScript Programs

Simply run the file from a `JavaScript Debug Terminal`, as you would in a `bash` terminal.

### Go Programs

Debug from the Launch configuration at `.vscode/launch.json`