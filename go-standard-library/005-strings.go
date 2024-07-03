package main 

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Trim("Ahmad Beni", " Beni"))
    fmt.Println(strings.Contains("Via Fitriana", "Hatiku"))
    fmt.Println(strings.Split("Via Fitriana", " ")) //returning slice
    fmt.Println(strings.ReplaceAll("Via Fitriana", "Via Fitriana", "My dear"))
}