package main

import "fmt"

type Customer struct{
    Name, Address string
    Age int
}

type Console struct{}

func (customer Customer) sayHello() {
    fmt.Println("Hello", customer.Name)
}

func (console Console) log(data string) {
    fmt.Println(data)
}

func main() {
    
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
    
    //struct method
    
    
    fitriana := Customer{Name: "Via Fitriana"}
    fitriana.sayHello()
    beni.sayHello();
    console := Console{}
    console.log("Hello World")
}