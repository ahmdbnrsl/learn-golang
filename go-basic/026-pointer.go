package main

import "fmt"

type Address struct{
    City string
    Province string
    Country string
}

func main() {
    address1 := Address{"Cilacap", "Jawa Tengah", "Indonesia"}
    address2 := address1 //pass by value
    
    address2.City = "Brebes"
    fmt.Println(address1)
    fmt.Println(address2)
    
    //pass by reference using pointer
    address3 := &address1 // use & operator
    address3.City = "Bumen"
    
    fmt.Println(address1)
    fmt.Println(address2)
    fmt.Println(address3)
}