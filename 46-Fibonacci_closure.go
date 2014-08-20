package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    current := 0
    last_fib := 1
    return func() int {
        current, last_fib = last_fib, current + last_fib
        return current
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}