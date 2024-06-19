package main

import "fmt"

func main() {
    type Customer struct{
        Name, Address string
        Age int
    }
    
    var customer Customer
    customer.Name = "Ahmad Beni Rusli"
    customer.Address = "Cilacap"
    customer.Age = 19
    
    fmt.Println(customer)
    
    //Struct literals 
    via := Customer {
        Name: "Via Fitriana",
        Address: "Sidareja",
        Age: 100,
    }
    fmt.Println(via)
    //or
    beni := Customer{"Ahmad Beni Rusli", "Cilacap", 19}
    fmt.Println(beni)
}