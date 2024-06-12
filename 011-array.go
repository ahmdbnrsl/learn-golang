package main

import "fmt"

func main() {
    var names [2]string
    
    names[0] = "Via"
    names[1] = "Fitriana"
    
    println(names[0])
    println(names[1])
    
    //direct array
    
    var val = [3]int8{
        10,
        20,
    }
    
    fmt.Println(val)
}