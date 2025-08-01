//You need a way to wait until your goroutines finish processing. Here’s a simple fix using sync.WaitGroup:

package main

import (
    "fmt"
    "sync"
)

func main() {
    cnp := make(chan func(), 10)
    var wg sync.WaitGroup

    for i := 0; i < 4; i++ {
        go func() {
            for f := range cnp {
                f()
                wg.Done()
            }
        }()
    }

    wg.Add(1)
    cnp <- func() {
        fmt.Println("HERE1")
    }

    fmt.Println("Hello")

    wg.Wait()
    close(cnp)
}

//This ensures "HERE1" gets printed before main exits.
