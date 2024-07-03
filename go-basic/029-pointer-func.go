package main

import "fmt"

type Address struct{
    City, Province, Country string
}

func changeAddress(address *Address) {
    address.Country = "Jepamg"
}

func main() {
    address := &Address{"Cilacap", "Jawa Tengah", "Indonesia"}
    changeAddress(address)
    
    fmt.Println(address)
}