package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    i := 0
    for z != math.Sqrt(x) {
        z = z - (z*z - x)/(2*z)
        i++
    }
    fmt.Printf("%d iterations.\n", i)
    return z
   
}

func main() {
    fmt.Println(Sqrt(2))
}