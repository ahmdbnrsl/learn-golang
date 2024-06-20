package main

import "fmt"

func console(data interface{}) interface{} {
    return data
}

func main() {
    var data any = console(234)
    fmt.Println(data)
    switch value := data.(type) {
        case string:
        println("string", value)
        case int:
        println("int", value)
        case bool:
        println("bool", value)
    }
}