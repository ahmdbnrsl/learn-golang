package main 

import (
    "slices"
    "fmt"
)

func main() {
    names := []string{
        "Via",
        "Fitriana",
        "Beni",
    }
    values := []int{
        19,
        20,
        19,
    }
    
    fmt.Println(slices.Min(values))
    fmt.Println(slices.Max(values))
    fmt.Println(slices.Contains(names, "Via"))
    fmt.Println(slices.Index(names, "Beni"))
}