package main

import "fmt"

type Address struct{
    City, Province, Country string
}

func main() {
    alamat1 := new(Address)
    alamat2 := alamat1
    
    alamat2.Country = "Jepang"
    
    fmt.Println(alamat1) //alamat1 berubah 
    fmt.Println(alamat2)
}