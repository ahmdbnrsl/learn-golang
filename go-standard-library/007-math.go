package main 

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Ceil(1.40))//bulatkan ke atas
    fmt.Println(math.Floor(1.80))//bulatkan kebawah
    fmt.Println(math.Round(1.23))//bulatkan ke yg terdekat, tapi aku sama dia jauh, jauh dihati dekat dimata :D
    fmt.Println(math.Min(12, 18))
    fmt.Println(math.Max(18, 12))
}