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
    
    var val1 = [...]int8{
        20,
        50,
        50,
    }
    
    fmt.Println(val)
    println(len(names))
    println(len(val))
    val1[1] = 100
    println(len(val1))
    println(val1[1])
}