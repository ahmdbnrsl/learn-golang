package main

import "fmt"

func main() {
    var name string

    name = "Via"
    fmt.Println(name)
    name = "Fitriana"
    fmt.Println(name)

    //optional data type

    var fullName = "Via"
    fmt.Println(fullName)

    //without var

    lastName := "Fitriana"
    lastName = "Fitria"
    fmt.Println(lastName)
    
    //grouping variable
    
    var(
        male = "Beni"
        female = "Via"
    )
    fmt.Println(male)
    fmt.Println(female)
    
    //constant
    
    const man string = "Beni"
    const woman = "Via"
    
    fmt.Println(woman)
    //error
    //man = "Via"
    fmt.Println(man)
    
    //constant grouping
    
    const(
        Man string = "Beni"
        Woman = "Via"
    )
}