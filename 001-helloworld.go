package main

import "fmt"

type Console struct{}

func (console Console) log(data any) {
    fmt.Println(data)
}

func main() {
    console := Console{}
    
    console.log("Hello World!!")
}