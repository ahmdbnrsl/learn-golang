package main

import "fmt"

//empety interface
func console(data interface{}) {
    fmt.Println(data)
}
//any
func printing(data any) {
    fmt.Println(data)
}

func main() {
    console("Hello World")
    printing("Hello World")
}