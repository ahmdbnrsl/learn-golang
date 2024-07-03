package main

import "fmt"

func main() {
    name := "Via Fitriana"
    
    switch name {
    case "Via Fitriana":
        fmt.Println("Hallo my dear")
    case "Via":
        fmt.Println("Hallo sayang")
    default:
        fmt.Println("Aku tidak suka kamu")
    }
    
    //shirt statement 
    switch length := len(name); length > 5 {
        case true:
        println("Haloooo sayang")
        case false:
        println("Hehe")
    }
    
    //without condition
    switch {
        case len(name) > 10:
        println("Wuhuhu")
        case len(name) > 5:
        println("hehehe")
        default:
        println("love you")
    }
}