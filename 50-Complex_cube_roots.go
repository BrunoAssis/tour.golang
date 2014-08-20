package main

import (
    "fmt"
    "math/cmplx"
)

func Cbrt(x complex128) complex128 {
    z := complex128(1)
    for cmplx.Abs(z - cmplx.Pow(x, 1/3.0)) > 0.0000001 {       
        z = z - (z*z*z - x)/(3*z*z)
    }
    return z
}

func main() {
    fmt.Println(Cbrt(2))
}