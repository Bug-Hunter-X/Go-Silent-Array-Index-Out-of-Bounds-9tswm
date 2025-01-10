```go
package main

import "fmt"

func main() {

    var x [5]int
    for i := 0; i < 10; i++ {
        if i < len(x) {
            x[i] = i
        } else {
            fmt.Printf("Index out of bounds: %d\n", i)
        }
    }
    fmt.Println(x)
}
```