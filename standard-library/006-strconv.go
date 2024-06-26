package main 

import (
    "strconv"
    "fmt"
)

func main() {
    boolean, err := strconv.ParseBool("true")
    if err == nil {
        fmt.Println(boolean)
    } else {
        fmt.Println(err.Error())
    }
    integer, errorr := strconv.Atoi("123")
    if err == nil {
        fmt.Println(integer)
    } else {
        fmt.Println(errorr.Error())
    }
    bin := strconv.FormatInt(999, 2)
    bin2 := strconv.Itoa(999)
    fmt.Println(bin, bin2)
}